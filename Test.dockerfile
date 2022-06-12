FROM golang:1.18-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

RUN go build -v -o /app/server github.com/ZNotify/server

FROM scratch

WORKDIR /app

ENV GIN_MODE release
ENV TEST true

EXPOSE 14444

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server ./server

ENTRYPOINT ["./server"]