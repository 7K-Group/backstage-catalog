# {{ values.name }}

{{ values.description }}

## Endpoints

- `GET /healthz` -> `ok`
- `GET /` -> service info JSON

## Run locally

```bash
go test ./...
go run ./cmd/server
```

## Build

```bash
go build ./cmd/server
```

## Docker

```bash
docker build -t {{ values.name }}:local .
docker run --rm -p 8080:8080 -e PORT=8080 {{ values.name }}:local
```
