FROM node:12.18.2

RUN mkdir /app
ADD public /app/public
ADD src /app/src
ADD .env /app/.env
ADD .env.development /app/.env.development
ADD babel.config.js /app/babel.config.js
ADD package.json /app/package.json
ADD package-lock.json /app/package-lock.json
ADD proxy.config.json /app/proxy.config.json
ADD vue.config.js /app/vue.config.js

WORKDIR /app
RUN npm install

CMD npm run serve

