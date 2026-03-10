# Catalog structure

This repository is intended to be registered in Backstage as a **single Location**.
All catalog entities live under `catalog/` and are separated by function:

- `catalog/sourcecode/`: Components (services, libraries)
- `catalog/apis/`: APIs provided/consumed by components
- `catalog/cicd/`: CI/CD resources (pipelines, workflows, deploy definitions)
- `catalog/tests/`: test-suite resources
- `catalog/monitoring/`: monitoring resources (dashboards, alerts)
- `catalog/systems/`: systems (collections of components)
- `catalog/domains/`: domains (business/engineering domains)
- `catalog/groups/`: groups (team ownership)
- `catalog/environments/`: environment resources (dev/stage/prod)

Conventions:

- `metadata.name`: kebab-case, globally unique within its kind
- `spec.owner`: `group:platform` (or another group defined in `catalog/groups/`)
- Components should reference their system via `spec.system`.

## Organization inheritance graph

```mermaid
graph TD
	ORG[Group: inari<br/>type: organization]
	ORG --> DCP[Group: developer-control-plane<br/>type: plane]
	ORG --> RP[Group: resource-plane<br/>type: plane]
	ORG --> IDP[Group: integration-delivery-plane<br/>type: plane]
	ORG --> SP[Group: security-plane<br/>type: plane]
	ORG --> OP[Group: observability-plane<br/>type: plane]
	ORG --> CT[Group: core-team<br/>type: team]
	DCP --> DET[Group: developer-experience-team<br/>type: team]
	RP --> CIT[Group: cloud-infrastructure-team<br/>type: team]
	IDP --> DCT[Group: delivery-control-team<br/>type: team]
	SP --> SCT[Group: security-compliance-team<br/>type: team]
	OP --> ORT[Group: observability-reliability-team<br/>type: team]
	U[User: joaodss] -->|memberOf| CT
	U -->|memberOf| DCP
	U -->|memberOf| ORG
```

Backstage relations are direct; user membership is not automatically inferred from child teams to parent groups.
To show a user on team, plane, and organization levels, include each ancestor in `User.spec.memberOf` (or define equivalent direct `Group.spec.members` entries).
