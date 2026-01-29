# Financial Agent

Financial agent is an important part of Xyber platform. It is responsible for executing financial tasks in a safe and verifiable way, utilizing Intel SGX TEE technology.

## Build

### Local binary

```bash
go build cmd/main.go
```

### Docker image

```bash
docker build -t financial-agent .
```


## Run tests

```bash
go test ./...
```