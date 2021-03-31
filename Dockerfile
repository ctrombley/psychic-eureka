FROM golang:1.16.2-alpine3.13

# Install make
RUN apk add --update-cache \
 make \
 git

# Install goreleaser
RUN go install github.com/goreleaser/goreleaser@v0.161.1

# Add source
ADD . ./src
WORKDIR ./src

# Build the binary
RUN make

ENTRYPOINT ["/go/src/dist/product-location-service_linux_amd64/product-location-service"]
