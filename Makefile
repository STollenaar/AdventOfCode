GO_VER              ?= go

skaff: ## Install skaff
	cd skaff && $(GO_VER) install github.com/STollenaar/AdventOfCode/skaff

.PHONY: skaff