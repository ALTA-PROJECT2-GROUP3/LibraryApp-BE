FROM golang:alpine as builder

WORKDIR /app/

COPY . .

RUN go build -o /app/main /app/main.go

FROM alpine

WORKDIR /app/

COPY --from=builder /app/ .

ENTRYPOINT ["/app/main"