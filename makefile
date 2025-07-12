# Project name
APP_NAME := afrconsole

# Default OS targets by user flag
TARGET ?= linux

# Architecture options
# ARCHS := amd64 arm64
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

# Determine which platforms to build for
ifeq ($(TARGET),all)
    PLATFORMS := $(ALL_TARGETS)
else ifeq ($(TARGET),macos)
    PLATFORMS := $(call expand_targets,darwin)
else ifeq ($(TARGET),windows)
    PLATFORMS := $(call expand_targets,windows)
else
    PLATFORMS := $(call expand_targets,linux)
endif

# Build all
build:
	@echo "Building $(APP_NAME) for: $(TARGET)"
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
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
