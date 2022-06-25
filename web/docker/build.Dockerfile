FROM joseluisq/static-web-server:2 AS build

ADD . /dist

WORKDIR /dist

ENTRYPOINT [ "./robots-simulator", "start" ]
