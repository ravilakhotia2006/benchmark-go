build: deps
    go build -o benchmark main.go routes.go

run:
    go run main.go