
test:
	go test -v -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./...