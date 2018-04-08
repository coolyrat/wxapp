FROM golang:1.10-alpine as build

WORKDIR $GOPATH/src/github.com/coolyrat/wxapp
ADD . .
RUN go build -o wxapp

FROM alpine:latest

WORKDIR /app
COPY --from=build /go/src/github.com/coolyrat/wxapp/wxapp .
EXPOSE 8080
ENTRYPOINT ./wxapp
