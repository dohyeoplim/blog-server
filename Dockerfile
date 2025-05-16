FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o blog-server .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/blog-server .

EXPOSE 8080

ENV PORT=8080

CMD ["./blog-server"]
