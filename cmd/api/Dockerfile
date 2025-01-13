FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# For some reason despite providing the GOOS and GOARCH output is ALWAYS an archive file
# From what I've found this is likely an issue deriving from using WSL, I have not been able to find a fix
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o /go-quiz-api ./cmd/api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go-quiz-api .
EXPOSE 8080

RUN chmod +x go-quiz-api
CMD ["./go-quiz-api"]