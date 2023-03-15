
FROM golang:1.16 as builder

ADD ./server /application/

WORKDIR /application
RUN go env -w GOPROXY=https://goproxy.cn && go build -ldflags "-w -s" -o main

FROM node:12-alpine3.14 as webbuilder

RUN mkdir -p /application/web

WORKDIR /application/web
ADD web/package.json package.json

RUN npm install 
RUN npm audit fix

ADD web/config.json config.json
ADD web/public public
ADD web/src src
ADD web/.env .env
ADD web/babel.config.js babel.config.js
ADD web/vue.config.js vue.config.js

RUN npm run build 

FROM jupyterhub/jupyterhub:1.4.2
# Create oauthenticator directory and put necessary files in it

WORKDIR /application
COPY requirements.txt /application/requirements.txt
COPY --from=builder /application/main /application/main
COPY --from=webbuilder /application/web/dist /application/web

RUN apt-get update -y && \
    apt-get install -y gettext sqlite3 nginx && \
    pip install --no-cache-dir -r requirements.txt && \
    mkdir /srv/oauthenticator && \
    mkdir /srv/ipython && \
    mkdir /srv/ipython/examples && \
    mkdir /root/.jupyter
COPY ./jupyter/examples  /srv/ipython/examples
ADD ./jupyter/jupyter_jupyterlab_server_config.py /application/jupyter_jupyterlab_server_config.py
WORKDIR /srv/oauthenticator
ENV OAUTHENTICATOR_DIR /srv/oauthenticator
ADD ./jupyter/templates templates
ADD ./jupyter/notebook_templates/page.html /usr/local/lib/python3.8/dist-packages/notebook/templates/page.html
ADD ./jupyter/base.py /usr/local/lib/python3.8/dist-packages/jupyterhub/handlers/base.py
ADD ./jupyter/pages.py /usr/local/lib/python3.8/dist-packages/jupyterhub/handlers/pages.py
ADD ./jupyter/login.py /usr/local/lib/python3.8/dist-packages/jupyterhub/handlers/login.py
ADD ./jupyter/custom.js /usr/local/lib/python3.8/dist-packages/notebook/static/custom/custom.js

RUN chmod 700 /srv/oauthenticator

WORKDIR /application

ADD ./run.sh /application
COPY ./server/static /application/static
COPY ./server/config/config.tmpl /application/config/config.tmpl
ADD ./jupyter/jupyterhub_config.py jupyterhub_config.py
ADD ./jupyter/ssl /application/ssl

ADD ./default.conf /etc/nginx/conf.d/default.conf

WORKDIR /application/web


WORKDIR /app
RUN mkdir /app/app_config
WORKDIR /application
ADD run_node.sh run_node.sh 
ADD gen_config.py gen_config.py
ADD gen_web_config.py gen_web_config.py
EXPOSE 8090

ENTRYPOINT ["sh","./run.sh"]