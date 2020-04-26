default: fmt test coverage clean

clean:
	rm c.out

fmt:
	go fmt ./...

test:
	go test ./... -coverprofile=c.out

coverage: test
	go tool cover -html=c.out
