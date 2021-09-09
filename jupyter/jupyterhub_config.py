# Configuration file for Jupyter Hub

from jinja2 import Template
from oauthenticator.github import GitHubOAuthenticator
from jupyterhub.auth import PAMAuthenticator
from jupyterhub.traitlets import Command
from jupyterhub.apihandlers.base import APIHandler
from jupyterhub.handlers.login import LogoutHandler
from jupyterhub.utils import maybe_future
import asyncio
import yaml
from tornado.escape import url_escape
from tornado.httputil import url_concat
from tornado.httpclient import HTTPRequest
from jupyterhub.handlers.base import BaseHandler
from traitlets import Unicode
SESSION_COOKIE_NAME = 'jupyterhub-session-id'
c = get_config()
c.JupyterHub.log_level = 10
c.JupyterHub.shutdown_on_logout = True
try:
    from http.cookies import Morsel
except ImportError:
    from Cookie import Morsel

Morsel._reserved[str('samesite')] = str('SameSite')

class WrappedGitHubAuthenticator(GitHubOAuthenticator):
    async def authenticate(self,handler,data):
        result = await GitHubOAuthenticator.authenticate(self, handler, data)
        result['name'] = 'github_user_' + result['name']
        return result
yamlconfig = {}
with open("/application/config/config.yaml", "r") as f:
    yamlconfig = yaml.full_load(f)
auth_url = "http://localhost:8080" + "/v1/users/auth"
backendhost = "localhost:8080"
class MyAuthenticator(WrappedGitHubAuthenticator,PAMAuthenticator):
    async def add_user(self, user):
        """Hook called whenever a new user is added

        If self.create_system_users, the user will attempt to be created if it doesn't exist.
        """
        user_exists = await maybe_future(self.system_user_exists(user))
        
        if not user_exists:
            if self.create_system_users:
                await maybe_future(self.add_system_user(user))
                dir = Path('/app/db/notebook_dir/' + user.name)
                if not dir.exists():
                    print('adding user',user.name)
                    os.makedirs(f'/app/db/notebook_dir/{user.name}/examples',exist_ok=True)
                    subprocess.check_call(['cp', '-r', '/srv/ipython/examples', '/app/db/notebook_dir/' + user.name + "/examples"])
                subprocess.check_call(['chown', '-R', user.name, '/app/db/notebook_dir/' + user.name ])
            else:
                raise KeyError("User %s does not exist." % user.name)

        await maybe_future(super().add_user(user))

    async def authenticate(self,handler,data):
        username = data['username']
        password = data['password']
        url = auth_url + "?user_name=" + username + "&token=" + password
        print('backendhost',backendhost)
        req = HTTPRequest(
            url,
            method="GET",
            headers={"Host": backendhost,
                     "Pragma":"no-cache",
                     "Cache-Control":"no-cache",
                     "Content-Type":"application/json;charset=UTF-8",
                     "Origin":backendhost,
                     "Sec-Fetch-Site":"same-site",
                     "Sec-Fetch-Mode":"cors",
                     "Sec-Fetch-Dest":"empty",
                     "Referer":backendhost,
                     "Accept-Encoding":"gzip,deflate,br",
                     "Accept-Language":"zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,zh-HK;q=0.6"
                     }
        )
        resp_data = await self.fetch(req)
        print(resp_data)
        if resp_data.get("message") and resp_data["message"] == "success":
            print(resp_data["data"]["id"])
            print('yes!!!')
            theUserName = 'user_' + resp_data["data"]["id"][0:10]
            print(theUserName)
            return theUserName
        else:
            raise ValueError("Authenticate Failed")

class ExtloginHandler(BaseHandler):
    """Render the login page."""

    def _render(self, login_error=None, username=None,token=None):
        context = {
            "next": url_escape(self.get_argument('next', default='')),
            "username": username,
            "password":token,
            "login_error": login_error,
            "login_url": self.settings['login_url'],
            "authenticator_login_url": url_concat(
                self.authenticator.login_url(self.hub.base_url),
                {'next': self.get_argument('next', '')},
            ),
        }
        custom_html = Template(
            self.authenticator.get_custom_html(self.hub.base_url)
        ).render(**context)
        return self.render_template('external_login.html',
                **context,
                custom_html=custom_html,
        )
    async def get(self):
        user = self.get_current_user()
        username = self.get_argument('username', default='')
        token = self.get_argument('token',default='')
        print(username,token)
        self.finish(await self._render(username=username,token=token))
    def _backend_logout_cleanup(self, name):
        """Default backend logout actions

        Send a log message, clear some cookies, increment the logout counter.
        """
        self.log.info("User logged out: %s", name)
        self.clear_login_cookie()
        self.statsd.incr('logout')

    async def default_handle_logout(self):
        """The default logout action

        Optionally cleans up servers, clears cookies, increments logout counter
        Cleaning up servers can be prevented by setting shutdown_on_logout to
        False.
        """
        user = self.current_user
        if user:
            if self.shutdown_on_logout:
                await self._shutdown_servers(user)

            self._backend_logout_cleanup(user.name)
    async def _shutdown_servers(self, user):
        """Shutdown servers for logout

        Get all active servers for the provided user, stop them.
        """
        active_servers = [
            name
            for (name, spawner) in user.spawners.items()
            if spawner.active and not spawner.pending
        ]
        if active_servers:
            self.log.info("Shutting down %s's servers", user.name)
            futures = []
            for server_name in active_servers:
                futures.append(maybe_future(self.stop_single_user(user, server_name)))
            await asyncio.gather(*futures)

c.JupyterHub.extra_handlers = [("/external/login",ExtloginHandler)]

c.JupyterHub.authenticator_class = MyAuthenticator




import os
import sys
import subprocess
from pathlib import Path

join = os.path.join

here = os.path.dirname(__file__)
root = os.environ.get('OAUTHENTICATOR_DIR', here)
sys.path.insert(0, root)
allowed_users = set()
def update_allowed_users():
    with open(join(root, 'userlist')) as f:
        for line in f:
            if not line:
                continue
            parts = line.split()
            name = parts[0]
            allowed_users.add(name)
def check_allowed(username):
    update_allowed_users()
    return username in allowed_users
c.JupyterHub.template_paths = [os.environ['OAUTHENTICATOR_DIR'] + '/templates']
c.JupyterHub.tornado_settings = {
    'headers': {
        'Access-Control-Allow-Origin':'*',
        'Content-Security-Policy': "frame-ancestors * " ,
    },
}
c.NotebookApp.tornado_settings = {
  'headers': {
     'Access-Control-Allow-Origin':'*',
      'Content-Security-Policy': "frame-ancestors * ",
   }
}
def generateExampleHook(spawner):
    username = spawner.user.name
    print(f'/app/db/notebook_dir/{username}/examples/en-horizontal-learning-task.ipynb')
    os.makedirs(f'/app/db/notebook_dir/{username}/examples',exist_ok=True)
    enFile = Path(f'/app/db/notebook_dir/{username}/examples/en-horizontal-learning-task.ipynb')
    if not enFile.exists():
      subprocess.check_call(['cp', '-f', '/srv/ipython/examples/en-horizontal-learning-task.ipynb', f'/app/db/notebook_dir/{username}/examples/en-horizontal-learning-task.ipynb'])
    zhFile = Path(f'/app/db/notebook_dir/{username}/examples/zh-horizontal-learning-task.ipynb')
    if not zhFile.exists():
      subprocess.check_call(['cp', '-f', '/srv/ipython/examples/zh-horizontal-learning-task.ipynb', f'/app/db/notebook_dir/{username}/examples/zh-horizontal-learning-task.ipynb'])
c.Spawner.pre_spawn_hook = generateExampleHook
c.Spawner.http_timeout = 100
c.Spawner.notebook_dir = '/app/db/notebook_dir/{username}'
c.Spawner.args = ['''--ServerApp.tornado_settings={
  'headers':{
    'Access-Control-Allow-Origin':'*',
    'Content-Security-Policy': "frame-ancestors 'self' * ",
    'cookie_options': {'samesite':'None','Secure':True},
  }
} --ServerApp.allow_remote_access=True''']
c.Spawner.args = ['''--config=/application/jupyter_jupyterlab_server_config.py''']
c.Spawner.cmd = ["jupyter-labhub"]
c.MyAuthenticator.create_system_users = True
# ssl config
ssl = "/application/ssl"
keyfile = join(ssl, 'jupyter.key')
certfile = join(ssl, 'jupyter.crt')
print('keyfile',keyfile)
print('certfile',certfile)
# if os.path.exists(keyfile):
#     c.JupyterHub.ssl_key = keyfile
# if os.path.exists(certfile):
#     c.JupyterHub.ssl_cert = certfile
