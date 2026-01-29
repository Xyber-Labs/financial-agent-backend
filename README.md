# Financial Agent Backend
> **General:** This repository is a core component of the Xyber platform, responsible for executing financial transactions in a secure and verifiable manner using Intel SGX TEE (Trusted Execution Environment) technology.
> It integrates with lending protocols for yield generation and manages user deposits/withdrawals through trust-managed wallet proxies.

## Capabilities

### 1. **REST API Endpoints**

Standard REST endpoints for financial operations.

| Method | Endpoint              | Description                                          |
| :----- | :-------------------- | :--------------------------------------------------- |
| `POST` | `/withdraw`           | Withdraw ERC20 tokens with accumulated yield         |
| `POST` | `/withdraw-native`    | Withdraw native tokens (ETH) with accumulated yield  |
| `GET`  | `/swagger/*`          | Interactive API documentation (Swagger UI)           |

### 2. **TEE/SGX Integration**

Secure transaction signing using Intel SGX attestation.

| Feature                  | Description                                              |
| :----------------------- | :------------------------------------------------------- |
| `Session Key Generation` | Creates ephemeral ECDSA keys within the enclave          |
| `Quote Extraction`       | Retrieves SGX quotes for hardware attestation            |
| `On-Chain Verification`  | Smart contracts verify TEE attestation before execution  |

## API Documentation

This server automatically generates OpenAPI documentation. Once the server is running, you can access the interactive API docs at:

- **Swagger UI**: [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

The interface allows you to explore all REST endpoints, view their schemas, and test them directly from your browser.

## Requirements

- **Go 1.23+**
- **Docker** (for containerization)
- **Intel SGX** (for TEE features, optional for development)
- **Access to Ethereum RPC endpoint** (Base Sepolia testnet by default)

## Setup

1.  **Clone & Configure**
    ```bash
    git clone <repository-url>
    cd financial-agent-backend
    cp config.yaml.example config.yaml
    # Edit config.yaml with your settings (RPC URL, private key, contract addresses)
    ```

2.  **Install Dependencies**
    ```bash
    go mod download
    ```

## Running the Server

### Locally

```bash
# Build the binary
go build -o financial-agent cmd/main.go

# Run with default config
./financial-agent

# Run with custom config path
./financial-agent --config path/to/config.yaml
```

### Using Docker

```bash
# Build the image
docker build -t financial-agent .

# Run the container
docker run -p 8081:8081 -v $(pwd)/config.yaml:/app/config.yaml financial-agent
```

### Using Docker with SGX (Gramine)

For production deployment with Intel SGX:

```bash
cd graminize_scripts

# Generate Gramine signing keys
./0_generate_gram_key.sh

# Build base image
./1_build_base_docker_image.sh

# Build production image
./2_build_prod_docker_image.sh

# Create Gramine enclave
./3_graminize_image.sh

# Deploy
./5_deploy_gram_image.sh
```

## Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Generate coverage report
make coverage
```

## Project Structure

```
financial-agent-backend/
├── cmd/
│   └── main.go                     # Entry point (Cobra CLI)
├── core/
│   ├── server/
│   │   └── http.go                 # Gin HTTP server & handlers
│   ├── transactor/
│   │   ├── transactor.go           # Transaction batching & execution
│   │   └── tee.go                  # TEE/SGX integration
│   ├── onchain/
│   │   └── provider.go             # On-chain operations (deposit, withdraw, claim)
│   ├── network/
│   │   └── event_handler.go        # Block scanner & event processor
│   ├── abi/
│   │   ├── bindings/               # Go contract bindings
│   │   │   ├── TrustManagementRouter/
│   │   │   ├── TrustManagementWallet/
│   │   │   ├── AavePool/
│   │   │   ├── ERC20/
│   │   │   └── WETH/
│   │   └── jsons/                  # Contract ABI JSON files
│   ├── protocol/
│   │   ├── common.go               # Protocol provider interface
│   │   └── aave/                   # Aave protocol implementation
│   └── utils/
│       └── eth.go                  # Ethereum utilities
├── config/
│   └── config.go                   # Configuration management
├── go-tee/                         # SGX quote parsing (git submodule)
│   └── sgx-quote/
│       ├── service.go              # TEE service implementation
│       └── parser.go               # SGX quote parser
├── docs/
│   ├── swagger.json                # OpenAPI specification
│   └── swagger.yaml
├── graminize_scripts/              # SGX Gramine deployment scripts
├── tests/                          # Integration tests
├── config.yaml                     # Application configuration
├── Dockerfile
├── Makefile
├── go.mod
└── README.md
```

## Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│                    HTTP REST API                         │
│                   (Gin, Port 8081)                       │
│           POST /withdraw  POST /withdraw-native          │
└────────────────────────┬────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────┐
│              TrustManagementProvider                     │
│  • Fetches user deposits from router contract           │
│  • Calculates yield-adjusted amounts via Aave aTokens   │
│  • Wraps operations into batched transactions           │
└────────────────────────┬────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────┐
│                     Transactor                           │
│  • Generates TEE session keys (ECDSA)                   │
│  • Extracts SGX quotes for attestation                  │
│  • Signs transaction bundles                            │
│  • Executes via Router.execute() multicall              │
└────────────────────────┬────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────┐
│               Ethereum Network                           │
│  • TrustManagementRouter (user management)              │
│  • TrustManagementWallet (proxy wallets)                │
│  • AavePool (lending protocol)                          │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│              Event Listener (Background)                 │
│  • Polls for new blocks every 10 seconds                │
│  • Processes Deposited events                           │
│  • Auto-supplies tokens to Aave                         │
└─────────────────────────────────────────────────────────┘
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -m 'Add my feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Create a Pull Request

## License

MIT
