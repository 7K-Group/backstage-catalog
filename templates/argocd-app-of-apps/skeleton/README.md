# Argo CD App-of-Apps Repository

This repository contains:

- `bootstrap/`: bootstrap manifests, including the root app-of-apps Argo CD `Application`
- `apps/`: place child Argo CD `Application` manifests here (in subfolders, one per domain/project)
- `${{ values.name }}-extras/`: extra Kubernetes manifests synced by a child `Application`
- `catalog-info.yaml`: Backstage catalog registration for the Resource
