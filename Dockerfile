# make nc server

FROM golang:alpine

RUN apk add --no-cache socat

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o server .

EXPOSE 5555

CMD socat TCP-LISTEN:5555,reuseaddr,fork EXEC:"./server"