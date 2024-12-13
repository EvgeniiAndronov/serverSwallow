FROM golang:1.23.3

WORKDIR /cmd/appServer

COPY go.* ./

RUN go mod download

COPY . .
RUN go build -o main cmd/appServer/main.go

EXPOSE 8080

CMD ["./main"]

