# Configuration file for Jupyter Hub

from jinja2 import Template
from oauthenticator.github import GitHubOAuthenticator
from jupyterhub.auth import PAMAuthenticator
from jupyterhub.traitlets import Command
from jupyterhub.apihandlers.base import APIHandler
from jupyterhub.handlers.login import LogoutHandler
import jupyterhub.orm as orm
import json
from jupyterhub.utils import maybe_future
import asyncio

from tornado.escape import url_escape
from tornado.httputil import url_concat
from tornado import web
import pipes
from subprocess import PIPE
from subprocess import Popen
from subprocess import STDOUT
from tornado.httpclient import HTTPRequest
from jupyterhub.handlers.base import BaseHandler
from traitlets import Unicode
SESSION_COOKIE_NAME = 'jupyterhub-session-id'
c = get_config()
c.JupyterHub.log_level = 10
class WrappedGitHubAuthenticator(GitHubOAuthenticator):
    async def authenticate(self,handler,data):
        result = await GitHubOAuthenticator.authenticate(self, handler, data)
        result['name'] = 'github_user_' + result['name']
        return result

class MyAuthenticator(WrappedGitHubAuthenticator,PAMAuthenticator):
    async def add_user(self, user):
        """Hook called whenever a new user is added

        If self.create_system_users, the user will attempt to be created if it doesn't exist.
        """
        user_exists = await maybe_future(self.system_user_exists(user))

        if not user_exists:
            if self.create_system_users:
                await maybe_future(self.add_system_user(user))
                dir = Path('/home/' + user.name + '/examples')
                print('add user',user.name)
                if not dir.exists():
                    print('adding user',user.name)
                    subprocess.check_call(['cp', '-r', '/srv/ipython/examples', '/home/' + user.name + '/examples'])
                    subprocess.check_call(['chown', '-R', user.name, '/home/' + user.name + '/examples'])
            else:
                raise KeyError("User %s does not exist." % user.name)

        await maybe_future(super().add_user(user))

    async def authenticate(self,handler,data):
        # if data is not None and 'localAuth' in data:
        #     result =  PAMAuthenticator.authenticate(self,handler,data)
        #     return result
        # else:
        #     return WrappedGitHubAuthenticator.authenticate(self, handler, data)
        user_email = data['username']
        password = data['password']
        print('AUTH_URL',os.environ['AUTH_URL'])
        url = os.environ['AUTH_URL'] + "?user_email=" + user_email + "&token=" + password
        req = HTTPRequest(
            url,
            method="GET",
            headers={"Host": "localhost:8081",
                     "Pragma":"no-cache",
                     "Cache-Control":"no-cache",
                     "Content-Type":"application/json;charset=UTF-8",
                     "Origin":"localhost:8081",
                     "Sec-Fetch-Site":"same-site",
                     "Sec-Fetch-Mode":"cors",
                     "Sec-Fetch-Dest":"empty",
                     "Referer":"localhost:8081",
                     "Accept-Encoding":"gzip,deflate,br",
                     "Accept-Language":"zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,zh-HK;q=0.6"
                     }
        )
        resp_data = await self.fetch(req)
        print(resp_data)
        if resp_data.get("message") and resp_data["message"] == "success":
            print(resp_data["data"]["id"])
            print('yes!!!')
            return resp_data["data"]["id"]
        else:
            raise ValueError("Authenticate Failed")

class ExtUserAPIHandler(APIHandler):
    def get(self):

        data = [
            self.user_model(u, include_servers=True, include_state=True)
            for u in self.db.query(orm.User)
        ]
        self.write(json.dumps(data))
    async def post(self):
        data = self.get_json_body()
        print(data)
        user = self.find_user(data['name'])
        if user is not None:
            raise web.HTTPError(409, "User %s already exists" % data['name'])
        user = self.user_from_username(data['name'])
        if data:
            self._check_user_model(data)
            if 'admin' in data:
                user.admin = data['admin']
                self.db.commit()
        try:
            await maybe_future(self.authenticator.add_user(user))
        except Exception:
            self.log.error("Failed to create user: %s" % data['name'], exc_info=True)
            # remove from registry
            self.users.delete(user)
            raise web.HTTPError(400, "Failed to create user: %s" % data['name'])

        self.write(json.dumps(self.user_model(user)))
        self.set_status(201)

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

c.JupyterHub.extra_handlers = [("/external/users", ExtUserAPIHandler),("/external/login",ExtloginHandler)]

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
# def pre_spawn_hook(spawner):
#     username = spawner.user.name
#     # if not username in allowed_users:
#     #     raise ValueError('User Not Registed ')
#     # else:
#     #     dir = Path('/home/' + username + '/examples')
#     #     if not dir.exists():
#     #         subprocess.check_call(['cp', '-r', '/srv/ipython/examples', '/home/' + username + '/examples'])
#     #         subprocess.check_call(['chown','-R',username,'/home/' + username + '/examples'])
#     dir = Path('/home/' + username + '/examples')
#     if not dir.exists():
#         subprocess.check_call(['cp', '-r', '/srv/ipython/examples', '/home/' + username + '/examples'])
#         subprocess.check_call(['chown','-R',username,'/home/' + username + '/examples'])
c.JupyterHub.template_paths = [os.environ['OAUTHENTICATOR_DIR'] + '/templates']
c.JupyterHub.tornado_settings = {
    'headers': {
        'Access-Control-Allow-Origin':'*',
        'Content-Security-Policy': "frame-ancestors https://localhost:8090",
    },
    'cookie_options': {'samesite':'None'}
}
c.NotebookApp.tornado_settings = {
  'headers': {
      'Content-Security-Policy': "frame-ancestors https://localhost:8090"
   }
}
c.Spawner.http_timeout = 100
c.Spawner.args = ['''--NotebookApp.tornado_settings={
  'headers':{
    'Content-Security-Policy': "frame-ancestors https://localhost:8090",
  },
  'cookie_options': {'samesite':'None','Secure':True},
  'xsrf_cookie_kwargs': {'samesite':'None','Secure':True}
}''']
c.Spawner.cmd = ["jupyter-labhub"]
c.MyAuthenticator.create_system_users = True
c.MyAuthenticator.client_id = os.environ['CLIENT_ID']
c.MyAuthenticator.client_secret = os.environ['CLIENT_SECRET']
# c.JupyterHub.tornado_settings = {'cookie_options': {'samesite': 'None'}}
# ssl config
ssl = join(root, 'ssl')
keyfile = join(ssl, 'jupyter.key')
certfile = join(ssl, 'jupyter.crt')
if os.path.exists(keyfile):
    c.JupyterHub.ssl_key = keyfile
if os.path.exists(certfile):
    c.JupyterHub.ssl_cert = certfile
