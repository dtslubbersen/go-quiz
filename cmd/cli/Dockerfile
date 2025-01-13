FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-quiz-cli ./cmd/cli

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go-quiz-cli .
CMD ["./go-quiz-cli"]