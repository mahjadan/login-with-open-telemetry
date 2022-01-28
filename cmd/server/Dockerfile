FROM golang:1.17-alpine3.13 as build_base

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -installsuffix cgo -mod=vendor -o ./out/go-server ./cmd/server

FROM alpine
RUN apk add ca-certificates

COPY --from=build_base /app/out/go-server /app/go-server

ENTRYPOINT ["/app/go-server"]
