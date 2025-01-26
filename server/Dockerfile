FROM golang:alpine as builder

WORKDIR /app

COPY go.mod server.go ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o compiler-server

FROM scratch

COPY --from=builder /app/compiler-server /

EXPOSE 3000

CMD ["/compiler-server"]
