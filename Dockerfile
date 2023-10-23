FROM golang:1.19.5-alpine as builder

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR '/app'

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod download
RUN go mod vendor
RUN go mod verify


# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN go build -o kreditplus
EXPOSE 9000

CMD ["./kreditplus"]