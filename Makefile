format:
	go fmt ./...

clean_tests:
	go clean -testcache

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

docker_build:
	docker build -t mars-rover .

docker_run:
	docker run -it mars-rover