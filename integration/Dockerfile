FROM openjdk:8u212-jre-alpine

RUN apk -U --no-cache upgrade && \
  apk -U --no-cache add \
    bash \
    curl \
    jq \
    openssl \
    vim \
    wget

RUN mkdir /tests
WORKDIR /tests

ENV USER=minecraft
ENV GROUP=minecraft

RUN addgroup -S ${USER} && adduser -S ${USER} -G ${USER} && \
  chown -R ${USER}:${USER} /tests
USER ${USER}

VOLUME /tests

EXPOSE 25565
EXPOSE 25575

ENTRYPOINT ["/tests/suite.sh"]
