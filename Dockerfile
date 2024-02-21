# build the app
FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

# run the app
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

COPY app.env .

EXPOSE 8080

CMD ["/app/main"]