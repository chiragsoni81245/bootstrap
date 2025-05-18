build:
	@go build -o bin/bootstrap ./main.go

run: build
	@./bin/bootstrap create $(ARGS)

test:
	@go test ./... -v --race
