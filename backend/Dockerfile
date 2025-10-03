FROM golang:1.25.1-alpine AS builder
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend ./backend/cmd/backend
RUN chmod +x /backend

FROM alpine:3.21
COPY --from=builder /backend /backend
ENTRYPOINT ["/backend"]