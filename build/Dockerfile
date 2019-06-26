FROM golang:alpine AS builder
ADD . /build/
WORKDIR /build/
RUN apk update \
    && apk add --no-cache git ca-certificates \
    && update-ca-certificates
RUN adduser -D -g '' appuser
RUN go mod download \
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app .

FROM scratch
COPY --from=builder /app ./
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER appuser
EXPOSE 8080
ENTRYPOINT ["./app"]
