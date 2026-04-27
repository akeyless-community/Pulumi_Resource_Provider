# Security policy

## Supported versions

Security fixes are applied to the **default branch** (`main`) and, when applicable, backported to the latest **release tag**.

## Reporting a vulnerability

Please **do not** open a public GitHub issue for security-sensitive reports.

Use one of the following:

1. **GitHub private vulnerability reporting** (recommended): open the repository on GitHub → **Security** → **Report a vulnerability**.
2. If private reporting is disabled for the org, contact the **akeyless-community** maintainers through your organization’s usual security channel.

Include enough detail to reproduce or assess impact (affected component, version or commit, and steps or proof-of-concept if safe to share).

## Scope

This repository contains the **Pulumi provider for Akeyless** (bridge + SDKs). Vulnerabilities in **upstream** components ([terraform-provider-akeyless](https://github.com/akeyless-community/terraform-provider-akeyless), Akeyless APIs, or Pulumi core) should be reported to those projects when appropriate.
