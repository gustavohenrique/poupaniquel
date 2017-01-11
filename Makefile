RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)
test:
	go test -v $(RUN_ARGS)

tests:
	go test -v ./api/...

run:
	go run main.go

build:
	export GIN_MODE=release
	go build -o poupaniquel main.go

install:
	curl https://glide.sh/get | sh
	glide install
	go install

packages:
	time xgo -go develop --targets=darwin/amd64,linux/amd64,windows/amd64 $(GOPATH)/src/github.com/gustavohenrique/poupaniquel

static:
	scripts/static.sh
