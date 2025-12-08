FROM golang:1.24 AS builder 

WORKDIR /app 
COPY go.mod go.sum ./
RUN go mod download

COPY ./src ./src

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./src/cmd

FROM alpine:3.16

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 5000

CMD ["./app"]
