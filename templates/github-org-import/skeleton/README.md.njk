# Backstage GitHub org import

{{ values.description }}

## What this does

This repository contains a single Backstage `Location` entity of type `github-discovery`.
When this repo is registered in Backstage, the GitHub discovery processor will scan:

- `https://{{ values.githubHost }}/{{ values.githubOrg }}/*/blob/{{ values.defaultBranch }}/{{ values.catalogInfoPath }}`

and import the catalog descriptor from each matching repository.

If you used the **pull request into an existing repo** option, make sure the generated `catalog-info.yaml` (in the target folder) is referenced by an already-ingested Backstage `Location` (for example by adding it to an index `Location` file in your catalog repo).

## Prerequisites in your Backstage instance

- Catalog backend has the GitHub discovery processor enabled (from `@backstage/plugin-catalog-backend-module-github`).
- Backstage has GitHub credentials configured (token or GitHub App) that can read the target repositories (especially for private repos).

If you used this template to **create a repo** or **open a PR**, the configured credentials must also have permissions to write to the destination repository/organization (e.g. GitHub App permissions such as `Contents: Read & write` and `Pull requests: Read & write`, and installation access to the org/repo).

Troubleshooting:

- If you see `Resource not accessible by integration`, the GitHub App is authenticated but does not have access to the repo (not installed on that repo/org, or missing required write permissions).
