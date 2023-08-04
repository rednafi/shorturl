.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./tests/...


.PHONY: build
build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) -v


.PHONY: run
run:
	@echo "Running..."
	@go run main.go


.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean
	@rm -f $(BINARY_NAME)


.PHONY: lint
lint:
	@echo "Linting..."
	@go vet ./...
	@go fmt ./...
	@go fix ./...
