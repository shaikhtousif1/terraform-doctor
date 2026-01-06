# Terraform Doctor

## What is Terraform Doctor

Terraform Doctor is a small, design-time static analysis CLI for Terraform codebases. It inspects project structure and configuration to highlight architectural and dependency issues that commonly cause operational problems—not syntax errors or runtime failures.

This repository is a skeleton (v1) focused on clear structure and future extensibility.

## The problem it solves

Terraform projects often grow organically, which leads to subtle design issues that are hard to catch with unit tests or linters:

- Excessive module coupling that makes changes risky
- Circular module dependencies that create resource ordering problems
- Mixing local and remote state backends unintentionally
- Cross-environment leakage (prod referencing dev resources)
- Copy-paste configuration drift (dev tags in prod)

These issues are about architecture and organization rather than provider-specific runtime failures.

## What Terraform Doctor does

- Scans a Terraform codebase from a path (no cloud credentials required)
- Builds a lightweight model of modules and relationships
- Runs a set of architectural checks and prints human-readable warnings

It does NOT run `terraform plan` / `apply`, does NOT require credentials, and does NOT enforce policy.

## What it does NOT do

- Validate Terraform syntax comprehensively (use `terraform validate` for that)
- Run or change infrastructure
- Replace policy frameworks like OPA or Sentinel

## Architecture overview

- `cmd/` — CLI entry points (uses `spf13/cobra`)
- `internal/parser` — lightweight scanner that discovers modules (files/directories)
- `internal/graph` — builds a dependency graph (placeholder logic)
- `internal/checks` — individual architectural checks (coupling, circular deps, state mismatch, tagging)
- `internal/report` — simple console reporting

This layout is intentionally simple: `internal/` contains the analysis implementation while `cmd/` wires the CLI.

## Example output

🩺 Terraform Doctor scanning: ./infra
⚠️ Circular dependency detected between modules "eks" and "rds"
⚠️ State mismatch detected between local and remote backends
⚠️ Environment tag mismatch detected (dev tag in prod)

## Roadmap

Planned improvements:

- Replace filesystem heuristics with a real HCL parser for block-level analysis
- Add more checks (secret sprawl, provider pinning, module versioning)
- Add machine-readable output (JSON) and CI integrations
- Provide rules tuning and configuration

Contributions and issues welcome. This is intentionally a small, reviewer-friendly starter.
