FROM golang:1.19-alpine as build-env

COPY . /build

WORKDIR /build/app

RUN  go build -o app

FROM alpine:3.16

COPY --from=build-env /build/app/app /

CMD ["/app"]