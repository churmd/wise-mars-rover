format:
	go fmt ./...

clean_tests:
	go clean -testcache

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out