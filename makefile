test-all:
	go list -f '{{.Dir}}' -m | xargs go test -v -race -cover -coverprofile=profile.out
test:
	go test -v -race -cover -coverprofile=coverage.out \
		./common/... \
	&& go tool cover -html=coverage.out -o coverage.html
clean:
	
all:

inspect:
	echo 'file://wsl.localhost/Ubuntu/home/lavantien/dev/personal/lavantien/simple-commerce/coverage.html'

.PHONY: test-all test clean inspect all
