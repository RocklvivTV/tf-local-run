build:
	go mod vendor
	go build -o bin/tf_local main.go

run:
	go run main.go