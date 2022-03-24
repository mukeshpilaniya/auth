# Build stage
FROM  golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN  go mod tidy -go=1.16 && go mod tidy -go=1.17
RUN go build -o main cmd/api/*

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8081
CMD ["/app/main"]