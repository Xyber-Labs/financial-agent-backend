GO=go
 
coverage:
	### Generate coverage report
	$(GO) test -covermode=atomic \
	  -coverpkg=./... \
	  -coverprofile=coverage.out \
	  ./...
	### Export to html
	$(GO) tool cover -html=coverage.out -o coverage.html

swagger:
	swag init -g ./cmd/main.go -d .