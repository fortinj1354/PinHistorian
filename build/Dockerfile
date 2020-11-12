FROM alpine:3.12.1
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY PinHistorian /go/bin

ENV GIN_MODE=release
ENV PORT=8080
ENV DATABASE_NAME=/config/PinHistorian.sqlite

EXPOSE 8080
VOLUME /config
ENTRYPOINT /go/bin/PinHistorian --port 80