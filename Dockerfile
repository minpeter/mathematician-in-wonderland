# make nc server

FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o server .

FROM alpine

COPY --from=builder /app/server /app

RUN apk add --no-cache socat

EXPOSE 5555

CMD socat TCP-LISTEN:5555,reuseaddr,fork EXEC:"/app"