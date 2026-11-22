# LEARNINGS

Dated, append-only. Never rewritten; new discoveries go at the bottom. Decision-CM learnings land here tagged with the CM's name (`[<cm-name>]`).

## 2026-11-20 — Fresh start + organization installed

- Repo reset: everything pre-2026-11-20 moved to gitignored `archive/` — curriculum company content now cannot reach GitHub by construction.
- Manager & Context Masters organization installed. Roster empty; promotion counter `cm-candidates.md` seeded with 5 candidates at 0 triggers.
- New laws in AGENTS.md: **Docs-first** (explanations cite official documentation), **Review law** (domain CMs review code/learnings, parallel where independent; Manager verifies everything).

## 2026-11-20 — Direction set; go-language CM born

- Direction: **Learn Go with Tests** (quii) — TDD, one part at a time: build → question → answer → confirm, then next.
- Go target: **1.26.0** (owner-set). Local: go1.27.0 darwin/arm64, GOTOOLCHAIN=auto — a 1.27 toolchain compiles `go 1.26.0` modules with 1.26 semantics; the exact toolchain auto-downloads if ever pinned. No reinstall.
- Docs decision (re-validated): full go.dev mirror = 171 MB → rejected. CM stack: `go doc` (offline, version-exact, primary) + `/usr/local/go/doc/` (spec, memory model, godebug) + book clone `.knowledge/learn-go-with-tests/` (8.5 MB shallow, gitignored).
- First Context Master: `go-language` (decision-CM, isolated, glm-5.3-flash). Reviews Luis's code with citations; NEVER writes/edits `.go` files (Luis hand-types everything; Socratic).

## 2026-11-20 — vscode CM born (owner-directed); Block 0 staged

- Luis uses VS Code daily at work (Windows + **WSL Alpine** remote). 11 work-proven editor practices recorded as canon in the CM brain (`.pi/agents/vscode.md` § The work-proven canon) — incl. the `\\wsl$` ban ("squiggles day", 06.11.) and Thomas's workspace-only-extensions rule.
- **Drill law** added to the teaching contract: one random VS Code action per Go block (pool: brain § Drills), Luis performs + explains in German.
- Docs: curated official set (15 files, 372 KB) at `.knowledge/vscode-docs/` — remote/wsl/ssh/server architecture, extensions/workspaces, keybindings (`default-keybindings.md` holds command IDs; rendered chords are `kb()` placeholders), terminal, debugging + `vscode-go-debugging.md`.
- Mac keymap delta is EXPECTED only (Cmd↔Ctrl swaps; Ctrl+` stays) — must be verified against the installed app in Block 0 before drilling.
- VS Code not installed on this Mac; brew available → `brew install --cask visual-studio-code`.

## 2026-11-22 — [vscode] Block 0: Mac default keymap VERIFIED against installed app (VS Code 1.147.0)

- App confirmed on disk: `code --version` (full path `/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code`) → `1.147.0 / 8a14f0cc1b08d531eb83255ea3764ae22aa932b5 / arm64`. CLI still not on PATH.
- Method: grepped the compiled keybinding rules in the app bundle (`Contents/Resources/app/out/vs/workbench/workbench.desktop.main.js`) and decoded the numeric `primary`/`secondary` fields. KeyMod: CtrlCmd=2048, Shift=1024, Alt=512, WinCtrl=256 (= physical ⌃ on macOS). Decoding self-validated via anchors matching public defaults: ⌘P=2094, ⌘⇧F=3108, ⌘⇧M=3115, zoomOut ⌘-=2136 (→ US_MINUS=88), jumpToBracket ⌘⇧\=3165 (→ Backslash=93), terminal.new ⌃⇧`=1371 (→ Backquote=91), navigateBack win Alt+←=527 (→ LeftArrow=15).
- VERIFIED Mac table for the canon actions (primary unless noted):

  | Action | command id | Mac default | bundle evidence |
  |---|---|---|---|
  | Go to Definition | `editor.action.revealDefinition` | **F12** | primary:70, no mac override (plus hidden ⌘F12 variant primary:2118; Ctrl/Cmd+Click is the mouse binding, not in keybinding rules) |
  | Go Back | `workbench.action.navigateBack` | **⌃- (Control+Minus)** | mac primary:344 = WinCtrl\|88 (win Alt+←:527 unchanged) |
  | Quick Open (file by name) | `workbench.action.quickOpen` | **⌘P** | primary:2094 = CtrlCmd\|KeyP(46), mac same |
  | Find in Files | `workbench.action.findInFiles` | **⌘⇧F** | primary:3108 = CtrlCmd\|Shift\|KeyF(36), no mac override |
  | Show Problems | `workbench.actions.view.problems` | **⌘⇧M** | primary:3115 = CtrlCmd\|Shift\|KeyM(43), no mac override |
  | Toggle Terminal | `workbench.action.terminal.toggleTerminal` | **Control+Backquote** (canon "Ctrl+` stays" HOLDS) | mac:347 = WinCtrl\|Backquote(91); win Ctrl+`:2139 |
  | Split Terminal | `workbench.action.terminal.split` | **⌘\** | mac:2141 = CtrlCmd\|Backslash(93); win ⌃⇧5:3098; mac secondary [1306] = ⌃⇧5 |

- Keymap deltas vs the Windows canon for drills: Go Back is ⌃- (NOT Alt+← — Alt+← is Windows-only); Split Terminal is ⌘\ (NOT Ctrl+Shift+5 — ⌃⇧5 survives only as mac secondary); everything else maps as expected (Cmd↔Ctrl swap; terminal toggle keeps physical ⌃).
- navigateBack/navigateForward `secondary:[122]/[123]` = mouse side-buttons (Back/Forward), not keyboard keys — do not decode as chords.
- Cross-check: command ids present in `.knowledge/vscode-docs/default-keybindings.md` for revealDefinition, quickOpen, navigateBack, view.problems, toggleTerminal (chords are `kb()` placeholders there — bundle is the chord ground truth); ids `findInFiles` and `terminal.split` are NOT listed in that file → verified from the bundle alone.
- Open: optional belt-and-braces in-app check (Command Palette → "Preferences: Open Default Keyboard Shortcuts (JSON)") before first drill; `code` on PATH still to be added by owner.

## 2026-11-22 — Block 0 closed; Block 1 opened

- `code` CLI on PATH: symlink `/opt/homebrew/bin/code` → app-bundle binary. (Owner's first attempt created `code~` by typo; Manager removed and relinked.) `code --version` → 1.147.0, verified in the owner's terminal.
- Block 1 (Hello, world) opened; instruction grounded from `.knowledge/learn-go-with-tests/hello-world.md`. Chapter folder: `hello/` at repo root.

## 2026-11-22 — [go-language] The `cd .. && go mod init` trap; one module per chapter folder

- Luis hit `go: cannot find main module, but found .git/config in /Users/luiszg/Desktop/GitHub/go_learnings` running bare `go test` in `hello/` (no go.mod anywhere in the tree; archive/ copies are gitignored history). Root cause: the go command (module mode, default since Go 1.16 — book cites the go116 module changes blog in `hello-world.md` § Go modules?) walks up parent directories looking for `go.mod`; finding none, it anchors its init hint on the nearest VCS boundary it finds — the repo `.git`. The suggestion is generic, NOT layout advice.
- Book prescription (`.knowledge/learn-go-with-tests/hello-world.md` § Go modules?): run `go mod init example.com/hello` IN the chapter folder; and: "In future chapters, you will need to run `go mod init SOMENAME` in each new folder" — i.e. one module per chapter folder, name must contain a dot. Confirmed by sandbox: same files without a parent .git give the plainer message `go: go.mod file not found in current directory or any parent directory; see 'go help modules'`.
- Repo-root module REFUTED for this layout (sandboxed mimic): it mechanically passes (`ok github.com/luiszg/go_learnings/hello`) but swallows all chapters into ONE module (`…/go_learnings/integers`, …) — against the book's per-folder prescription and `go help modules` semantics ("a module is a collection of packages that are released, versioned, and distributed together" — chapters are independent units).
- Toolchain gotcha: on go1.27.0, `go mod init example.com/hello` writes `go 1.27.0` into the fresh go.mod (NOT the owner-set 1.26.0). Fix verified: `go mod edit -go=1.26.0` → go.mod declares `go 1.26.0`, `go test` PASSes under GOTOOLCHAIN=auto with 1.26 semantics. Future chapter init commands should include this step.
- Book-faithful Block 1 fix (all inside `hello/`): `go mod init example.com/hello` → `go mod edit -go=1.26.0` → `go test` → `PASS ok example.com/hello`.
- **Owner decision (2026-11-22): book-verbatim mode.** NO `go mod edit -go=` version corrections — the book's examples carry no version pinning, so neither do we. `go mod init` stamps the local toolchain version (1.27.0 here); the 1.26.0 "target" stays informational only. CM's two-step recommendation is thereby overridden; `go mod init SOMENAME` + nothing else is the per-chapter ritual.

## 2026-11-22 — [go-language] `go` directive = language gate + toolchain selector; GOTOOLCHAIN forms verified live on this machine

- Language gating VERIFIED live (go1.27.0 toolchain, /tmp module): with `go 1.18` in go.mod, code using builtin `min` fails with `built-in min requires go1.21 or later (-lang was set to go1.18; check go.mod)`; after `go mod edit -go=1.26.0` the IDENTICAL code builds with the SAME toolchain. Ground truth: go.dev/ref/mod § "go directive" — "the compiler rejects use of language features introduced after the version specified by the go directive" + "Before Go 1.21, the directive was advisory only; now it is a mandatory requirement: Go toolchains refuse to use modules declaring newer Go versions."
- GOTOOLCHAIN form semantics VERIFIED live, using a downloaded-and-cached go1.25.6 toolchain on a module with `go 1.26.0`:
  - `GOTOOLCHAIN=go1.25.6+auto go build ./...` → `go: downloading go1.26.0 (darwin/arm64)`, exit 0 — an older default toolchain transparently fetches and runs a newer one. (= what a colleague with only Go 1.25 installed experiences under the standard default `auto`.)
  - bare `GOTOOLCHAIN=go1.25.6 go build ./...` → `go: go.mod requires go >= 1.26.0 (running go 1.25.6; GOTOOLCHAIN=go1.25.6)`, exit 1 — bare `<name>` means "exactly that toolchain, never switch" (go.dev/doc/toolchain: "the go command always runs that specific Go toolchain").
  - `GOTOOLCHAIN=local` with bundled 1.27.0 on the same module → build OK (local toolchain ≥ go line is fine; switching only ever goes UP, never down).
- GOTOOLCHAIN default chain (go.dev/doc/toolchain): `auto` is shorthand for `local+auto`; standard toolchains set `GOTOOLCHAIN=auto` via `$GOROOT/go.env`; with `local` the go command "will refuse to run in a main module … in which the go line requires a newer version of Go". `go help environment` GOTOOLCHAIN entry points to https://go.dev/doc/toolchain.
- Toolchains are distributed as ordinary modules: `golang.org/toolchain@v0.0.1-go<VERSION>.<GOOS>-<GOARCH>` under `$GOPATH/pkg/mod/` — this machine now caches go1.25.6 and go1.26.0 (darwin/arm64) alongside the bundled go1.27.0.
- `go mod edit -go=` semantics (`go help mod edit`): "-go=version flag sets the expected Go language version … takes a version like '1.26' or '1.26.2'. Using the version 'none' removes the go directive. Users should prefer 'go get go@version'."
- ref/mod's `go mod init` section does NOT spell out which version init writes — on this machine the empirical answer is stable: init writes the running toolchain's version (`go 1.27.0` on go1.27.0). Keep the init→`go mod edit -go=1.26.0` two-step in chapter scripts.
