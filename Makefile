RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)
test:
	go test -v $(RUN_ARGS)

run:
	go run main.go

build:
	go build -o poupaniquel main.go

install:
	curl https://glide.sh/get | sh
	glide install
	go install

packages:
	time xgo --targets=windows/amd64,darwin/amd64,linux/amd64 $(GOPATH)/src/github.com/gustavohenrique/poupaniquel

static:
	scripts/static.sh
