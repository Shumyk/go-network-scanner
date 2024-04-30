NMAP_VERSION = 7.95
NMAP_ARCHIVE = nmap-$(NMAP_VERSION)
NMAP_ARCHIVE_EXT = $(NMAP_ARCHIVE).tar.bz2
NMAP_SOURCE_DIR = nmap-source-code-$(NMAP_VERSION)

build:
	go build -o network-scanner main.go

clone-nmap-source:
	@if [[ ! -d "$(NMAP_SOURCE_DIR)" ]]; then \
		echo "### Cloning nmap source..."; \
		wget https://nmap.org/dist/$(NMAP_ARCHIVE_EXT); \
		tar -xvf $(NMAP_ARCHIVE_EXT); \
		rm $(NMAP_ARCHIVE_EXT); \
		mv $(NMAP_ARCHIVE) $(NMAP_SOURCE_DIR); \
	fi

remove-nmap-source:
	@if [[ -d "$(NMAP_SOURCE_DIR)" ]]; then \
		echo "### Removing nmap source..."; \
		sudo rm -rf $(NMAP_SOURCE_DIR); \
	fi

install-nmap:
	@if command -v nmap >/dev/null 2>&1; then \
		echo "### Nmap is already installed. Skipping installation."; \
		exit 0; \
	fi
	$(MAKE) clone-nmap-source
	@cd $(NMAP_SOURCE_DIR); \
		./configure; \
		make; \
		sudo make install

uninstall-nmap:
	@if [[ ! -d "$(NMAP_SOURCE_DIR)" ]]; then \
		echo "### Nmap source directory does not exist. Skipping uninstallation."; \
		exit 1; \
	fi
	@echo "### Uninstalling nmap..."
	cd $(NMAP_SOURCE_DIR); sudo make uninstall
	$(MAKE) remove-nmap-source
	@echo "### Done."