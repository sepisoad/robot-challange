FROM golang:1.18 AS build

ADD . /src

WORKDIR /src
RUN --mount=type=cache,target=/go make dep
RUN --mount=type=cache,target=/go make build-api

FROM alpine:latest

WORKDIR /
COPY --from=build /src/bin/robots-api .
COPY ../web/dist/* .

ENTRYPOINT [ "./robots-api", "start" ]
