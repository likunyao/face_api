ARG GO_VERSION=1.14.2

FROM golang:${GO_VERSION}-alpine AS builder

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

RUN mkdir -p /face_api
WORKDIR /face_api

COPY . .
RUN go mod download
RUN go build -o ./face_app ./main.go

FROM alpine:latest

RUN mkdir -p /face_api
WORKDIR /face_api
COPY --from=builder /face_api/face_app .
COPY --from=builder /face_api/config.ini .

EXPOSE 8000

ENTRYPOINT ["./face_app"]