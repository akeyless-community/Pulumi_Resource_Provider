---
title: Akeyless Installation & Configuration
meta_desc: Install the Akeyless Pulumi provider and configure gateway URL and credentials.
layout: package
---

## Installation

Use the package manager for your language (see the [package overview](..) for package names and links).

Ensure the Pulumi CLI can install the provider plugin: releases must publish `pulumi-resource-akeyless-v*-*-*.tar.gz` assets on GitHub for each OS and architecture you support (see the upstream repository release notes).

## Authentication

Most stacks use API key authentication against your Akeyless gateway.

### Example: API key via Pulumi config

Set the gateway origin (scheme, host, and optional port):

```bash
pulumi config set akeyless:apiGatewayAddress "https://api.akeyless.io"
```

Set API key login objects (use `--secret` for sensitive values):

```bash
pulumi config set --path 'akeyless:apiKeyLogins[0].accessId' "YOUR_ACCESS_ID"
pulumi config set --path 'akeyless:apiKeyLogins[0].accessKey' "YOUR_ACCESS_KEY" --secret
```

### Environment variables

For API key auth, the Terraform-compatible variables **`AKEYLESS_ACCESS_ID`** and **`AKEYLESS_ACCESS_KEY`** are supported. **`AKEYLESS_GATEWAY`** is commonly used for the same value as `apiGatewayAddress`.

## Configuration options

Common provider settings:

* `akeyless:apiGatewayAddress` (environment: `AKEYLESS_GATEWAY`, and related gateway URL variables per upstream provider) — base URL of the Akeyless API gateway.
* `akeyless:apiKeyLogins` — list of objects with `accessId` and `accessKey` for API key authentication.

The full configuration model (additional auth methods and optional fields) is defined on the `akeyless` provider resource in the generated API documentation once the package is listed in the Pulumi Registry.
