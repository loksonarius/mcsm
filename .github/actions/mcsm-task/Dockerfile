FROM golang:1.15-alpine3.12

RUN apk -U --no-cache upgrade && \
  apk -U --no-cache add \
    bash \
    curl \
    gcc \
    git \
    jq \
    musl-dev \
    openjdk11 \
    openssl \
    vim \
    wget

RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | \
  bash -s -- --to /usr/local/bin

ENTRYPOINT ["/usr/local/bin/just"]
