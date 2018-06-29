FROM golang

ARG app_env
ENV APP_ENV $app_env

RUN apt-get update
RUN apt-get install -y --no-install-recommends apt-utils git


COPY ./ /go/src/github.com/raulgonz/almanac/server
WORKDIR /go/src/github.com/raulgonz/almanac/server

RUN go get ./
RUN go build