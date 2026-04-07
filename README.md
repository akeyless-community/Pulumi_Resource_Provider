# Pulumi Akeyless provider

[Pulumi](https://www.pulumi.com/) provider for managing [Akeyless](https://www.akeyless.io/) resources. It is bridged from the open-source [terraform-provider-akeyless](https://github.com/akeyless-community/terraform-provider-akeyless) (Terraform provider API and behavior; MPL 2.0). Report provider-specific issues here first; for upstream behavior gaps, see that repository.

## Requirements

- [Go](https://go.dev/dl/) (see `provider/go.mod` / `sdk/go.mod` for supported versions)
- [Pulumi CLI](https://www.pulumi.com/docs/install/)
- Optional: [mise](https://mise.jdx.dev/) for pinned tooling used by `make lint_provider`

## Installing (published packages)

When published to npm, PyPI, NuGet, and the Go module proxy, use:

### Node.js (TypeScript / JavaScript)

```bash
npm install @pulumi/akeyless
```

### Python

```bash
pip install pulumi_akeyless
```

### Go

```bash
go get github.com/akeyless-community/pulumi-akeyless/sdk/go/akeyless
```

Imports use the path:

`github.com/akeyless-community/pulumi-akeyless/sdk/go/akeyless/...`

### .NET

```bash
dotnet add package Pulumi.Akeyless
```

## Building from source

From the repository root:

```bash
make build
```

This builds the provider binary, generates or builds SDKs, and installs local SDK artifacts for testing. Other useful targets:

| Target | Purpose |
|--------|---------|
| `make help` | List targets |
| `make provider` | Provider binary only |
| `make generate` | Regenerate SDKs, schema, and docs (commit outputs as your workflow requires) |
| `make test_provider` | Provider unit tests |
| `make test` | Example integration tests (run `make build` first) |

Optional: install [mise](https://mise.jdx.dev/) so `make lint_provider` uses the linter version from `mise.toml`.

## Provider configuration

Configure the `akeyless` provider with Pulumi config (recommended for secrets) or environment variables. Common settings:

- **`akeyless:apiGatewayAddress`** — API gateway origin URL (scheme, host, port). Environment variable **`AKEYLESS_GATEWAY`** is often used for the same value (default in many setups is `https://api.akeyless.io`).
- **`akeyless:apiKeyLogins`** — List of objects with `accessId` and `accessKey` (mark `accessKey` as secret). Aligns with API key authentication in Akeyless.

The schema also supports other authentication blocks (AWS IAM, Azure AD, GCP, JWT, certificate, token, Universal Identity, email). See the [Pulumi Registry API docs](https://www.pulumi.com/registry/packages/akeyless/api-docs/) for the full provider configuration model.

Environment variable names for API key auth (**`AKEYLESS_ACCESS_ID`**, **`AKEYLESS_ACCESS_KEY`**) match the [Akeyless Terraform provider documentation](https://registry.terraform.io/providers/akeyless-community/akeyless/latest/docs).

## Examples and smoke testing

Under `examples/`:

- **`basic-ts`** — TypeScript stack that creates a test folder in Akeyless.
- **`basic-py`** — Same for Python.

End-to-end notes, required permissions, and automation with **`AKEYLESS_SMOKE_PARENT_PATH`** (for path-scoped roles) are documented in [examples/SMOKE.md](examples/SMOKE.md).

Quick manual check after `make build`:

```bash
cd examples/basic-ts
npm install
npm install "file:$(cd ../.. && pwd)/sdk/nodejs/bin" --save
export AKEYLESS_ACCESS_ID="..." AKEYLESS_ACCESS_KEY="..."
# If your role only allows writes under a specific folder:
export AKEYLESS_SMOKE_PARENT_PATH="/your/allowed/parent/path"
pulumi stack init dev
pulumi up
pulumi destroy
```

Integration tests (Go) live in `examples/`; they are skipped unless `AKEYLESS_ACCESS_ID` and `AKEYLESS_ACCESS_KEY` are set. See [examples/SMOKE.md](examples/SMOKE.md) for `go test` invocations.

## GitHub Actions CI

Boilerplate workflows originally called **Pulumi ESC** (`imports/github-secrets`) with OIDC to the `pulumi` organization. That only works inside Pulumi’s own infrastructure; community repos get **401 Unauthorized** on the token exchange. These workflows are patched to use **GitHub repository (or organization) secrets** and a **GitHub App installation token** instead.

Add these secrets so jobs such as **Upgrade provider**, **prerequisites**, **build**, and **lint** can authenticate to the API and clone private tooling:

| Secret | Purpose |
|--------|---------|
| `PULUMI_PROVIDER_AUTOMATION_APP_ID` | Numeric App ID for a GitHub App installed on this repository (or org). |
| `PULUMI_PROVIDER_AUTOMATION_PRIVATE_KEY` | PEM private key for that app (full key including `BEGIN` / `END` lines). |

Install the app with permissions appropriate for automation (at minimum: **Contents** and **Pull requests** read/write where you want upgrade PRs; **Metadata** read). See [Creating a GitHub App](https://docs.github.com/en/apps/creating-github-apps/registering-a-github-app).

Optional secrets for specific jobs: `CODECOV_TOKEN`; publishing (`PYPI_API_TOKEN`, `NPM_TOKEN`, `NUGET_PUBLISH_KEY`, Java/OSSRH-related names in `publish.yml`); `AZURE_SIGNING_*` for Windows signing; `RELEASE_BOT_ENDPOINT` / `RELEASE_BOT_KEY` only if you use Pulumi’s release-by-label integration.

## Repository layout

- **`provider/`** — Bridge and provider implementation; `schema.json` and the `pulumi-resource-akeyless` binary.
- **`sdk/`** — Generated Node.js, Python, Go, and .NET SDKs.
- **`examples/`** — Sample programs and integration tests.

## License

This repository is licensed under the Apache License 2.0; see [LICENSE](LICENSE). The bridged Terraform provider is MPL 2.0; attribution and upstream links are noted in generated package metadata where applicable.
