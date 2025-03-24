# Colors
color_green   := $(shell printf "\e[32m")  # Green color
color_reset        := $(shell printf "\e[0m")

# ==================================================================================== #
# HELPERS
# ==================================================================================== #
.PHONY: help
help: ## Print this help message
	@echo ""
	@echo "Usage: make [action]"
	@echo ""
	@echo "Available Actions:"
	@echo ""
	@awk -v green="$(color_green)" -v reset="$(color_reset)" -F ':|##' \
		'/^[^\t].+?:.*?##/ {printf " %s* %-15s%s %s\n", green, $$1, reset, $$NF}' $(MAKEFILE_LIST) | sort
	@echo ""

# ==================================================================================== #
# PUBLIC TASKS
# ==================================================================================== #

.PHONY: live/templ
# Run templ generation in watch mode
live/templ:
	@templ generate --watch --proxy="http://localhost:3300" --open-browser=false

.PHONY: live/server
# Run air to detect file changes, re-build and re-run the server
live/server:
	@go run github.com/air-verse/air@v1.61.7 \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

.PHONY: live/sync
# Watch for asset changes and reload browser via templ proxy
live/sync:
	@go run github.com/air-verse/air@v1.61.7 \
	--build.cmd "tempo sync && templ generate -notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css"

.PHONY: dev
dev: ## Run the dev server with live reload
	make -j3 live/templ live/server live/sync

.PHONY: build
build: ## Build for production with minified asset files
	@tempo sync --prod
	@templ fmt .
	@templ generate
	@go build -o tmp/bin/tempo-demo
