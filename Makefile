.PHONY: main
GOPATH=$(PWD)/.gopath
PKG=$(shell cat .goimportpath)


main:
	mkdir -p $(GOPATH)/src/$(shell dirname $(PKG))
	-rm -fr $(GOPATH)/src/$(PKG)
	ln -s $(PWD) $(GOPATH)/src/$(PKG)
	GOPATH=$(GOPATH) go install -v $(PKG)/main

clean:
	go clean -cache
