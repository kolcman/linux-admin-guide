# AGENTS.md

## Cursor Cloud specific instructions

This repo is a single Go module (`linux-guide`) — an offline terminal (TUI) cheat sheet
for Linux/DevOps commands built with Bubble Tea + Lip Gloss. There is no backend,
database, or web service; the only "service" is the interactive TUI binary itself.

Standard commands (see also `README.md` and `.github/workflows/release.yml`):

- Lint / static check: `go vet ./...`
- Build (dev): `go build -o linux-guide .`
- Run: `./linux-guide` (or `go run .`)

Notes:

- There are currently no automated tests (`*_test.go`) in the repo. `go test ./...`
  is a no-op ("no test files") but still safe to run.
- The app uses the terminal alt-screen (`tea.WithAltScreen()`) and requires an
  interactive TTY. It will not render usefully when stdout is not a terminal.
  To exercise it non-interactively, run it inside a `tmux` session and drive it with
  `tmux send-keys` / inspect with `tmux capture-pane`.
- Controls: arrows or `j`/`k` navigate, `Enter` opens, `Esc` goes back, `q` (or `й`)
  quits/back. Content menus are all in Russian.
- `install-plank-mint.sh` and `scripts/install.sh` are unrelated helper scripts (desktop
  dock setup / release-binary installer); they are not needed to build or run the app.
