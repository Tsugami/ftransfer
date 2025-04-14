# FTransfer

A Go-based file transfer application.

## Project Structure

```
.
├── api/            # API protocol definitions, OpenAPI/Swagger specs, JSON schema files
├── build/          # Packaging and CI
├── cmd/            # Main applications
├── configs/        # Configuration file templates or default configs
├── deploy/         # IaaS, PaaS, system and container orchestration deployment configs
├── docs/           # Design and architecture documents
├── internal/       # Private application and library code
├── pkg/            # Library code that's ok to use by external applications
├── scripts/        # Scripts to perform various build, install, analysis, etc.
└── test/           # Additional external test applications and test data
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Tsugami/ftransfer.git
cd ftransfer
```

2. Install dependencies:
```bash
go mod tidy
```

### Building

```bash
go build ./...
```

### Running Tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 