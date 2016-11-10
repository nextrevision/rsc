TEST?=

default: test

test:
	go vet `glide nv`
	go test -v `glide nv`

build:
	gox -os="darwin linux" -osarch="amd64"

build-ci:
	go get github.com/mitchellh/gox
	sudo chown -R ${USER}: /usr/local/go
	$(MAKE) build
