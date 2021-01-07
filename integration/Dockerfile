FROM alpine:3.12

RUN apk -U --no-cache upgrade && \
  apk -U --no-cache add \
    bash \
    curl \
    jq \
    openjdk11 \
    openssl \
    vim \
    wget

RUN mkdir /tests
WORKDIR /tests
VOLUME /tests

EXPOSE 25565
EXPOSE 25575

ENTRYPOINT ["/tests/suite.sh"]
