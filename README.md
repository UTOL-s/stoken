# SToken Go Module

A Go module that provides token-related functionality with versatile environment configuration support.

## Features

- **Environment-based Configuration**: Automatically adapts behavior based on the current environment (development, testing, production)
- **Multiple Configuration Sources**: Load configuration from files, environment variables, or use defaults
- **Flexible Settings**: Customize token generation with environment-specific settings

## Installation

```bash
go get github.com/UTOL-s/stoken
```


## How to Push Updates to This Go Module

### CI/CD Process

This repository uses GitHub Actions for continuous integration and delivery:

#### Pull Request Workflow
When a pull request is opened against the main branch, the following checks are automatically run:
- Dependency verification
- Unit tests
- Code formatting checks
- Static analysis with go vet

#### Release Workflow
When changes are merged to the main branch, a release workflow is triggered that:
1. Runs tests to ensure code quality
2. Determines the next semantic version based on commit history
3. Creates a new tag and GitHub release
4. Publishes the module to the Go package registry

### Manual Release Process (if needed)

If you need to manually release a new version:

1. **Clone the repository** (if you haven't already):
   ```bash
   git clone https://github.com/UTOL-s/stoken.git
   cd stoken
   ```

2. **Make your changes** to the code.

3. **Update dependencies** if necessary:
   ```bash
   go mod tidy
   ```

4. **Commit your changes**:
   ```bash
   git add .
   git commit -m "Description of your changes"
   ```

5. **Push your changes** to trigger the automated release:
   ```bash
   git push origin main
   ```

6. **Verify the new version** after the GitHub Action completes:
   ```bash
   go list -m github.com/UTOL-s/stoken@<new-version>
   ```

### Best Practices for Go Module Versioning

- Follow [Semantic Versioning](https://semver.org/) (SemVer) for your tags.
- Major version changes (v1 → v2) that include breaking changes should use a different module path (e.g., `/v2` suffix).
- Include a CHANGELOG.md to document changes between versions.
- Use go.mod's `replace` directive during local development if needed.

## Project Structure

The module is organized with the following structure:

```
stoken/
├── .github/                  # GitHub specific files
│   └── workflows/            # CI/CD workflow definitions
│       ├── release.yml       # Release workflow
│       └── test.yml          # Test workflow
├── configs/                  # Configuration files for different environments
│   ├── config.development.json
│   ├── config.testing.json
│   └── config.production.json
├── internal/                 # Internal packages
│   └── token/                # Token implementation
│       └── token.go
├── pkg/                      # Public packages
│   └── config/               # Configuration handling
│       └── config.go
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
├── main.go                   # Main package entry point
└── README.md                 # Documentation
```

## License

MIT License. See the LICENSE file for details.
