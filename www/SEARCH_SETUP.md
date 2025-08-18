# Search and Versioning Setup Guide

This guide explains how to configure search functionality and versioning for the LLM Operator documentation.

## Search Configuration (Algolia DocSearch)

### 1. Get Algolia Credentials

1. Sign up for a free account at [Algolia](https://www.algolia.com/)
2. Create a new application
3. Get your Application ID and Search API Key

### 2. Update Configuration

Replace the placeholder values in `docusaurus.config.ts`:

```typescript
algolia: {
  appId: 'YOUR_ACTUAL_APP_ID',
  apiKey: 'YOUR_ACTUAL_SEARCH_API_KEY',
  indexName: 'YOUR_ACTUAL_INDEX_NAME',
  // ... other options
}
```

### 3. Configure DocSearch

1. Go to [DocSearch](https://docsearch.algolia.com/)
2. Submit your site for indexing
3. Wait for approval and indexing to complete
4. Update your configuration with the provided credentials

### 4. Test Search

1. Start the development server: `npm run start`
2. Navigate to the documentation
3. Use the search bar in the navigation
4. Test the search page at `/search`

## Versioning Configuration

### 1. Current Setup

The documentation is currently configured with:
- **Current version**: Latest development version
- **Version label**: "Next 🚀"
- **Version path**: `/next`

### 2. Adding New Versions

To add a new version:

1. **Create version directory**:
   ```bash
   mkdir -p www/versioned_docs/version-1.0
   ```

2. **Copy current docs**:
   ```bash
   cp -r www/docs/* www/versioned_docs/version-1.0/
   ```

3. **Update versions.json**:
   ```json
   [
     "current",
     "1.0"
   ]
   ```

4. **Update docusaurus.config.ts**:
   ```typescript
   versions: {
     current: {
       label: 'Next 🚀',
       path: 'next',
     },
     '1.0': {
       label: '1.0',
       path: '1.0',
     },
   }
   ```

### 3. Version Management

- **Current**: Always points to the latest development version
- **Stable**: Points to the latest stable release
- **Legacy**: Previous stable versions for backward compatibility

## API Documentation

### 1. API Sidebar

The API documentation uses a separate sidebar configuration in `sidebars.api.ts`:

```typescript
const sidebars: SidebarsConfig = {
  api: [
    {
      type: 'doc',
      id: 'crd-reference',
      label: 'CRD Reference',
    },
    // ... other API docs
  ],
};
```

### 2. API Routes

API documentation is available at:
- **Current**: `/api/crd-reference`
- **Versioned**: `/api/1.0/crd-reference` (when versions are added)

## Customization

### 1. Search Styling

Customize search appearance in `src/css/custom.css`:

```css
/* Custom search bar styling */
.algolia-autocomplete {
  /* Your custom styles */
}

/* Search results styling */
.algolia-autocomplete .ds-dropdown-menu {
  /* Your custom styles */
}
```

### 2. Version Styling

Customize version dropdown appearance:

```css
/* Version dropdown styling */
.dropdown__menu {
  /* Your custom styles */
}

.dropdown__link {
  /* Your custom styles */
}
```

## Troubleshooting

### Common Issues

1. **Search not working**:
   - Check Algolia credentials
   - Verify index is properly configured
   - Clear browser cache

2. **Version dropdown not showing**:
   - Check `versions.json` configuration
   - Verify version directories exist
   - Restart development server

3. **API docs not loading**:
   - Check sidebar configuration
   - Verify file paths
   - Check for syntax errors

### Getting Help

- **Algolia Support**: [Algolia Help Center](https://help.algolia.com/)
- **Docusaurus Docs**: [Docusaurus Documentation](https://docusaurus.io/docs)
- **GitHub Issues**: Report bugs in the LLM Operator repository

## Production Deployment

### 1. Environment Variables

For production, use environment variables:

```typescript
algolia: {
  appId: process.env.ALGOLIA_APP_ID || 'YOUR_APP_ID',
  apiKey: process.env.ALGOLIA_SEARCH_API_KEY || 'YOUR_SEARCH_API_KEY',
  indexName: process.env.ALGOLIA_INDEX_NAME || 'YOUR_INDEX_NAME',
}
```

### 2. Build and Deploy

1. **Build documentation**:
   ```bash
   npm run build
   ```

2. **Deploy to your hosting platform**:
   - GitHub Pages
   - Netlify
   - Vercel
   - Custom server

3. **Verify functionality**:
   - Test search on production
   - Check version navigation
   - Verify API documentation

## Security Notes

- **Public API Key**: The search API key is safe to commit to version control
- **Admin API Key**: Never commit the admin API key
- **Rate Limiting**: Algolia provides generous free tier limits
- **Content Filtering**: Search results are based on your indexed content only
