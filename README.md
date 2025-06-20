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

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/UTOL-s/stoken"
)

func main() {
    // Uses configuration based on the current environment
    token := stoken.SToken("Hello, World!")
    fmt.Println(token)
}
```

### Advanced Usage

For more advanced scenarios, you can directly use the configuration and token packages:

```go
package main

import (
    "fmt"
    "github.com/UTOL-s/stoken/internal/token"
    "github.com/UTOL-s/stoken/pkg/config"
)

func main() {
    // Create a custom configuration
    customConfig := config.Config{
        Environment: "custom",
        LogLevel:    "debug",
        Settings: map[string]string{
            "api_url":     "https://custom.example.com",
            "timeout":     "60s",
            "max_retries": "10",
        },
    }

    // Create token service with custom configuration
    tokenService := token.NewTokenService(customConfig)
    customToken, _ := tokenService.Generate("Custom message")
    fmt.Println(customToken)
}
```

See the `examples` directory for more usage examples:
- `examples/basic_usage.go` - Basic usage with different environments
- `examples/advanced_usage.go` - Advanced usage with custom configurations

## Environment Configuration

The module supports different environments through configuration files and environment variables:

### Using Environment Variables

Set the environment using the `STOKEN_ENVIRONMENT` variable:

```bash
# Set environment to production
export STOKEN_ENVIRONMENT=production

# Set log level
export STOKEN_LOG_LEVEL=info

# Set custom settings
export STOKEN_SETTING_API_URL=https://api.example.com
export STOKEN_SETTING_TIMEOUT=10s
```

### Configuration Files

The module looks for configuration files in the `configs` directory with naming pattern `config.<environment>.json`:

- `configs/config.development.json` - Development environment (default)
- `configs/config.testing.json` - Testing environment
- `configs/config.production.json` - Production environment

Example configuration file:

```json
{
  "environment": "production",
  "log_level": "warn",
  "settings": {
    "api_url": "https://api.example.com",
    "timeout": "5s",
    "max_retries": "5"
  }
}
```

## How to Push Updates to This Go Module

### Prerequisites

- Git installed on your machine
- A GitHub account with access to the repository
- Go installed on your machine

### Steps to Push a New Version

1. **Clone the repository** (if you haven't already):
   ```bash
   git clone https://github.com/UTOL-s/stoken.git
   cd stoken
   ```

2. **Make your changes** to the code.

3. **Update the version** in your go.mod file if necessary:
   ```bash
   go mod tidy
   ```

4. **Commit your changes**:
   ```bash
   git add .
   git commit -m "Description of your changes"
   ```

5. **Tag your release** with a semantic version:
   ```bash
   git tag v1.0.0  # Replace with appropriate version
   ```

6. **Push your changes and tags**:
   ```bash
   git push origin main
   git push origin v1.0.0  # Push the tag
   ```

7. **Publish to Go package registry**:
   Go modules are automatically available once you push to a public GitHub repository with proper versioning tags.

8. **Verify the new version**:
   ```bash
   go list -m github.com/UTOL-s/stoken@v1.0.0
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
├── configs/                  # Configuration files for different environments
│   ├── config.development.json
│   ├── config.testing.json
│   └── config.production.json
├── examples/                 # Example usage
│   ├── basic_usage.go
│   └── advanced_usage.go
├── internal/                 # Internal packages
│   └── token/                # Token implementation
│       └── token.go
├── pkg/                      # Public packages
│   └── config/               # Configuration handling
│       └── config.go
├── go.mod                    # Go module definition
├── main.go                   # Main package entry point
└── README.md                 # Documentation
```

## License

[Add your license information here]
