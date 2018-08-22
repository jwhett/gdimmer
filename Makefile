.PHONY: test clean
default: test;

test:
	go test

clean:
	git clean -dxf
