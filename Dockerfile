FROM golang:1.14.4-alpine as build-env
COPY ./* /go/src/audit-example/
WORKDIR /go/src/audit-example/
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/audit-example/app /app/
COPY --from=build-env /go/src/audit-example/webhook-server-tls* /app/
CMD ["./app"]