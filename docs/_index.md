---
title: Akeyless
meta_desc: Pulumi provider for managing Akeyless secrets and configuration resources.
layout: package
---

## Installation

The Akeyless provider is published as language SDKs alongside the `pulumi-resource-akeyless` plugin (downloaded automatically by the Pulumi CLI from GitHub Releases when `pluginDownloadURL` is set in the package schema).

* JavaScript/TypeScript: [`@pulumi/akeyless`](https://www.npmjs.com/package/@pulumi/akeyless)
* Python: [`pulumi_akeyless`](https://pypi.org/project/pulumi-akeyless/) — `pip install pulumi_akeyless`
* Go: [`github.com/akeyless-community/pulumi-akeyless/sdk/go/akeyless`](https://github.com/akeyless-community/pulumi-akeyless/tree/main/sdk/go/akeyless)
* .NET: [`Pulumi.Akeyless`](https://www.nuget.org/packages/Pulumi.Akeyless)

## Overview

The Akeyless provider for Pulumi manages resources exposed by the [Akeyless](https://www.akeyless.io/) platform. It is bridged from the open-source [terraform-provider-akeyless](https://github.com/akeyless-community/terraform-provider-akeyless) (MPL 2.0). Report issues in [`akeyless-community/pulumi-akeyless`](https://github.com/akeyless-community/pulumi-akeyless/issues) first.

## Authentication

Configure API gateway URL and credentials (for example API key access ID and secret). See [Installation & configuration](installation-configuration/) for `pulumi config` examples and environment variables.

After the package appears in the [Pulumi Registry](https://www.pulumi.com/registry/), the generated API reference for the provider resource lists every supported auth block (AWS IAM, Azure AD, GCP, JWT, and others).
