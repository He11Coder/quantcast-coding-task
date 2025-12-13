#Build the project into a executable binary called 'cookie-tool'
PHONY.: build
build:
	go build -o cookie-tool ./cmd

#Run all tests, print coverage statistics and save it to 'tmp_cover.out'
PHONY.: test
test:
	go test -cover -coverprofile=tmp_cover.out ./pkg/logparse ./pkg/stats ./pkg/argparse ./cmd

#Print total coverage statistics (main() function's coverage is not considered)
PHONY.: coverage
coverage: test
	grep -v "most_active_cookie.go" tmp_cover.out > cover.out && go tool cover -func=cover.out