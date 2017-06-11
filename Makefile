# Replace data-sipper with your desired executable name
appname := data-sipper

sources := $(wildcard *.go)

build = GOOS=$(1) GOARCH=$(2) go build -o $(GOPATH)/bin/$(appname)$(3)
build_sqlite3 = cd $(GOPATH)/src/github.com/mattn/go-sqlite3 && go build --tags "libsqlite3 $(1)"
tar = cd $(GOPATH)/bin && tar -cvzf $(appname)_$(1)_$(2).tar.gz $(appname)$(3) && rm $(appname)$(3)
zip = cd $(GOPATH)/bin && zip $(appname)_$(1)_$(2).zip $(appname)$(3) && rm $(appname)$(3)

.PHONY: all windows darwin linux clean

all: windows darwin linux

clean:
	rm -rf $(GOPATH)/bin/*
	rm -rf $(GOPATH)/pkg/*

##### LINUX BUILDS #####
linux: $(GOPATH)/bin/$(appname)_linux_arm.tar.gz $(GOPATH)/bin/$(appname)_linux_arm64.tar.gz $(GOPATH)/bin/$(appname)_linux_386.tar.gz $(GOPATH)/bin/$(appname)_linux_amd64.tar.gz

$(GOPATH)/bin/$(appname)_linux_386.tar.gz: $(sources)
	$(call build_sqlite3,linux,386,)
	$(call build,linux,386,)
	$(call tar,linux,386)

$(GOPATH)/bin/$(appname)_linux_amd64.tar.gz: $(sources)
	$(call build_sqlite3,linux,amd64,)
	$(call build,linux,amd64,)
	$(call tar,linux,amd64)

$(GOPATH)/bin/$(appname)_linux_arm.tar.gz: $(sources)
	$(call build_sqlite3,linux,arm,)
	$(call build,linux,arm,)
	$(call tar,linux,arm)

$(GOPATH)/bin/$(appname)_linux_arm64.tar.gz: $(sources)
	$(call build_sqlite3,linux,arm64,)
	$(call build,linux,arm64,)
	$(call tar,linux,arm64)

##### DARWIN (MAC) BUILDS #####
darwin: $(GOPATH)/bin/$(appname)_darwin_amd64.tar.gz

$(GOPATH)/bin/$(appname)_darwin_amd64.tar.gz: $(sources)
	$(call build,darwin,amd64,)
	$(call tar,darwin,amd64)

##### WINDOWS BUILDS #####
windows: $(GOPATH)/bin/$(appname)_windows_386.zip $(GOPATH)/bin/$(appname)_windows_amd64.zip

$(GOPATH)/bin/$(appname)_windows_386.zip: $(sources)
	$(call $(GOPATH)/bin,windows,386,.exe)
	$(call zip,windows,386,.exe)

$(GOPATH)/bin/$(appname)_windows_amd64.zip: $(sources)
	$(call build,windows,amd64,.exe)
	$(call zip,windows,amd64,.exe)