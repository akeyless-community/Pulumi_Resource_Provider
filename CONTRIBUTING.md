# Contributing

Thank you for helping improve this provider.

## Code of conduct

Participation is governed by the [Contributor Covenant](CODE-OF-CONDUCT.md).

## Development setup

1. Install [Go](https://go.dev/dl/), the [Pulumi CLI](https://www.pulumi.com/docs/install/), and (for a full build) [.NET SDK](https://dotnet.microsoft.com/download), Node.js, and Python as needed.
2. From the repository root:
   - `make build` — provider + all SDKs (slow first time).
   - `make provider` — provider binary only (faster iteration).
   - `make test_provider` — provider tests.

## Smoke / integration tests

End-to-end examples and environment variables are documented in [examples/SMOKE.md](examples/SMOKE.md). Integration tests under `examples/` are skipped unless `AKEYLESS_ACCESS_ID` and `AKEYLESS_ACCESS_KEY` are set.

## Pull requests

- Keep changes focused and describe **what** changed and **why** in the PR description.
- Run `make test_provider` (and relevant smoke tests if you touch examples) before submitting.

## Regenerating SDKs

If you change `provider/resources.go` or bridge metadata, run `make generate` (or the narrower `make tfgen` / language targets per `make help`) and commit the generated files as appropriate for your workflow.
