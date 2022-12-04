.SHELL := /usr/bin/env bash
.DEFAULT_GOAL := run
.PHONY: *

day_dir=$(firstword $(subst test,,$(MAKECMDGOALS)))

test:
	cd day$(day_dir) && go test

run:
	@go run .

%:
	@: