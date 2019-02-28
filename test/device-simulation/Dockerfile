FROM node:lts

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . ./

EXPOSE 9999
ENTRYPOINT [ "node", "/app/index.js" ]