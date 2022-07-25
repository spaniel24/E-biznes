FROM node:18.6-alpine3.16

WORKDIR /app/e-biznes_react

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .


CMD ["npm", "start"]