# Geeper.AI Documentation

This directory contains the documentation website for Geeper.AI, built with Docusaurus.

## What is Geeper.AI?

Geeper.AI is a powerful Kubernetes operator that simplifies the deployment and management of Large Language Models (LLMs) in Kubernetes clusters.

## Features

- **Local Search**: Powered by docusaurus-lunr-search for fast, offline-capable search
- **Versioning**: Support for multiple documentation versions
- **Responsive Design**: Mobile-friendly documentation interface
- **Dark/Light Theme**: Automatic theme switching based on system preferences

## Development

### Prerequisites
- Node.js 18+
- npm or yarn

### Local Development
```bash
# Install dependencies
npm install

# Start development server
npm start

# Build for production
npm run build

# Serve production build
npm run serve
```

### Docker Development
```bash
# Build and run with Docker
make dev

# Build production
make build

# Clean up
make clean
```

## Project Structure

- `docs/` - Documentation pages
- `src/` - Source code and components
- `static/` - Static assets
- `docusaurus.config.ts` - Main configuration
- `sidebars.ts` - Sidebar configuration

## Search Configuration

The documentation uses `docusaurus-lunr-search` for local search functionality:

- **Offline Search**: Search works without external services
- **Fast Indexing**: Built-in search index generation
- **Customizable**: Configurable search bar placement and styling
- **No External Dependencies**: Self-contained search solution

## Contributing

1. Make changes to documentation in the `docs/` directory
2. Update sidebar configuration if needed
3. Test locally with `npm start`
4. Submit a pull request

## Deployment

The documentation can be deployed to:
- GitHub Pages
- Netlify
- Vercel
- Any static hosting service

For GitHub Pages deployment:
```bash
npm run deploy
```

## License

Same as the main project - see the root LICENSE file.
