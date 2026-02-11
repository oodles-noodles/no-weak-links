# no-weak-links
Zelda would be proud

## Overview

This repository contains a basic Go application with intentionally vulnerable dependencies for demonstrating security tracking through Microsoft Defender for Cloud integration with GitHub.

## Features

- Simple HTTP server built with Gin framework (v1.7.0 - intentionally vulnerable)
- Docker containerization
- Automated CI/CD workflow with GitHub Actions
- Azure Container Registry integration
- Build attestation and provenance generation
- Vulnerable dependencies for security scanning demonstration

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Docker
- Azure Container Registry (for deployment)

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/oodles-noodles/no-weak-links.git
cd no-weak-links
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on port 8080. You can access:
- `/` - Hello World endpoint
- `/health` - Health check endpoint

### Building with Docker

Build the Docker image:
```bash
docker build -t no-weak-links:latest .
```

Run the container:
```bash
docker run -p 8080:8080 no-weak-links:latest
```

## GitHub Actions Workflow

The repository includes a CI/CD workflow (`.github/workflows/build-and-push.yml`) that:

1. Builds the Docker image
2. Pushes to Azure Container Registry
3. Generates build attestation for provenance tracking

### Required Secrets

Configure the following secrets in your GitHub repository:

- `ACR_REGISTRY` - Your Azure Container Registry URL (e.g., `myregistry.azurecr.io`)
- `ACR_USERNAME` - Azure Container Registry username
- `ACR_PASSWORD` - Azure Container Registry password

### Workflow Triggers

The workflow runs on:
- Push to `main` branch
- Push to branches matching `copilot/**`
- Pull requests to `main` branch
- Manual trigger via `workflow_dispatch`

## Intentionally Vulnerable Dependencies

This application uses older versions of dependencies with known vulnerabilities:

- `github.com/gin-gonic/gin v1.7.0` - Has known security vulnerabilities
- `golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9` - Outdated crypto library

These vulnerabilities are intentional for demonstration purposes and will be tracked through Microsoft Defender for Cloud.

## Security Scanning

The vulnerable dependencies in this repository allow you to:

1. Demonstrate GitHub's Dependabot security alerts
2. Track vulnerabilities through Microsoft Defender for Cloud
3. Test security scanning integrations
4. Practice vulnerability management workflows

## License

This is a demonstration project. Use at your own risk.
