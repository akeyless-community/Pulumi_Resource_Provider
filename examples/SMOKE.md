# Akeyless provider smoke test

These examples create **one folder** in Akeyless (no static secrets). They verify the bridged provider end-to-end.

## What you need from Akeyless

1. **Access ID** and **Access Key** for an identity that can authenticate (e.g. API key auth method).
2. **Authorization on an item path** — in Akeyless, access is usually scoped by **role** to allowed paths (and capabilities such as create/update/delete on folders). The smoke test calls the **Folder Create** API, so the identity must be allowed to **create a folder under the path you choose** (exact permission names depend on your role template; typical failures are HTTP 403 or a 4xx with a “permission” / “access” message in logs).
3. Optional: custom gateway URL if not using the public SaaS API.

### Choosing the folder path (`AKEYLESS_SMOKE_PARENT_PATH`)

By default the program creates a folder named `pulumi-smoke-<stack-name>` at the **root** of your allowed namespace. Many accounts **do not** allow that; the role only permits writes **under a specific parent path** (for example a folder an admin created for automation).

Set **`AKEYLESS_SMOKE_PARENT_PATH`** to that **allowed parent item path** (no trailing slash required). The test creates:

`{AKEYLESS_SMOKE_PARENT_PATH}/pulumi-smoke-<stack-name>`

Example:

```bash
export AKEYLESS_SMOKE_PARENT_PATH="/my-tenant/automation/pulumi-smoke"
```

Work with your Akeyless admin to obtain a path your auth method’s **associated role** may use for **folder** operations, or adjust the role’s path rules in the console/docs to include this subtree.

Environment variables (same as the [Terraform provider](https://registry.terraform.io/providers/akeyless-community/akeyless/latest/docs)):

| Variable | Required | Description |
|----------|----------|-------------|
| `AKEYLESS_ACCESS_ID` | Yes | Access ID |
| `AKEYLESS_ACCESS_KEY` | Yes | Access key |
| `AKEYLESS_GATEWAY` | No | Defaults to `https://api.akeyless.io` |
| `AKEYLESS_SMOKE_PARENT_PATH` | **Often** (restricted roles) | Parent path under which `pulumi-smoke-<stack>` may be created |

## Prerequisite: build the repo

From the repository root:

```bash
make build
```

## TypeScript (`basic-ts`)

Pulumi’s `ProgramTest` copies the app to a temp directory, so **never** use a relative `file:../../...` for the SDK there. The Go test uses **`Overrides`** to rewrite `package.json` / Yarn resolutions to an absolute `file:/.../sdk/nodejs/bin` before `yarn install`.

`package.json` lists `@pulumi/akeyless` as `0.0.0` as a placeholder; integration tests replace it with the real path. For a **manual** run, point it at the built SDK:

```bash
cd examples/basic-ts
npm install
npm install "file:$(cd ../.. && pwd)/sdk/nodejs/bin" --save
export AKEYLESS_ACCESS_ID="..." AKEYLESS_ACCESS_KEY="..."
export AKEYLESS_SMOKE_PARENT_PATH="/path/your-role-may-write"   # if required by your role
pulumi stack init dev
pulumi up
pulumi destroy
```

(`make build` must have produced `sdk/nodejs/bin` first.)

## Python (`basic-py`)

Use a venv, then:

```bash
cd examples/basic-py
python3 -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate
pip install -r requirements.txt
export AKEYLESS_ACCESS_ID="..." AKEYLESS_ACCESS_KEY="..."
export AKEYLESS_SMOKE_PARENT_PATH="/path/your-role-may-write"   # if required by your role
pulumi stack init dev
pulumi up
pulumi destroy
```

`requirements.txt` installs `pulumi_akeyless` editable from `../../sdk/python`.

## Automated integration test (Go)

From `examples/`:

```bash
export AKEYLESS_ACCESS_ID="..."
export AKEYLESS_ACCESS_KEY="..."
export AKEYLESS_SMOKE_PARENT_PATH="/path/your-role-may-write"   # if your role is path-scoped
go test -v -tags=nodejs -timeout=60m .
```

Python:

```bash
go test -v -tags=python -timeout=60m .
```

Tests are **skipped** if the two `AKEYLESS_*` variables are unset.

The Go tests also run `pulumi config set-all --secret` with `akeyless:apiKeyLogins` built from those variables, so the provider gets credentials from **stack config** (not only the process environment). That avoids **401 Unauthorized** when the test runner does not forward `AKEYLESS_*` to child processes (common in IDEs). Optional `AKEYLESS_GATEWAY` is set as plaintext `akeyless:apiGatewayAddress` when present.

## Cleanup

`pulumi destroy` removes the folder. The leaf name is `pulumi-smoke-<stack>`; the full item path is `{AKEYLESS_SMOKE_PARENT_PATH}/pulumi-smoke-<stack>` when the parent path env var is set.
