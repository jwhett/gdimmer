.PHONY: test clean doctest build docs doctest install-docs install-cmd install

MAN ?= /usr/bin/man
DOCSRC = doc/gdimmer.md
DOCOUT = doc/gdimmer.1.gz
DOCDIR ?= /usr/local/share/man/man1/
BINDIR = cmd/gdimmer
BINSRC = $(BINDIR)/main.go
BINOUT = $(BINDIR)/gdimmer
INSTALLOPTS = -D -o root -g root -m 0644


default: build;

test:
	go test

clean:
	$(RM) $(DOCOUT) $(BINOUT)
	
install: install-cmd install-docs

docs:
	pandoc $(DOCSRC) -s -t man | gzip > $(DOCOUT)
	
build: docs
	go build -o $(BINOUT) $(BINSRC)

doctest: docs
	$(MAN) $(DOCOUT)

install-docs: docs
	install $(INSTALLOPTS) $(DOCOUT) -t $(DOCDIR)
	mandb

install-cmd:
	cd $(BINDIR)
	go install
