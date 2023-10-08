FROM golang:1.20-buster AS builder
RUN apt-get update && \
	apt-get install -y ca-certificates openssl && \
	update-ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o main

FROM debian:buster-slim

WORKDIR /app
COPY --from=builder /app/main ./
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
# front end
COPY www ./www
# GCS service account
COPY gcs.json ./

EXPOSE 8090

CMD ["./main"]
