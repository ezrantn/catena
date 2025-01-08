test:
	@go test -v .

bench:
	@go test -bench . -benchmem

test-cov:
	@go test -coverprofile=catena.out

cov: test-cov
	@go tool cover -html=catena.out

fmt:
	@go fmt .