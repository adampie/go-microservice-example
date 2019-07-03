FROM golang:alpine AS builder
COPY go.mod /build/
COPY go.sum /build/
WORKDIR /build/
RUN apk update \
    && apk add --no-cache git ca-certificates \
    && update-ca-certificates
RUN adduser -D -g '' appuser
RUN go mod download
RUN go mod verify
ADD . /build/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app cmd/api/main.go
RUN chmod +x /app

FROM scratch
COPY --from=builder /app ./
COPY --from=builder /build/configs/migrations /migrations/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER appuser
EXPOSE 8080
ENTRYPOINT ["./app"]
