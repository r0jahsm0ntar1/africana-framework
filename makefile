# Colors - Basic
Endc   := \033[0m
Bold   := \033[1m
Dim    := \033[2m
Italic := \033[3m
Underl := \033[4m
Blink  := \033[5m
Blink2 := \033[6m
Invert := \033[7m
Hidden := \033[8m
Strike := \033[9m

# Colors - Standard
Black   := \033[30m
Red     := \033[31m
Green   := \033[32m
Yellow  := \033[33m
Blue    := \033[34m
Magenta := \033[35m
Cyan    := \033[36m
White   := \033[37m

# Colors - Bright
Grey          := \033[90m
BrightRed     := \033[91m
BrightGreen   := \033[92m
BrightYellow  := \033[93m
BrightBlue    := \033[94m
BrightMagenta := \033[95m
BrightCyan    := \033[96m
BrightWhite   := \033[97m

# Define banner as a shell function
define PRINT_BANNER
	echo "${BrightYellow} ,__,             ${Endc}"; \
	echo "${BrightYellow} (oo)____         ${Endc}"; \
	echo "${BrightYellow} (__)    )\\       ${Endc}"; \
	echo "${BrightYellow}    ||--||        ${Endc}"; \
	echo "${BrightYellow}    John 3:16     ${Endc}";
endef

define PRINT_BANNER0
	echo "${BrightYellow}           .--,   ${Endc}"; \
	echo "${BrightYellow}       ,.-( (o)\\  ${Endc}"; \
	echo "${BrightYellow}      /   .)/\\ ') ${Endc}"; \
	echo "${BrightYellow}    .',./'/   )/  ${Endc}"; \
	echo "${BrightYellow}()=///=))))==()   ${Endc}"; \
	echo "${BrightYellow}  / John 3:16     ${Endc}";
endef

# Project name
APP_NAME := afrconsole

# Output directory
BUILD_DIR := build

# Detect current architecture
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
    ARCHS := amd64
else ifeq ($(UNAME_M),aarch64)
    ARCHS := arm64
else ifeq ($(UNAME_M),armv7l)
    ARCHS := arm
else
    ARCHS := $(UNAME_M)
endif

# Map OS keyword to GOOS
define expand_targets
$(foreach arch,$(ARCHS),$(1)/$(arch))
endef

# All supported targets
ALL_TARGETS := \
    $(call expand_targets,linux) \
    $(call expand_targets,windows) \
    $(call expand_targets,darwin)

# Detect current OS
UNAME_S := $(shell uname -s | tr A-Z a-z)
ifeq ($(UNAME_S),linux)
    # Check if we're in Termux (Android)
    ifneq ($(wildcard /data/data/com.termux/files/usr/bin/termux-info),)
        PLATFORMS := $(call expand_targets,android)
    else
        PLATFORMS := $(call expand_targets,linux)
    endif
else
    PLATFORMS := $(call expand_targets,$(UNAME_S))
endif

# Default make (build for current OS)
default: build

# Build for current system
build:
	@echo "${BrightGreen}${Bold}[*] ${Endc}Detecting system ...${Endc}"
	@echo "${BrightBlue}${Bold}[+] ${Endc}Building ${Green}${APP_NAME}${Endc} for current system ...${Endc}"
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d'/' -f1); \
		GOARCH=$$(echo $$platform | cut -d'/' -f2); \
		EXT=$$( [ "$$GOOS" = "windows" ] && echo ".exe" || echo "" ); \
		OUT=$(BUILD_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH$$EXT; \
		echo "${Magenta}${Dim}   >${Endc} ${BrightWhite}$$OUT${Endc}"; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$OUT .; \
		if [ $$? -eq 0 ]; then \
			$(PRINT_BANNER) \
			echo "${Blue}${Bold}[+] ${Endc}Building completed successfully ...${Endc}"; \
		else \
			echo "${Red}${Bold}[!] ${Endc}${Cyan}Build failed. Please retry again ...${Endc}"; \
			exit 1; \
		fi; \
	done

# Build for specified OS (e.g., make distro windows)
distro:
	@OS_NAME=$(filter-out $@,$(MAKECMDGOALS)); \
	if [ -z "$$OS_NAME" ]; then \
		$(PRINT_BANNER0) \
		echo "${BrightRed}${Bold}Usage:${Endc} make distro <os>"; \
		echo "${BrightBlue}${Underl}Supported OS:${Endc} linux, windows, darwin, android, all"; \
		exit 1; \
	fi; \
	if [ "$$OS_NAME" = "all" ]; then \
		PLATFORMS="$(ALL_TARGETS)"; \
		echo "${BrightBlue}${Bold}[+] ${Endc}Building for ${BrightGreen}${Underl}all platforms${BrightBlue} ...${Endc}"; \
	else \
		PLATFORMS="$(call expand_targets,$$OS_NAME)"; \
		echo "${BrightBlue}${Bold}[+] ${Endc}Building for ${BrightCyan}${Underl}$$OS_NAME${BrightBlue} ...${Endc}"; \
	fi; \
	mkdir -p $(BUILD_DIR); \
	for platform in $$PLATFORMS; do \
		GOOS=$$(echo $$platform | cut -d'/' -f1); \
		GOARCH=$$(echo $$platform | cut -d'/' -f2); \
		EXT=$$( [ "$$GOOS" = "windows" ] && echo ".exe" || echo "" ); \
		OUT=$(BUILD_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH$$EXT; \
		echo "${Magenta}${Dim}   >${Endc} ${BrightWhite}$$OUT${Endc}"; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$OUT .; \
		if [ $$? -eq 0 ]; then \
			$(PRINT_BANNER) \
			echo "${Blue}${Bold}[+] ${Endc}Building completed successfully ...${Endc}"; \
		else \
			echo "${BrightRed}[!] ${Endc}${Cyan}Build failed. Please retry again ...${Endc}"; \
			exit 1; \
		fi; \
	done

# Clean output
clean:
	@echo "${BrightYellow}${Bold}[!] ${Endc}Cleaning build directory ...${Endc}"
	@go clean -cache
	@rm -rf $(BUILD_DIR)
	@echo "${BrightBlue}${Bold}[+] ${Endc}Clean completed!${Endc}"

# Native run helper
run:
	@GOOS=$(shell uname -s | tr A-Z a-z); \
	if [ -f /data/data/com.termux/files/usr/bin/termux-info ]; then GOOS="android"; fi; \
	GOARCH=$(shell uname -m); \
	# Map architecture names \
	if [ "$$GOARCH" = "x86_64" ]; then GOARCH="amd64"; fi; \
	if [ "$$GOARCH" = "aarch64" ]; then GOARCH="arm64"; fi; \
	if [ "$$GOARCH" = "armv7l" ]; then GOARCH="arm"; fi; \
	EXT=; \
	if [ "$$GOOS" = "windows" ]; then EXT=".exe"; fi; \
	BIN=$(BUILD_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH$$EXT; \
	if [ ! -f "$$BIN" ]; then \
		echo "${BrightRed}[!] ${Endc}Binary not found: ${BrightCyan}$$BIN${Endc}"; \
		echo "${BrightBlue}Run ${BrightCyan}make build${BrightBlue} first${Endc}"; \
		exit 1; \
	fi; \
	echo "${BrightBlue}${Bold}[+] ${Endc}Running ${Green}$$BIN ${Endc}...${Endc}"; \
	chmod +x $$BIN; \
	./$$BIN -q; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then \
		echo "${Blue}${Bold}[*] ${Endc}Execution completed successfully ...${Endc}"; \
	else \
		echo "${Red}${Bold}[!] ${Endc}Execution failed with exit code $$EXIT_CODE${Endc}"; \
	fi

# Help target
help:
	@echo "${BrightCyan}${Bold}${Underl}Available commands:${Endc}"
	@echo "  ${BrightGreen}make${Endc}           - Build for current system"
	@echo "  ${BrightGreen}make build${Endc}     - Build for current system"
	@echo "  ${BrightGreen}make distro <os>${Endc} - Build for specific OS"
	@echo "  ${BrightGreen}make clean${Endc}     - Clean build artifacts"
	@echo "  ${BrightGreen}make run${Endc}       - Run the built binary"
	@echo "  ${BrightGreen}make help${Endc}      - Show this help"
	@echo ""
	@echo "${BrightYellow}${Bold} Examples:${Endc}"
	@echo "  ${BrightWhite}make distro linux${Endc}   - ${Dim}Build for Linux${Endc}"
	@echo "  ${BrightWhite}make distro windows${Endc} - ${Dim}Build for Windows${Endc}"
	@echo "  ${BrightWhite}make distro darwin${Endc}  - ${Dim}Build for macOS${Endc}"
	@echo "  ${BrightWhite}make distro android${Endc} - ${Dim}Build for Android${Endc}"
	@echo "  ${BrightWhite}make distro all${Endc}     - ${Dim}Build for all platforms${Endc}"
	@echo ""
	@echo "${BrightBlue}Current OS: ${BrightCyan}$(shell uname -s)${Endc}"
	@echo "${BrightBlue}Current Arch: ${BrightCyan}$(shell uname -m) → $(ARCHS)${Endc}"

# Prevent Make from trying to treat distro targets like files
.PHONY: default build distro clean run help
%:
	@:
