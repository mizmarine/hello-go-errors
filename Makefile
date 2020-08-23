.PHONY: run

run: target=generator
run:
	go run examples/${target}/main.go
