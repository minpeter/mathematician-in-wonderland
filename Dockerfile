# Step 1: Modules caching
FROM --platform=$BUILDPLATFORM golang:latest as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder, with cross compile
FROM --platform=$BUILDPLATFORM golang:latest AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
ARG TARGETOS TARGETARCH 
ENV CGO_ENABLED=0
WORKDIR /app
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/app .

# GOPATH for scratch images is /
FROM scratch
COPY --from=builder /bin/app /app

RUN apk add --no-cache socat websocketd

EXPOSE 1337
EXPOSE 1338

CMD socat TCP-LISTEN:1337,reuseaddr,fork EXEC:"/app" & \
    websocketd --port=1338 /app