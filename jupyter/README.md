# OAuthenticator

github https://github.com/jackeyli/jupyterhub-github-oauth

# ADD your own ssl key in the ssl folder
cert file :ssl/jupyter.crt
key file :ssl/jupyter.key

## build

Build the container with:

    docker build -t jupyterhub-oauth .

### ssl

To run the server on HTTPS, put your ssl key and cert in ssl/ssl.key and
ssl/ssl.cert.

## run

Add your oauth client id, client secret, and callback URL to the `env file`.
Once you have built the container, you can run it with:

    docker run -it -p 8000:8000 --env-file=env jupyterhub-oauth

Which will run the Jupyter server.
