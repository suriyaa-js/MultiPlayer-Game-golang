# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .
RUN go mod download

RUN go build -o ./userapi

FROM debian:12
# FROM gcr.io/distroless/base-debian12  // For smaller image size

WORKDIR /app
COPY --from=builder /build/userapi ./userapi
COPY .env /app

# Add a new file to the final image
COPY read.txt /app/read.txt

# Run
CMD ["/app/userapi"]