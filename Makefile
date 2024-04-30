NMAP_SOURCE_DIR = nmap-source-code

clone-nmap-repo:
	@if [ ! -d "$(NMAP_SOURCE_DIR)" ]; then \
		echo "Installing nmap source code..."; \
		git clone --depth 1 --no-tags https://github.com/nmap/nmap.git $(NMAP_SOURCE_DIR); \
	fi

remove-nmap-repo:
	@if [ -d "$(NMAP_SOURCE_DIR)" ]; then \
		echo "Removing nmap source..."; \
		sudo rm -rf $(NMAP_SOURCE_DIR); \
	fi

install-nmap: clone-nmap-repo
	@sudo echo "Installing nmap from source..."
	@cd $(NMAP_SOURCE_DIR); \
		./configure; \
		make; \
		sudo make install

uninstall-nmap: clone-nmap-repo
	@echo "Uninstalling nmap..."
	@cd $(NMAP_SOURCE_DIR); \
		sudo make uninstall
	@remove-nmap-repo