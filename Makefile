VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null)

all: build-production package checksums

build-production:
	rm -rf build
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.currentVersion=$(VERSION) -s -w" -o build/run-macos-64/bin/run
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.currentVersion=$(VERSION) -s -w" -o build/run-windows-64/bin/run.exe
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.currentVersion=$(VERSION) -s -w" -o build/run-linux-64/bin/run

package:
	rm -rf dist
	mkdir dist
	tar -czf dist/run-$(VERSION)-macos-64.tar.gz LICENSE -C build/run-macos-64 .
	tar -czf dist/run-$(VERSION)-windows-64.tar.gz LICENSE -C build/run-windows-64 .
	tar -czf dist/run-$(VERSION)-linux-64.tar.gz LICENSE -C build/run-linux-64 .

checksums:
	cd dist && shasum -a 256 * > run-$(VERSION)-checksums.txt