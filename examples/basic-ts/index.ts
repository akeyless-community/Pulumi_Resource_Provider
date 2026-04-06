import * as pulumi from "@pulumi/pulumi";
import * as akeyless from "@pulumi/akeyless";

/**
 * Smoke test: creates a single Akeyless folder (no secrets).
 *
 * Auth (same as Terraform provider):
 *   export AKEYLESS_ACCESS_ID="p-..."
 *   export AKEYLESS_ACCESS_KEY="..."
 * Optional:
 *   export AKEYLESS_GATEWAY="https://api.akeyless.io"
 *
 * Path / RBAC:
 *   Most roles are limited to a subtree. Set AKEYLESS_SMOKE_PARENT_PATH to an item path your
 *   identity may manage (e.g. /dev/pulumi-tests). The folder will be created at
 *   {parent}/pulumi-smoke-{stack}. See examples/SMOKE.md.
 */
const stack = pulumi.getStack();
const parent = (process.env.AKEYLESS_SMOKE_PARENT_PATH ?? "").trim().replace(/\/+$/, "");
const leaf = `pulumi-smoke-${stack}`;
const folderName = parent ? `${parent}/${leaf}`.replace(/\/+/g, "/") : leaf;

const folder = new akeyless.Folder("smoke", {
    name: folderName,
    description: "Pulumi Akeyless provider smoke test; safe to delete.",
    deleteProtection: "false",
});

export const folderNameOut = folder.name;
export const folderIdOut = folder.folderId;
