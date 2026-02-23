# Bootstrap

Apply the manifests in this folder to bootstrap Argo CD with the root app-of-apps `Application`.

Typical usage:

- `kubectl apply -f bootstrap/`

After the root `Application` is created, Argo CD will reconcile child `Application` manifests under `apps/` (recursively).
