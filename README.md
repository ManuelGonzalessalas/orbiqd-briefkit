# OrbiqD BriefKit

OrbiqD BriefKit is a solution that exposes an **MCP server** to drive **subscription-based coding CLIs** (Codex / Claude Code / Gemini) — **no APIs, no API keys**.

It’s built for workflows where you want multiple coding agents to work on the same problem, while keeping a clean, explicit session log you can reuse, inspect, and summarize.

You can use BriefKit in two ways:
- as an MCP server (via the available tools),
- via the `briefkit-ctl` CLI.

## What it does

BriefKit runs locally and provides two operating modes:

### 1) Direct Agent Ask
Ask a single agent a question and get a response back.

Use it when you already know *who* should answer (e.g. “Claude Code: review this approach”, “Codex: propose a refactor”, “Gemini: generate tests”).

### 2) BriefKit Collaboration
Let agents collaborate on an idea inside a **shared session**.

## Why

Most multi-model tools rely on APIs. BriefKit is for people who already use coding agents via their CLIs and want:
- a shared session context,
- structured collaboration rather than copy-pasting between terminals.

## Configuration

BriefKit reads global configuration from `~/.orbiqd/briefkit/config.yaml`.

### Agents

Agent definitions live in `~/.orbiqd/briefkit/agents/*.yaml`. On startup, BriefKit scans that directory and loads every YAML file. The agent id is derived from the file name (for example `codex.yaml` -> `codex`).

Minimal agent schema:

```yaml
kind: codex | claude-code | gemini
executable:
  path: /path/to/binary
  environment:
    # environment variable overrides
```
