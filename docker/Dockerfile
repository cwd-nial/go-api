#################################
# STEP 1 build executable binary
#################################
FROM golang:1.12-alpine AS builder

# Set env values
ENV PACKAGE=github.com/cwd-nial/go-api
ENV GOOS=linux
ENV LISTEN=0.0.0.0:8080
ENV CGO_ENABLED=0

WORKDIR /go/src/${PACKAGE}
COPY . .

# Build the binary
RUN go build -o /go/bin/api

#############################
# STEP 2 build a small image
#############################
FROM scratch

COPY --from=builder /go/bin/api /go/bin/api

# Run the api binary.
ENTRYPOINT ["/go/bin/api"]