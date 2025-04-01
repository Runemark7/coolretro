# Build Stage
FROM golang:alpine AS builder
ENV CGO_ENABLED=1 GOOS=linux GOCACHE=/root/.cache/go-build
WORKDIR /app
RUN apk add --no-cache gcc musl-dev sqlite-dev
# Copy dependency files first to leverage caching
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
# Copy the rest of the application source code
COPY . .
# Cache build artifacts during compilation
RUN --mount=type=cache,target=/root/.cache/go-build go build -o main .

# Runtime Stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x ./main
ENTRYPOINT ["./main"]
