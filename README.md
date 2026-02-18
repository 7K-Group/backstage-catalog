# backstage-catalog

Backstage catalog repository for the SaaS developer platform.

## Register this repo in Backstage

In Backstage:

- Go to **Catalog** → **Register Existing Component**
- Use the URL to this repo’s `catalog-info.yaml`

Once registered, Backstage will ingest:

- Entities under `catalog/`
- Templates under `templates/`

## Creating templates with multiple entities

The recommended pattern is:

- The generated service repo has a **root** `catalog-info.yaml` of kind `Location`
- That Location points to multiple entity files (e.g. `./catalog/component.yaml`, `./catalog/resources.yaml`)

This lets a single `catalog:register` step register *many* entities at once (Component + CI Resource + Docker image Resource + Helm chart Resource), even when each entity lives in its own YAML file.

## Included starter template

- `Go Microservice (Docker + GitHub Actions + Helm)` in `templates/go-microservice/`
