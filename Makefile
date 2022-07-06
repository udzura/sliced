VERSION = $(shell grep 'const VERSION' sliced.go | sed 's/.*"\(.*\).*"/\1/')

.PHONE: build
build:
	docker build -t udzura/sliced-pkg:$(VERSION) .

.PHONY: pkg
pkg: build
	$(eval cid := $(shell docker create udzura/sliced-pkg:$(VERSION)))
	mkdir -p pkg
	docker cp $(cid):'/app/build/sliced-linux-amd64' pkg/sliced-linux-amd64-$(VERSION)
	docker cp $(cid):'/app/build/sliced-linux-arm64' pkg/sliced-linux-arm64-$(VERSION)
	docker cp $(cid):'/app/build/sliced-darwin-arm64' pkg/sliced-darwin-arm64-$(VERSION)
	docker cp $(cid):'/app/build/sliced-darwin-amd64' pkg/sliced-darwin-amd64-$(VERSION)
	docker rm -f $(cid)

.PHONY: clean
clean:
	rm -vrf pkg/*
