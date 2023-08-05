.PHONY: doc test

doc:
	env GFNCWD=`pwd` go run ./cmd/generate.go

test:
	go test -v ./...