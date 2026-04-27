# Pulumi Akeyless provider

[Pulumi](https://www.pulumi.com/) provider for managing [Akeyless](https://www.akeyless.io/) resources. It is bridged from the open-source [terraform-provider-akeyless](https://github.com/akeyless-community/terraform-provider-akeyless) (Terraform provider API and behavior; MPL 2.0). Report provider-specific issues in this repository first; for upstream behavior gaps, see that repository.

## Source, security, and contributing

| | |
|--|--|
| **Repository** | [github.com/akeyless-community/pulumi-akeyless](https://github.com/akeyless-community/pulumi-akeyless) |
| **Security** | [SECURITY.md](SECURITY.md) (use GitHub private vulnerability reporting when possible) |
| **Contributing** | [CONTRIBUTING.md](CONTRIBUTING.md) |
| **Code of conduct** | [CODE-OF-CONDUCT.md](CODE-OF-CONDUCT.md) |
| **Releasing** | [RELEASING.md](RELEASING.md) (tag, GitHub Release, tokens, registry PR) |

### Go module path

The Go SDK import path is **`github.com/akeyless-community/pulumi-akeyless/sdk/go/akeyless`**, matching this repository (**[akeyless-community/pulumi-akeyless](https://github.com/akeyless-community/pulumi-akeyless)**).

## Requirements

- [Go](https://go.dev/dl/) (see `provider/go.mod` / `sdk/go.mod` for supported versions)
- [Pulumi CLI](https://www.pulumi.com/docs/install/)
- Optional: [mise](https://mise.jdx.dev/) for pinned tooling used by `make lint_provider`

## Installing (published packages)

**Status:** SDK packages are **not** published to npm / PyPI / NuGet under the names below yet. Use [Building from source](#building-from-source) and the [examples](#examples-and-smoke-testing) until you run a release. When published, install with:

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

The schema also supports other authentication blocks (AWS IAM, Azure AD, GCP, JWT, certificate, token, Universal Identity, email). After the package is listed on the Pulumi Registry, see the [provider API reference](https://www.pulumi.com/registry/packages/akeyless/api-docs/provider/) for the full configuration model; until then, use the [Terraform provider argument reference](https://registry.terraform.io/providers/akeyless-community/akeyless/latest/docs) as a parallel guide.

Environment variable names for API key auth (**`AKEYLESS_ACCESS_ID`**, **`AKEYLESS_ACCESS_KEY`**) match the [Akeyless Terraform provider documentation](https://registry.terraform.io/providers/akeyless-community/akeyless/latest/docs).

## Examples and smoke testing

Under `examples/`:

- **`basic-ts`** — TypeScript stack that creates a test folder in Akeyless.
- **`basic-py`** — Same for Python.

End-to-end notes, required permissions, and automation with **`AKEYLESS_SMOKE_PARENT_PATH`** (for path-scoped roles) are documented in [examples/SMOKE.md](examples/SMOKE.md).

Quick manual check after `make build` (the example depends on the **local** SDK at `sdk/nodejs/bin`, not the npm registry):

```bash
cd examples/basic-ts
npm install
export AKEYLESS_ACCESS_ID="..." AKEYLESS_ACCESS_KEY="..."
# If your role only allows writes under a specific folder:
export AKEYLESS_SMOKE_PARENT_PATH="/your/allowed/parent/path"
pulumi stack init dev
pulumi up
pulumi destroy
```

Integration tests (Go) live in `examples/`; they are skipped unless `AKEYLESS_ACCESS_ID` and `AKEYLESS_ACCESS_KEY` are set. See [examples/SMOKE.md](examples/SMOKE.md) for `go test` invocations.

## Releases, GitHub assets, and Pulumi Registry

**Step-by-step (push, tag, CI upload, npm/PyPI/NuGet tokens):** [RELEASING.md](RELEASING.md). Pushing a tag like `v0.1.0` triggers [.github/workflows/release.yml](.github/workflows/release.yml), which builds and **uploads the six provider `.tar.gz` files** to that GitHub Release automatically.

The schema sets **`pluginDownloadURL`** to `github://api.github.com/akeyless-community/pulumi-akeyless`, so the Pulumi CLI downloads provider binaries from **GitHub Releases** on this repository (CLI 3.35.3+). The release **tag** must be **`vMAJOR.MINOR.PATCH`** (for example **`v0.1.0`**), matching the version baked into the published SDKs and plugin archives without the leading `v` in the archive basename.

### GitHub Release assets (example `v0.1.0`)

From a clean tree at the release commit, build archives (**version must not include a leading `v` in `PROVIDER_VERSION`**):

```bash
make provider_dist PROVIDER_VERSION=0.1.0
```

Attach **all** of these files from `bin/` to the GitHub Release for tag `v0.1.0` (six platforms):

| Asset |
|-------|
| `pulumi-resource-akeyless-v0.1.0-linux-amd64.tar.gz` |
| `pulumi-resource-akeyless-v0.1.0-linux-arm64.tar.gz` |
| `pulumi-resource-akeyless-v0.1.0-darwin-amd64.tar.gz` |
| `pulumi-resource-akeyless-v0.1.0-darwin-arm64.tar.gz` |
| `pulumi-resource-akeyless-v0.1.0-windows-amd64.tar.gz` |
| `pulumi-resource-akeyless-v0.1.0-windows-arm64.tar.gz` |

Each archive contains the `pulumi-resource-akeyless` binary at the root (plus `README.md` and `LICENSE` per the Makefile), which matches [Pulumi’s executable plugin layout](https://www.pulumi.com/docs/iac/guides/building-extending/packages/executable-plugin/). Optional extras such as a `SHA256SUMS` file are not required for the CLI.

**Suggested release order:** (1) set `PROVIDER_VERSION=0.1.0` and run `make generate` (or `make build`) so SDK metadata matches; (2) commit generated outputs; (3) tag `v0.1.0` and push; (4) run `make provider_dist PROVIDER_VERSION=0.1.0` and upload the six `.tar.gz` files; (5) publish npm / PyPI / NuGet packages (requires maintainer credentials; the generated Node and .NET package names are `@pulumi/akeyless` and `Pulumi.Akeyless` — confirm you can publish to those registries or change the package names in `provider/resources.go` and regenerate); (6) for Go, consumers use the **`sdk` module**, for example `go get github.com/akeyless-community/pulumi-akeyless/sdk@v0.1.0`.

### Listing on the Pulumi Registry

This repo includes **`docs/_index.md`** and **`docs/installation-configuration.md`** for registry docs. To appear on [pulumi.com/registry](https://www.pulumi.com/registry/), open a pull request against **[pulumi/registry](https://github.com/pulumi/registry)** that adds one object to `community-packages/package-list.json` (alphabetical order by `repoSlug` is conventional), for example:

```json
{
  "repoSlug": "akeyless-community/pulumi-akeyless",
  "schemaFile": "provider/cmd/pulumi-resource-akeyless/schema.json"
}
```

Pulumi’s team reviews community additions; you do not need hosting beyond GitHub.

## Continuous integration

Build and test locally (`make build`, `make test_provider`, [examples/SMOKE.md](examples/SMOKE.md)). **Tag pushes** run [`.github/workflows/release.yml`](.github/workflows/release.yml) to attach provider binaries to GitHub Releases. Add other workflows if you want pull-request CI or SDK publishing.

## Repository layout

- **`provider/`** — Bridge and provider implementation; `schema.json` and the `pulumi-resource-akeyless` binary.
- **`sdk/`** — Generated Node.js, Python, Go, and .NET SDKs.
- **`examples/`** — Sample programs and integration tests.

## License

This repository is licensed under the Apache License 2.0; see [LICENSE](LICENSE). The bridged Terraform provider is MPL 2.0; attribution and upstream links are noted in generated package metadata where applicable.
