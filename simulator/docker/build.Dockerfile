FROM golang:1.18 AS build

ADD . /src

WORKDIR /src
RUN --mount=type=cache,target=/go make dep
RUN --mount=type=cache,target=/go make build-simulator

FROM alpine:latest

WORKDIR /
COPY --from=build /src/bin/robots-simulator .

ENTRYPOINT [ "./robots-simulator", "start" ]
