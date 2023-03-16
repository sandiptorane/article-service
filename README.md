# article service

## Prerequisites

1. docker 
2. mockery to mock interface functions github.com/vektra/mockery/v2

## tools used
1. mockery to mock interface functions github.com/vektra/mockery/v2
2. github.com/DATA-DOG/go-sqlmock to mock sql query

## how to
1. to create mocks run `make gen-mock`
2. to run unit test run `make test-cover`
3. to deploy start run for linux `sh start.sh` for mac `bash start.sh`