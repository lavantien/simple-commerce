test-all:
	go list -f '{{.Dir}}' -m | xargs go test -v -race -cover -coverprofile=profile.out \
	&& go tool cover -html=coverage.out -o coverage.html \
	&& ./generate-coverage-badge.sh
test:
	go test -v -race -cover -coverprofile=coverage.out \
		./common/... \
	&& go tool cover -html=coverage.out -o coverage.html \
	&& ./generate-coverage-badge.sh
clean:
	
all:

inspect:
	echo 'file://wsl.localhost/Ubuntu/home/lavantien/dev/personal/lavantien/simple-commerce/coverage.html'

.PHONY: test-all test clean inspect all
