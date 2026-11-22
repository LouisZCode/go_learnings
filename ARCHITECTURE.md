# go_learnings — system map

> Append-only phase ledger at the bottom. Never delete phases.

## System map

- **Purpose:** Luis learns Go from zero; agents teach (contract in AGENTS.md).
- **Layout:** repo root holds the organization files; all pre-2026-11-20 material lives in gitignored `archive/` (reference only, never built, never pushed).
- **Learning path:** "Learn Go with Tests" (quii) — chapter folders appear at repo root as Luis builds them. Book source: `.knowledge/learn-go-with-tests/` (gitignored, local only, 8.5 MB). Target Go 1.26.0 (local go1.27.0 + GOTOOLCHAIN=auto).
- **Organization:** Manager & Context Masters pattern since 2026-11-20 (roster in AGENTS.md, promotion counter in cm-candidates.md).

## Phase ledger

- **2026-11-20 · Phase 0 — Fresh start.** All prior learning material (00_intro, 00_basic_http, listapp, listapp-evolution, curriculum, docs, media, old AGENTS/LEARNINGS) moved to `archive/`; `archive/` gitignored — company content can never be committed.
- **2026-11-20 · Phase 1 — Organization installed.** Manager & Context Masters pattern (skill: `manager-and-context-master-organization`). Roster empty (earned, not planned). Custom laws added: Docs-first, Review law.
- **2026-11-20 · Phase 2 — Direction set; first CM promoted.** Learning path: Learn Go with Tests, chapter by chapter (TDD, confirm each part). `go-language` Context Master created (decision-CM, brain `.pi/agents/go-language.md`, glm-5.3-flash) — docs access: `go doc`, `/usr/local/go/doc/`, book clone. Target Go 1.26.0.
- **2026-11-22 · Phase 3a — Block 0 closed.** VS Code 1.147.0 installed (owner); `code` CLI on PATH via symlink `/opt/homebrew/bin/code` → app bundle. Mac keymap VERIFIED by `vscode` CM from the app bundle (Go Back = ⌃-, Split Terminal = ⌘\, rest as expected — LEARNINGS.md). Drills armed.
- **Phase 3b — in progress: Block 1 (book chapter: Hello, world).**
- **2026-11-20 · Phase 2+ — `vscode` CM promoted (owner-directed).** Second CM (decision-CM, brain `.pi/agents/vscode.md`, glm-5.3-flash). Curated official docs (15 files, 372 KB) at `.knowledge/vscode-docs/` (remote/wsl/ssh/server architecture, extensions/workspaces, keybindings, terminal, debugging, vscode-go). Work-proven canon (11 items) + drill pool embedded in brain; drill law added to teaching contract.
