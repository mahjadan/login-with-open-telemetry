#FROM golang:1.12.0-alpine3.9
FROM golang:1.17-buster AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -installsuffix cgo -mod=vendor  -o ./go-client ./cmd/client

FROM alpine
RUN apk add ca-certificates
COPY --from=builder /app/go-client /app/go-client

CMD ["/app/go-client"]


