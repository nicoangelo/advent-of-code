.SHELL := /usr/bin/env bash
.DEFAULT_GOAL := run
.PHONY: *

day_dir=$(firstword $(subst new,,$(subst test,,$(MAKECMDGOALS))))

test:
	cd day$(day_dir) && go test

run:
	@go run .

new:
	@cp -R template/ day$(day_dir)
	@sed -i 's/dayX/day$(day_dir)/' day$(day_dir)/*
	@sed -i 's/}/\tday$(day_dir).PrintSolutions()\n}/' main.go
	@sed -i 's/^)/\t"github.com\/nicoangelo\/advent-of-code-2024\/day$(day_dir)"\n)/' main.go

%:
	@:
