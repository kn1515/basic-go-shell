# syntax=docker/dockerfile:1
FROM golang:1.22-bullseye AS build

WORKDIR /app

RUN apt update -y && apt upgrade -y \
    && apt install -y \
    iproute2 \
    iputils-ping \
    arping \
    netcat \
    sudo \
    git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*