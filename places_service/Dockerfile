FROM golang:1.17.0-alpine AS builder

# RUN apk --no-cache add gcc g++ make git
#RUN apk add upx

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

# Compress the build
#RUN upx ./main

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM scratch
MAINTAINER heri@heri.co

COPY --from=builder /dist/main /

EXPOSE 3000

# Command to run
ENTRYPOINT ["/main"]