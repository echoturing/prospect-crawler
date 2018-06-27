.PHONY: main
GOPATH=$(PWD)/.gopath
PKG=$(shell cat .goimportpath)


main:
	-rm -fr $(GOPATH)
	mkdir -p $(GOPATH)/src/$(shell dirname $(PKG))
	mkdir -p bin
	ln -s $(PWD) $(GOPATH)/src/$(PKG)
	ln -s $(PWD)/bin $(GOPATH)/bin
	GOPATH=$(GOPATH) go install -v $(PKG)/main

clean:
	-rm -fr $(GOPATH)
	go clean -cache
