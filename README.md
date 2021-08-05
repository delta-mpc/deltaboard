# delta-board-server

## Run with provided mysql compose up
docker-compose up -d


## Run in your own environment

replace the docker-compose.yml with code as below
```
version: '3'
services:
   api:
      build: ./server
      container_name: go-service-api
      ports:
         - 8080:8080
      expose:
         - 8080
      depends_on:
         - migrate
      environment:
         WORK: server
         ENV: config
         CONNECTOR: "Your own mysql connection}"
         FRONT_END_DOMAIN: "Your dashboard front end domain"
      volumes:
         - ./server/config:/app/config
   migrate:
      build: ./server
      container_name: go-service-migrate-a
      environment:
         WORK: migrate
         ENV: config
         CONNECTOR: "Your own mysql connection}"
         FRONT_END_DOMAIN: "Your dashboard front end domain"
      volumes:
         - ./server/config:/app/config
   jupyter:
      build: ./jupyter
      container_name: jupyterhub
      ports:
         - 8000:8000
      links:
         - api
      environment:
         AUTH_URL: http://api:8080/v1/users/auth
      volumes:
         - "You local folder to store user's jupyter data":/home
```