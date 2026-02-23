# Keycloak Instance (Argo CD App-of-Apps)

This repository contains:

- `bootstrap/`: bootstrap manifests, including the root app-of-apps Argo CD `Application`
- `apps/`: child Argo CD `Application` manifests (organized by domain/project subfolders)
- `catalog-info.yaml`: Backstage catalog registration for the Keycloak Resource
