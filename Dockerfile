FROM golang:1.14.3-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080
RUN go build -o benchmark main.go routes.go

CMD ["./benchmark"]
