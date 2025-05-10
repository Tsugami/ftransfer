# FTransfer

FTransfer is a robust file transfer system designed to automate and manage file transfers between different storage systems. It provides a unified interface for handling transfers across various protocols like FTP, SFTP, and S3.

## Features

- **Multi-Protocol Support**
  - FTP/SFTP file transfers
  - S3 Protocol
  - Extensible architecture for new protocols

- **Automated Transfers**
  - Scheduled transfers
  - Real-time monitoring
  - Error handling and retry mechanisms

- **Centralized Management**
  - Web-based dashboard
  - Transfer history and logs
  - Status monitoring

## Quick Start

### Prerequisites

- Go 1.23 or higher
- Node.js 20 or higher
- PostgreSQL 15 or higher
- Docker (optional)

### Installation

#### Using Docker

The easiest way to run FTransfer is using our pre-built Docker image:

```bash
# Pull the latest image
docker pull ghcr.io/tsugami/ftransfer:latest

# Run the container
docker run -p 8080:8080 \
  -e DATABASE_URL="postgres://user:password@host:5432/dbname?sslmode=disable" \
  ghcr.io/tsugami/ftransfer:latest
```

#### Using Releases

The easiest way to get started is to download the latest release from our [releases page](https://github.com/Tsugami/ftransfer/releases).

1. Download the appropriate release for your system
2. Extract the archive
3. Configure your environment variables
4. Run the application

## API Documentation

The API is RESTful and follows standard HTTP conventions. For detailed API documentation, see the [API Documentation](docs/api.md).

### Example API Usage

```bash
# List storage providers
curl -X GET "http://localhost:8080/api/v1/storage-providers"

# Create a new transfer
curl -X POST "http://localhost:8080/api/v1/transfers" \
  -H "Content-Type: application/json" \
  -d '{
    "source_storage_provider_id": "provider-id",
    "destination_storage_provider_id": "provider-id",
    "source_dir": "/source/path",
    "destination_dir": "/destination/path"
  }'
```

## Development

### Project Structure

```
ftransfer/
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── storage_provider/  # Storage provider implementations
│   ├── transfer/         # Transfer logic
│   └── events/          # Event handling
├── frontend/             # Web interface
├── migrations/           # Database migrations
└── docs/                # Documentation
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, please open an issue in the GitHub repository or contact the maintainers.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin) - Go web framework
- [React Admin](https://marmelab.com/react-admin/) - Admin interface framework
- [PostgreSQL](https://www.postgresql.org/) - Relational database