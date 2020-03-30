default: coverage

fmt:
	go fmt .

test: fmt
	go test -coverprofile=c.out

coverage: test
	go tool cover -html=c.out
