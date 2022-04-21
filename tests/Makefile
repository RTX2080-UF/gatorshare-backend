server: 
	go run main.go

changedir:
	cd  ./tests && If (Test-Path -Path "../db/share-v1.0-test.db" ) { del "../db/share-v1.0-test.db" }

test: 
	go clean -testcache && go test -v