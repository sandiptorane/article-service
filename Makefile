## Generate mocks for all the interfaces
gen-mock:
	rm -rf mocks
	mockery --all --keeptree --case underscore --with-expecter --exported

# test case coverage
test-cover:
	go test ./... -v -coverpkg=./... -coverprofile=coverage.out
	go tool cover -html=coverage.out