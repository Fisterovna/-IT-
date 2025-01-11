echo "Building the project..."
go build -o calc_service ./cmd/calc_service

echo "Running tests..."
go test ./...

echo "Build and tests complete."