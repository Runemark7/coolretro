# Build Stage
FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN apk add --no-cache gcc musl-dev sqlite-dev
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Runtime Stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x ./main
ENTRYPOINT ["./main"]
