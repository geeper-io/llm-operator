# GitHub Actions Setup for Documentation

This guide explains how to set up and use the GitHub Actions workflows for building and deploying the LLM Operator documentation.

## Overview

The repository includes three GitHub Actions workflows:

1. **`docs.yml`** - Main workflow for building and deploying to GitHub Pages
2. **`docs-test.yml`** - Testing workflow for pull requests
3. **`docs-custom-deploy.yml`** - Custom deployment options (Netlify, Vercel, S3, etc.)

## Quick Setup

### 1. Enable GitHub Pages

1. Go to your repository settings
2. Navigate to "Pages" section
3. Set source to "GitHub Actions"
4. Choose a branch (usually `main` or `master`)

### 2. Configure Repository Secrets

For custom deployments, add these secrets in your repository settings:

#### Netlify Deployment
```
NETLIFY_AUTH_TOKEN=your_netlify_token
NETLIFY_SITE_ID=your_site_id
```

#### Vercel Deployment
```
VERCEL_TOKEN=your_vercel_token
ORG_ID=your_org_id
PROJECT_ID=your_project_id
```

#### AWS S3 Deployment
```
AWS_ACCESS_KEY_ID=your_access_key
AWS_SECRET_ACCESS_KEY=your_secret_key
AWS_REGION=your_region
S3_BUCKET=your_bucket_name
CLOUDFRONT_DISTRIBUTION_ID=your_distribution_id
```

## Workflow Details

### Main Deployment Workflow (`docs.yml`)

**Triggers:**
- Push to `main`/`master` branch
- Changes in `www/`, `docs/`, or `examples/` directories
- Manual trigger (`workflow_dispatch`)

**What it does:**
1. Sets up Node.js 18 environment
2. Installs dependencies with caching
3. Builds the Docusaurus site
4. Deploys to GitHub Pages
5. Handles concurrency to prevent conflicts

**Features:**
- Automatic deployment on push
- GitHub Pages integration
- Build artifact caching
- Concurrent deployment protection

### Test Workflow (`docs-test.yml`)

**Triggers:**
- Pull requests to `main`/`master` branch
- Manual trigger

**What it does:**
1. Tests the build process
2. Runs linting (if configured)
3. Verifies build output
4. Uploads build artifacts for inspection

**Features:**
- Prevents broken builds from merging
- Build artifact retention for debugging
- No deployment (safe for PRs)

### Custom Deployment Workflow (`docs-custom-deploy.yml`)

**Triggers:**
- Push to `main`/`master` branch
- Manual trigger with deployment target selection

**Deployment Options:**
- **Netlify**: Static site hosting with CDN
- **Vercel**: Edge deployment platform
- **S3**: AWS static website hosting
- **Custom**: Your own deployment logic

## Usage

### Automatic Deployment

The main workflow runs automatically when you:
- Push changes to the main branch
- Modify files in `www/`, `docs/`, or `examples/` directories

### Manual Deployment

1. Go to "Actions" tab in your repository
2. Select the desired workflow
3. Click "Run workflow"
4. Choose deployment target (for custom workflow)
5. Click "Run workflow"

### Testing Pull Requests

1. Create a pull request
2. The test workflow runs automatically
3. Check the build status and artifacts
4. Merge only if tests pass

## Customization

### Modifying Build Process

Edit the build step in any workflow:

```yaml
- name: Build documentation
  working-directory: www
  run: |
    npm ci
    npm run build
    # Add custom build steps here
```

### Adding Environment Variables

```yaml
- name: Build documentation
  working-directory: www
  run: npm run build
  env:
    NODE_ENV: production
    CUSTOM_VAR: ${{ secrets.CUSTOM_VAR }}
```

### Conditional Steps

```yaml
- name: Deploy to staging
  if: github.ref == 'refs/heads/develop'
  run: echo "Deploying to staging"

- name: Deploy to production
  if: github.ref == 'refs/heads/main'
  run: echo "Deploying to production"
```

## Troubleshooting

### Common Issues

1. **Build fails with dependency errors:**
   - Check `www/package.json` and `www/package-lock.json`
   - Ensure all dependencies are properly specified
   - Clear npm cache if needed

2. **GitHub Pages not updating:**
   - Check workflow run status
   - Verify GitHub Pages source is set to "GitHub Actions"
   - Check for deployment errors in workflow logs

3. **Custom deployment fails:**
   - Verify all required secrets are set
   - Check deployment target configuration
   - Review deployment logs for specific errors

4. **Workflow not triggering:**
   - Check file paths in workflow triggers
   - Verify branch names match
   - Ensure workflow files are in `.github/workflows/` directory

### Debugging

1. **Check workflow runs:**
   - Go to "Actions" tab
   - Click on failed workflow
   - Review step-by-step logs

2. **Download artifacts:**
   - Failed builds create artifacts
   - Download and inspect locally
   - Check for missing files or build errors

3. **Local testing:**
   - Run `npm run build` in `www/` directory
   - Check for local build errors
   - Verify all dependencies are installed

## Advanced Configuration

### Multiple Environments

Create environment-specific workflows:

```yaml
- name: Deploy to staging
  if: github.ref == 'refs/heads/develop'
  run: echo "Staging deployment"

- name: Deploy to production
  if: github.ref == 'refs/heads/main'
  run: echo "Production deployment"
```

### Scheduled Deployments

Add scheduled triggers:

```yaml
on:
  schedule:
    - cron: '0 0 * * 0'  # Weekly on Sunday
  push:
    branches: [ main ]
```

### Matrix Builds

Test against multiple Node.js versions:

```yaml
strategy:
  matrix:
    node-version: [16, 18, 20]

- name: Setup Node.js ${{ matrix.node-version }}
  uses: actions/setup-node@v4
  with:
    node-version: ${{ matrix.node-version }}
```

## Security Considerations

### Secrets Management

- Never commit sensitive information
- Use repository secrets for all credentials
- Rotate secrets regularly
- Limit secret access to necessary workflows

### Permissions

The workflows use minimal required permissions:
- `contents: read` - Read repository content
- `pages: write` - Deploy to GitHub Pages
- `id-token: write` - GitHub token authentication

### Dependency Security

- Use `npm ci` for reproducible builds
- Enable Dependabot for security updates
- Regularly audit dependencies
- Pin dependency versions when needed

## Monitoring and Maintenance

### Workflow Health

- Monitor workflow success rates
- Set up notifications for failures
- Review and optimize build times
- Update dependencies regularly

### Performance Optimization

- Use dependency caching
- Optimize build steps
- Consider build matrix for parallel builds
- Monitor artifact sizes

### Cost Management

- GitHub Actions provides 2000 minutes/month free
- Monitor usage in repository insights
- Optimize workflows to reduce execution time
- Use self-hosted runners for high-volume builds

## Support

For issues with the workflows:

1. Check workflow logs for specific errors
2. Review GitHub Actions documentation
3. Create an issue in the repository
4. Check GitHub Actions status page

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Docusaurus Deployment Guide](https://docusaurus.io/docs/deployment)
- [GitHub Pages Documentation](https://docs.github.com/en/pages)
- [Workflow Syntax Reference](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
