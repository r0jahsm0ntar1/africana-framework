# Project name
APP_NAME := afrconsole

# Architecture options
ARCHS := amd64

# Output directory
BUILD_DIR := build

# Map OS keyword to GOOS
define expand_targets
$(foreach arch,$(ARCHS),$(1)/$(arch))
endef

# All supported targets
ALL_TARGETS := \
    $(call expand_targets,linux) \
    $(call expand_targets,windows) \
    $(call expand_targets,darwin)

# Default platform based on system
PLATFORMS := $(call expand_targets,$(shell uname -s | tr A-Z a-z))

# Default make (build for current OS)
default: build

# Build for current system
build:
	@echo "Building $(APP_NAME) for current system..."
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d'/' -f1); \
		GOARCH=$$(echo $$platform | cut -d'/' -f2); \
		EXT=$$( [ "$$GOOS" = "windows" ] && echo ".exe" || echo "" ); \
		OUT=$(BUILD_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH$$EXT; \
		echo " > $$OUT"; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -v -x -o $$OUT . || exit 1; \
	done

# Build for specified OS (e.g., make distro windows)
distro:
	@OS_NAME=$(filter-out $@,$(MAKECMDGOALS)); \
	if [ -z "$$OS_NAME" ]; then \
		echo "Usage: make distro <os>"; \
		echo "Supported OS: linux, windows, darwin, all"; \
		exit 1; \
	fi; \
	if [ "$$OS_NAME" = "all" ]; then \
		PLATFORMS="$(ALL_TARGETS)"; \
	else \
		PLATFORMS="$(call expand_targets,$$OS_NAME)"; \
	fi; \
	echo "Building for: $$OS_NAME"; \
	mkdir -p $(BUILD_DIR); \
	for platform in $$PLATFORMS; do \
		GOOS=$$(echo $$platform | cut -d'/' -f1); \
		GOARCH=$$(echo $$platform | cut -d'/' -f2); \
		EXT=$$( [ "$$GOOS" = "windows" ] && echo ".exe" || echo "" ); \
		OUT=$(BUILD_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH$$EXT; \
		echo " > $$OUT"; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -v -x -o $$OUT . || exit 1; \
	done

# Clean output
clean:
	@echo "Cleaning build directory..."
	@go clean -r -x -cache
	@rm -rf $(BUILD_DIR)

# Native run helper
run:
	@BIN=$(BUILD_DIR)/$(APP_NAME)-$(shell uname -s | tr A-Z a-z)-$(shell uname -m); \
	echo "Running $$BIN..."; \
	chmod +x $$BIN && $$BIN

# Prevent Make from trying to treat distro targets like files
.PHONY: default build distro clean run linux windows darwin all
%:
	@:
