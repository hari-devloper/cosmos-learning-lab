# Development Environment Audit

Status: **Observed from the environment checks shared during setup. Re-run before implementation if the machine or toolchain changes.**

Audit date: 2026-10-24

## Host and shell

| Item | Observed value |
|---|---|
| Host OS | Windows 11 with WSL2 |
| Linux distribution | Ubuntu 24.04.5 LTS |
| Linux architecture | `x86_64` / `linux/amd64` |
| Shell | `/bin/bash` |
| Linux user | `harikrishna` |
| Working directory created | `/home/harikrishna/src` |
| Notes directory created | `/home/harikrishna/workspace-notes` |

## Resources observed during the initial audit

| Item | Observed value |
|---|---:|
| Logical CPUs visible | 20 |
| Memory | 7.6 GiB total; approximately 7.1 GiB available at audit time |
| Swap | 2 GiB |
| Filesystem containing WSL home | Approximately 1007 GiB total; approximately 955 GiB available at audit time |

Resource availability is time-dependent; these figures are a snapshot, not a capacity guarantee.

## Linux toolchain observed

| Tool | Observed version/path |
|---|---|
| Go | `go1.26.10 linux/amd64`, `/usr/local/go/bin/go` |
| Git | `git version 2.43.0`, `/usr/bin/git` |
| jq | `1.7` |
| GNU Make | `4.3` |
| GCC | `13.3.0` |
| curl | `8.5.0` |
| VS Code command | `code --version` reported `1.143.1 x64` |

The following shell checks were shared and observed:

```bash
/bin/bash
Linux
/usr/local/go/bin/go
go version go1.26.10 linux/amd64
/usr/bin/git
git version 2.43.0
```

## Important environment caveat

Earlier checks found that Node/npm and Docker commands from Ubuntu resolved to Windows executables under `/mnt/c/Program Files/`. Do not assume those are Linux-native WSL tools. Before using them in a Linux build or test, inspect `command -v`, `type -a`, and version output in the VS Code WSL remote terminal.

Rust was not installed at the time of the audit. Do not install it unless an approved dependency or lab requires it.

## Re-run audit commands

Run these in the **Ubuntu WSL terminal**, and later in the VS Code integrated terminal after confirming that VS Code is connected to WSL:

```bash
printf 'shell=%s\n' "$SHELL"
uname -a
uname -m
command -v go && go version
command -v git && git --version
command -v node || true
command -v npm || true
command -v docker || true
command -v rustc || true
df -h "$HOME"
free -h
nproc
```

For VS Code, confirm the window indicates `WSL: Ubuntu-24.04`. A PowerShell terminal or a Windows Go executable is not evidence that the integrated terminal is using the Linux toolchain.

## Not yet verified

- Linux-native Node.js/npm version and installation method
- Docker Desktop WSL integration and Linux-side Docker behavior
- VS Code remote window state at the latest check
- Rust toolchain
- Framework dependency graph and reproducible application build

## Next action

Verify the WSL remote terminal and inspect official framework release/module metadata before pinning application dependencies. Do not start application implementation based only on this environment audit.

## Follow-up audit — 2026-10-25

Learner-shared Ubuntu WSL terminal output:

| Check | Observed result | Assessment |
|---|---|---|
| Distribution | Ubuntu 24.04.5 LTS | Confirmed |
| Architecture | `x86_64` | Confirmed |
| Go | `go1.26.10 linux/amd64` at `/usr/local/go/bin/go` | Linux-native path confirmed |
| Go environment | `GOROOT=/usr/local/go`; `GOPATH` under the Linux home directory | Observed |
| Git | `git version 2.43.0` at `/usr/bin/git` | Linux-native path confirmed |
| Node.js | `node` command not found | Not available natively in this shell |
| npm | `10.8.2`, resolves under `/mnt/c/Program Files/nodejs/` | Windows executable; do not rely on it for Linux builds |
| Docker | Resolves under `/mnt/c/Program Files/Docker/`; Docker reports unavailable in this WSL distro | WSL integration not verified; unavailable from this shell |
| Working directory | `/home/harikrishna/src` | Observed |

### Decision

The Go/Git foundation is sufficient for read-only source and dependency investigation. Do not install Node.js, configure Docker, or begin application implementation as part of this gate. They are not needed for the next Cosmos SDK source/dependency review.

This audit does not establish a successful Cosmos SDK build, dependency compatibility, or runtime behavior. Next: review official version metadata and the provisional toolchain compatibility draft, then explicitly decide the dependency baseline before creating the runnable lab.
