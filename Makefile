all: clean karousel

clean:

	rm -f karousel
	rm -f karousel.db

karousel:

	go build karousel.go

.PHONY: clean
