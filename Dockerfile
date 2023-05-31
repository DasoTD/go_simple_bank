# build stage
FROM golang:1.19.9-alpine3.18 AS builder

WORKDIR /app

COPY . .
RUN go build -o main main.go

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify


# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .



EXPOSE 8080

CMD ["/app/main"]
