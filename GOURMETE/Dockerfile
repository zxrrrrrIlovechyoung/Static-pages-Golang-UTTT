FROM golang:1.26.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /gourmete ./BACKEND/CMD

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /gourmete /app/gourmete
COPY FRONTEND/public /app/FRONTEND/public

EXPOSE 8080

CMD ["/app/gourmete"]
