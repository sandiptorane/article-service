## Generate mocks for all the interfaces
gen-mock:
	rm -rf mocks
	mockery --all --keeptree --case underscore --with-expecter --exported

# test case coverage
unit-test:
	go test ./... -v -cover -coverprofile=coverage.txt

# check coverage
test-cover:
	go tool cover -func coverage.txt
