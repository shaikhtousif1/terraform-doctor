# Terraform Doctor

## What is Terraform Doctor

**Terraform Doctor** is a design-time static analysis CLI for Terraform codebases.  
It inspects **structure, boundaries, and dependencies** to detect architectural issues that typically surface **late** in the during audits, scaling, or production changes.

Terraform Doctor focuses on **how Terraform is designed**, not whether it compiles or applies.

This repository is an intentionally lightweight **v1 skeleton**, built for clarity, extensibility, and real-world relevance.

---

## The problem it solves

Terraform projects rarely fail because of syntax.  
They fail because **design decisions age poorly**.

As Terraform usage scales, teams commonly encounter:

- **Excessive module coupling** that increases blast radius
- **Circular dependencies** between infrastructure domains (e.g., EKS ↔ RDS)
- **State topology mismatches** (local vs Terraform Enterprise / remote execution)
- **Cross-environment leakage** (prod referencing dev or shared state)
- **Configuration drift caused by copy-paste** (e.g., dev tags applied to prod)

These issues are architectural, not provider-specific—and are often invisible to linters or `terraform validate`.

Terraform Doctor exists to catch these problems **early**, before they turn Terraform into a fragile system that teams are afraid to change.

---

## What Terraform Doctor does

Terraform Doctor is a **read-only, static analysis tool**.

It:

- Scans a Terraform codebase from a given path
- Builds a lightweight model of modules and dependencies
- Applies architectural heuristics
- Prints **human-readable warnings with context**

It does **not** require cloud credentials and does **not** interact with live infrastructure.

---

## What Terraform Doctor does NOT do

Terraform Doctor is intentionally opinionated and limited in scope.

It does **not**:

- Replace `terraform validate`
- Run `terraform plan` or `apply`
- Enforce policy (OPA / Sentinel)
- Perform cost estimation or runtime drift detection

It is a **design-time diagnostic**, not an enforcement or execution engine.

---

## Architecture overview

The project is structured for clarity and future growth:

- `cmd/`  
  CLI entry points (built with `spf13/cobra`)

- `internal/parser`  
  Discovers Terraform files and module structure (placeholder logic in v1)

- `internal/graph`  
  Builds a dependency graph used for architectural analysis

- `internal/checks`  
  Independent checks for:
  - Module coupling
  - Circular dependencies
  - State execution mismatch
  - Environment and tagging leakage

- `internal/report`  
  Simple console-based reporting

This separation keeps the CLI thin and the analysis logic testable and extensible.

---

## Example output

```text
🩺 Terraform Doctor scanning: ./infra

⚠️ Circular dependency detected between modules "eks" and "rds"
   Impact: Infrastructure domains depend on each other, blocking clean automation.
   Recommendation: Move access wiring into a dedicated integration module.

⚠️ State execution mismatch detected
   Impact: Local plans may attempt to recreate remotely managed resources.
   Recommendation: Align local and remote backends or use remote execution consistently.

⚠️ Environment tag mismatch detected (dev tag in prod)
   Impact: Incorrect cost allocation and audit confusion.
   Recommendation: Centralize environment context and inject it via variables.
````

---

## Roadmap

Planned improvements (post-v1):

* Replace filesystem heuristics with full HCL block-level parsing
* Expand checks around:

  * State boundaries
  * Provider and module version pinning
  * Secret and variable sprawl
* Add machine-readable output (JSON) for CI integration
* Support configurable rule severity and suppression

Terraform Doctor is intentionally starting small to prioritize **signal over noise**.

---

## Status

This project is an early-stage skeleton designed to communicate intent, structure, and real-world relevance.
Contributions, discussions, and feedback are welcome.
