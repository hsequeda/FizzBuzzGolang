
build_docker:
	docker build -t "fizzbuzz-golang" "." -f "./Dockerfile"

linter:
	golint ./...

test:
	go test ./app
