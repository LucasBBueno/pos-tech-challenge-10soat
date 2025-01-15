FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main ./cmd/main.go
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g  ./cmd/main.go

CMD ["./main"]
