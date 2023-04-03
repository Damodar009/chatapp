FROM golang:1.16-alpine

# Required because go requires gcc to build
RUN apk update \
    && apk add build-base

RUN apk add inotify-tools

RUN apk add --no-cache ffmpeg

COPY . /clean_web

ARG VERSION="4.13.0"

WORKDIR /clean_web

RUN go mod tidy



CMD go run main.go