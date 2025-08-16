.PHONY: help dev build clean

# Default target
help:
	@echo "Available targets:"
	@echo "  dev     - Run Docusaurus dev server using Docker"
	@echo "  build   - Build Docusaurus site using Docker"
	@echo "  clean   - Clean up Docker containers and images"
	@echo "  help    - Show this help message"

# Run dev server using Docker
dev:
	@echo "Starting Geeper.AI documentation dev server with Docker..."
	@echo "The site will be available at http://localhost:3000"
	@echo "Press Ctrl+C to stop the server"
	docker build -t geeper-ai-docs-dev .
	docker run --rm -it \
		-p 3000:3000 \
		-v $(PWD):/app \
		-v /app/node_modules \
		--name geeper-ai-docs-dev \
		geeper-ai-docs-dev

# Build the site using Docker
build:
	@echo "Building Geeper.AI documentation site with Docker..."
	docker build -t geeper-ai-docs-build .
	docker run --rm \
		-v $(PWD):/app \
		-v /app/node_modules \
		geeper-ai-docs-build npm run build

# Clean up Docker resources
clean:
	@echo "Cleaning up Docker resources..."
	docker stop geeper-ai-docs-dev 2>/dev/null || true
	docker rm geeper-ai-docs-dev 2>/dev/null || true
	docker rmi geeper-ai-docs-dev 2>/dev/null || true
	docker rmi geeper-ai-docs-build 2>/dev/null || true
