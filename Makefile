.PHONY: test clean doctest build docs doctest install-docs install
default: build;

test:
	go test

clean:
	git clean -dxf

docs:
	pandoc doc/gdimmer.md -s -t man | gzip > doc/gdimmer.1.gz
	
build: docs
	go build -o cmd/gdimmer/gdimmer cmd/gdimmer/main.go

doctest: docs
	/usr/bin/man doc/gdimmer.1.gz

install-docs: docs
	install -o root -g root -m 0644 doc/gdimmer.1.gz /usr/local/share/man/man1/

install-cmd:
	cd cmd/gdimmer
	go install

install: install-cmd install-docs
