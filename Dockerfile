# Build nebula
FROM golang:1.19 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(cat version)'" -o nebula cmd/nebula/*

# Create lightweight container to run nebula
FROM alpine:latest

# Create user ot
RUN adduser -D -H nebula
WORKDIR /home/nebula
RUN mkdir .config && chown nebula:nebula .config
USER nebula

COPY --from=builder /build/nebula /usr/local/bin/nebula

CMD nebula
