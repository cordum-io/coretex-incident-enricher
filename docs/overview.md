# CoretexOS Pack Overview

This repo is a public, minimal example of how coretexOS is meant to scale:
keep the platform stable and install behavior as packs.

## What coretexOS provides (platform-only)

- CAP protocol (bus envelopes, job requests/results, heartbeats).
- Scheduler + safety kernel + workflow engine.
- Gateway APIs for workflows, artifacts, approvals, config, and policy.
- Redis-backed pointers for job context/results.

## What a pack adds

- Workflow templates (what to run, and in what order).
- Topics and routing overlays (topic -> pool mapping).
- Policy overlays (allow, deny, require approval).
- Schemas for inputs/outputs.
- External workers that implement the topics.

The platform never executes pack code during install. Packs only write data into
existing stores via the gateway.

## Why this pack exists

Incident Enricher is a first-pack reference. It proves:

- Workflow CRUD + run lifecycle works without any special-case product code.
- Policy fragments are loaded from config service and can force approvals.
- Scheduler routing overlays work (topic -> pool mapping).
- Artifacts provide an audit trail of evidence and outputs.

## What is intentionally minimal

- LLM integration is mock by default (OpenAI stub only).
- Incident fetch is mock by default (no PagerDuty integration).
- Slack posting works via webhook, gated by policy approval.

The goal is to keep the surface area small while proving the platform loop.
