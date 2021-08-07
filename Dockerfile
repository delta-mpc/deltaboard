
FROM jupyterhub/jupyterhub

RUN  apt-get update -y 
RUN  apt-get install -y gettext \
                        sqlite3 \
                        golang \
                        nodejs
RUN sqlite3 delta_dashboard.db
ENV GOROOT /usr/lib/go
ENV PATH $PATH:/usr/lib/go/bin
ENV GOPATH /root/go
ENV PATH $GOPATH/bin/:$PATH
# Install oauthenticator from git
RUN python3 -m pip install oauthenticator && \
    python3 -m pip install jupyterlab -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host=mirrors.aliyun.com && \
    python3 -m pip install notebook -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host=mirrors.aliyun.com

RUN mkdir /app && \
    mkdir /app/web
WORKDIR /app/web

ADD front/package.json package.json
RUN npm install

ADD ./server /app/
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn && go build -ldflags "-w -s" -o main


#don't touch codes above this line or you'll waste ton's of time in installing things


# Create oauthenticator directory and put necessary files in it
RUN mkdir /srv/oauthenticator && \
    mkdir /srv/ipython && \
    mkdir /srv/ipython/examples && \
    mkdir /root/.jupyter
ADD ./jupyter/helloworld.ipynb /srv/ipython/examples/helloworld.ipynb
ADD ./jupyter/jupyter_notebook_config.py /root/.jupyter/jupyter_notebook_config.py
WORKDIR /srv/oauthenticator
ENV OAUTHENTICATOR_DIR /srv/oauthenticator
ADD ./jupyter/templates templates
ADD ./jupyter/notebook_templates/page.html /usr/local/lib/python3.8/dist-packages/notebook/templates/page.html
ADD ./jupyter/base.py /usr/local/lib/python3.8/dist-packages/jupyterhub/handlers/base.py
ADD ./jupyter/pages.py /usr/local/lib/python3.8/dist-packages/jupyterhub/handlers/pages.py
ADD ./jupyter/login.py /usr/local/lib/python3.8/dist-packages/jupyterhub/handlers/login.py
ADD ./jupyter/custom.js /usr/local/lib/python3.8/dist-packages/notebook/static/custom/custom.js
ADD ./jupyter/ssl /srv/oauthenticator/ssl
RUN chmod 700 /srv/oauthenticator


WORKDIR /app

ADD ./run.sh /app
COPY ./server/static /app/static
COPY ./server/config/config.tmpl config/config.tmpl
ADD ./run_jupyter.sh /app
ADD ./jupyter/jupyterhub_config.py jupyterhub_config.py
ADD ./jupyter/ssl /app/ssl

WORKDIR /app/web

ADD front/public public
ADD front/src src
ADD front/.env.development .env.development
ADD front/babel.config.js babel.config.js
ADD front/package.json package.json
ADD front/proxy.config.json proxy.config.json
ADD front/vue.config.js vue.config.js
WORKDIR /app
ADD run_node.sh run_node.sh 
CMD ["sh", "./run.sh"]
