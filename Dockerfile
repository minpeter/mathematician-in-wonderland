# make nc server

FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o server .

FROM alpine

RUN apk add --no-cache socat

COPY --from=builder /app/server /app

ENV PORT 1337
EXPOSE $PORT

CMD socat TCP-LISTEN:$PORT,reuseaddr,fork EXEC:"/app"