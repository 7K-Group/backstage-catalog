# ${{ values.name }}

${{ values.description }}

## Local development

- `go test ./...`
- `go run ./cmd/${{ values.name }}`

## Container

- `docker build -t ${{ values.name }}:local .`
- `docker run --rm -p 8080:8080 ${{ values.name }}:local`
