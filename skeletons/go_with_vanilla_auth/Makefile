build:
	@go build -o bin/{{ .project.folder_name }} ./main.go

run: build
	@./bin/{{ .project.folder_name }} start $(ARGS)

test:
	@go test ./... -v --race

# Database migrations
migrate-up: build
	@./bin/{{ .project.folder_name }} migrate up $(ARGS)

migrate-down: build
	@./bin/{{ .project.folder_name }} migrate down $(ARGS)
