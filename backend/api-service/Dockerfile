FROM --platform=linux/amd64 golang:1.24.2-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /docker-gs-ping

FROM alpine:3.20
WORKDIR /
COPY --from=builder /docker-gs-ping /docker-gs-ping
EXPOSE 8080
CMD ["/docker-gs-ping"]
