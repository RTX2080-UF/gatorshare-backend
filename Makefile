server: 
	go run main.go

test: 
	rm share-v1.0-test.db && go clean -testcache && go test -v .\tests