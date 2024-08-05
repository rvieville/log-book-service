GO = go
APP = cmd/main.go

RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
BLUE=\033[0;34m
RESET=\033[0m

.PHONY: all run local clean test

all: clean
	@printf "$(BLUE)Start building binary    $(RESET) "
	@$(GO) build -o bin/main $(APP)
	@printf "$(GREEN)Done.$(RESET)\n"

run:
	@printf "$(BLUE)Starting Server\n$(RESET)"
	@${GO} run  $(APP)

local:
	air -c .air.toml

migrate:
	@${GO} run cmd/migration/init.go

clean:
	@printf "$(BLUE)Clean useless packages    $(REST) "
	@${GO} mod tidy
	@printf "$(GREEN)Done.$(RESET)\n"

test:
	@${GO} test -v -coverprofile=./coverage/coverage.out ./internal/services/...
	@${GO} tool cover -html=./coverage/coverage.out -o ./coverage/covergae.html