build:
	@echo "Building the binary..."
	@go build -o bin/go-crud ./cmd	

run: build
	@echo "Running the application..."
	@./bin/go-crud
