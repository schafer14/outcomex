# pull official base image
FROM node:13.12.0-alpine as reactor

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY front-end/package.json ./
COPY front-end/package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

# add app
COPY front-end ./

# start app
RUN npm run-script build

# Go APP
FROM golang:alpine

RUN mkdir /app 

ADD . /app/

WORKDIR /app 

COPY --from=reactor /app/build ./front-end/build

RUN go build -o api ./cmd/api

RUN adduser -S -D -H -h /app appuser

USER appuser

ENTRYPOINT ["./api"]

