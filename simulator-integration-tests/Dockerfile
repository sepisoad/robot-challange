FROM golang:1.18 AS build

ADD . /src

WORKDIR /src
RUN --mount=type=cache,target=/go make dep

WORKDIR /src/simulator-integration-tests

ENTRYPOINT ["go", "test", "./..."]
