FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
  apt-get install -y \
    curl \
    jq \
    openjdk-16-jre \
    unzip \
    vim \
    wget \
    && \
  apt-get clean &&\
  rm -rf /var/lib/apt/lists/*

RUN mkdir /tests
WORKDIR /tests
VOLUME /tests

EXPOSE 25565
EXPOSE 25575

ENTRYPOINT ["/tests/suite.sh"]
