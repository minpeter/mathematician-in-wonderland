# Step 2: Builder, with cross compile
FROM --platform=$BUILDPLATFORM golang:latest AS builder
COPY . /app
WORKDIR /app
RUN go mod download

ARG TARGETOS TARGETARCH 
ENV CGO_ENABLED=0
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/app .

# GOPATH for scratch images is /
FROM alpine:latest
COPY --from=builder /bin/app /app

RUN apk add --no-cache socat websocketd

EXPOSE 1337
EXPOSE 1338

CMD socat TCP-LISTEN:1337,reuseaddr,fork EXEC:"/app" & \
    websocketd --port=1338 /app