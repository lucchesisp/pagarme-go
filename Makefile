dependencies:
	go mod download

quality:
	go install golang.org/x/lint/golint@latest
	go vet ./...
	golint -set_exit_status ./...

test:
	go test ./pagarme/... --cover -test.paniconexit0