# Copyright 2024, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.

"""Smoke test: one Akeyless folder. Set AKEYLESS_ACCESS_ID and AKEYLESS_ACCESS_KEY."""

import os
import re

import pulumi
import pulumi_akeyless as akeyless

stack = pulumi.get_stack()
_parent = os.environ.get("AKEYLESS_SMOKE_PARENT_PATH", "").strip().rstrip("/")
_leaf = f"pulumi-smoke-{stack}"
folder_name = re.sub(r"/+", "/", f"{_parent}/{_leaf}") if _parent else _leaf

folder = akeyless.Folder(
    "smoke",
    name=folder_name,
    description="Pulumi Akeyless provider smoke test; safe to delete.",
    delete_protection="false",
)

pulumi.export("folder_name", folder.name)
pulumi.export("folder_id", folder.folder_id)
