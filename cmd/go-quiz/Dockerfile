FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY ./data/ ./data/
COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o /go-quiz ./cmd/go-quiz
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-quiz ./cmd/go-quiz/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/go-quiz .
COPY --from=builder /app/data ./data
EXPOSE 8080

RUN chmod +x go-quiz
ENTRYPOINT ["/app/go-quiz"]






