FROM golang:1.20.7-alpine as builder

WORKDIR /app

RUN apk add git

ENV GO111MODULE on

# ENV CGO_ENABLED=1
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go.mod .
COPY go.sum .

COPY . .

RUN apk update

RUN apk add  make

RUN make ci && make build

FROM alpine:latest as release

RUN apk add --no-cache --update ca-certificates

COPY --from=builder /app/main /app/cmd/

RUN chmod +x /app/cmd/main

WORKDIR /app

EXPOSE 80

CMD ["cmd/main"]
