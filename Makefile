TEST?=

default: test

test:
	go vet `glide nv`
	go test -v `glide nv`

build:
	gox -osarch="darwin/amd64 linux/amd64 windows/amd64"

build-ci:
	go get github.com/mitchellh/gox
	sudo chown -R ${USER}: /usr/local/go
	$(MAKE) build
