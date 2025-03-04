# build context at repo root: docker build -f Dockerfile ../..
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY config config
COPY . .
ARG TAG
ARG GIT_COMMIT
WORKDIR /app
RUN make build 

# stage 2: production image
FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/simple_cinema /simple_cinema

RUN chmod +x /simple_cinema
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Run the web service on container startup.
CMD ["/simple_cinema"]