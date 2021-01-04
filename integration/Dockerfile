FROM openjdk:8u212-jre-alpine

RUN apk -U --no-cache upgrade && \
  apk -U --no-cache add \
    bash \
    curl \
    jq \
    openssl \
    vim \
    wget

RUN mkdir /server /server-defs
WORKDIR /server

ENV USER=minecraft
ENV GROUP=minecraft

RUN addgroup -S ${USER} && adduser -S ${USER} -G ${USER} && \
  chown -R ${USER}:${USER} /server
USER ${USER}

ARG MCSM_BIN=build/mcsm-linux-amd64
COPY ${MCSM_BIN} /usr/local/bin/mcsm

VOLUME /server-defs

EXPOSE 25565
EXPOSE 25575

ENTRYPOINT ["/bin/bash"]
