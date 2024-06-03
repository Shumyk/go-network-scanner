NMAP_VERSION = 7.95
NMAP_ARCHIVE = nmap-$(NMAP_VERSION)
NMAP_ARCHIVE_EXT = $(NMAP_ARCHIVE).tar.bz2
NMAP_SOURCE_DIR = nmap-source-code-$(NMAP_VERSION)

EXECUTABLE=network-scanner
BUILD_DIR=build
VERSION= 0.0.1
BUILD= $(shell git rev-parse --short HEAD)

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-v -X main.version=${VERSION} -X main.build=${BUILD}"


help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	go test -v ./...

clone-nmap-source: ## Clone nmap source code from the official repository.
	@if [[ ! -d "$(NMAP_SOURCE_DIR)" ]]; then \
		echo "### Cloning nmap source..."; \
		wget https://nmap.org/dist/$(NMAP_ARCHIVE_EXT); \
		tar -xvf $(NMAP_ARCHIVE_EXT); \
		rm $(NMAP_ARCHIVE_EXT); \
		mv $(NMAP_ARCHIVE) $(NMAP_SOURCE_DIR); \
	fi

remove-nmap-source: ## Remove nmap source code.
	@if [[ -d "$(NMAP_SOURCE_DIR)" ]]; then \
		echo "### Removing nmap source..."; \
		sudo rm -rf $(NMAP_SOURCE_DIR); \
	fi

install-nmap: ## Install nmap. This checks if nmap is installed, if not, it installs nmap.
	@type nmap >/dev/null 2>&1 && echo "### Nmap is already installed. Skipping installation."
	@type nmap >/dev/null 2>&1 || $(MAKE) --install-nmap

--install-nmap: ## Install nmap. This will clone the source code if absent, compile and install nmap.
	@echo "### Installing nmap..."
	$(MAKE) clone-nmap-source
	@cd $(NMAP_SOURCE_DIR); \
		./configure; \
		make; \
		sudo make install
	@echo "### nmap installed."

uninstall-nmap: ## Uninstall nmap and remove the source code.
	@if [[ ! -d "$(NMAP_SOURCE_DIR)" ]]; then \
		echo "### Nmap source directory does not exist. Skipping uninstallation."; \
		exit 1; \
	fi
	@echo "### Uninstalling nmap..."
	cd $(NMAP_SOURCE_DIR); sudo make uninstall
	$(MAKE) remove-nmap-source
	@echo "### nmap uninstalled."

build: ## Build the application for current OS
	@sudo echo "### Installing Nmap..."
	$(MAKE) install-nmap

	@echo "### Build frontend..."
	@cd ui/pages && sudo yarn install && yarn build
	@echo "### Frontend build done."

	$(MAKE) setup-webui

	@echo "### Building the application..."
	go build  -o $(EXECUTABLE) -ldflags=$(LDFLAGS) main.go
	@echo "### Go binary build done. Run `sudo ./network-scanner` to start the application."

setup-webui: ## Setup webui library
	@echo "### Setting up WebUI C library..."
	ifeq ($(OS),Windows_NT)
		@echo "### Windows OS detected. Setting up..."
		@irm https://raw.githubusercontent.com/webui-dev/go-webui/main/setup.ps1 | iex
	else
		@sh -c "$(curl -fsSL https://raw.githubusercontent.com/webui-dev/go-webui/main/setup.sh)"
	endif
	@echo "### WebUI C library setup done."

install-xgo: ## Install xgo
	@echo "### Installing xgo..."
	@go install src.techknowlogick.com/xgo@latest
	@echo "### xgo installed."

cross-build: install-xgo ## Cross compile the application. Requires xgo and Docker.
	@echo "### Cross compiling the application..."
	@echo "### version $(VERSION), build $(BUILD)"
	@xgo \
		--hooksdir buildhooks \
		--targets=linux/amd64,linux/arm64,darwin-11.1.0/amd64,darwin-11.1.0/arm64,windows/amd64 \
		-out $(BUILD_DIR)/$(EXECUTABLE) \
		-ldflags=$(LDFLAGS) \
		./main.go

build-windows: install-xgo ## Cross build for Windows. Requires xgo and Docker.
	@echo "### Building for Windows..."
	@echo "### version $(VERSION), build $(BUILD)"
	@xgo \
		--targets=windows/amd64 \
		-out $(BUILD_DIR)/$(EXECUTABLE) \
		-ldflags=$(LDFLAGS) \
		./main.go
	@echo "### Windows build done."


build-linux: install-xgo ## Cross build for Linux. Requires xgo and Docker.
	@echo "### Building for Linux..."
	@echo "### version $(VERSION), build $(BUILD)"
	@xgo \
		--targets=linux/amd64,linux/arm64 \
		-out $(BUILD_DIR)/$(EXECUTABLE) \
		-ldflags=$(LDFLAGS) \
		./main.go
	@echo "### Linux build done."

build-darwin: install-xgo ## Cross build for Darwin (macOS). Requires xgo and Docker.
	@echo "### Building for Darwin..."
	@echo "### version $(VERSION), build $(BUILD)"
	@xgo \
		--targets=darwin-11.1.0/amd64,darwin-11.1.0/arm64 \
		-out $(BUILD_DIR)/$(EXECUTABLE) \
		-ldflags=$(LDFLAGS) \
		./main.go

clean-builds: ## Remove previous build
	rm -r $(BUILD_DIR)

clean-frontend: ## Remove frontend build
	rm -rf ui/pages/build

clean-deep: ## Remove previous builds, nmap source code, frontend build
	$(MAKE) clean-builds
	$(MAKE) remove-nmap-source
	$(MAKE) clean-frontend
