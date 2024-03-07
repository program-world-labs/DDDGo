# Step 1: Modules caching
FROM golang:1.22.0-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Intermediate
FROM alpine:3.14 as intermediate
RUN apk update && \
    apk add --no-cache wget=1.21.1-r1 && \
    wget --progress=dot:giga https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz && \
    tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.6.1.tar.gz && \
    rm dockerize-linux-amd64-v0.6.1.tar.gz

# Step 3: Tests
FROM golang:1.22.0-alpine
COPY --from=modules /go/pkg /go/pkg
COPY --from=intermediate /usr/local/bin/dockerize /dockerize
COPY . /app
WORKDIR /app

RUN go env -w CGO_ENABLED=0
RUN go env -w GOOS=linux
RUN go env -w GOARCH=amd64

CMD ["go", "test", "-v", "./integration-test/..."]