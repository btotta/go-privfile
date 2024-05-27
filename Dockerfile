FROM golang:1.22

WORKDIR /go/src/app

COPY . .

# Instalar swag para gerar documentação da API
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Gerar documentação da API com swag
RUN swag init -g internal/server/routes.go -o docs

RUN go get -d -v ./...

RUN go build -o main ./cmd/api

EXPOSE 8080

CMD ["./main"]
