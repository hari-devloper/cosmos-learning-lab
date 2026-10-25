# Cosmos Toolchain Compatibility — Draft

Status: **Evidence collected; no final dependency lock approved.**

Research date: 2026-10-24. This is a compatibility investigation, not a claim that a complete application build has passed.

## Why the version choice matters

The SDK and CometBFT must form a compatible release family for the initial single-chain learning application. IBC-Go and optional CosmWasm dependencies will be evaluated in a separate later-stage compatibility gate because the project explicitly defers IBC implementation. Do not force a newer SDK into an older integration with ad hoc `replace` directives. A module's declared dependency is useful evidence, but a complete application build and relevant tests are still required.

## Primary-source evidence inspected

The following official `go.mod` files were fetched from their exact release tags:

| Component | Tag | Declared Go version | Direct ecosystem dependency evidence |
|---|---|---:|---|
| Cosmos SDK | `v0.55.0` | `1.26.5` | CometBFT `v0.40.0` |
| Cosmos SDK | `v0.54.4` | `1.26.5` | CometBFT `v0.39.4` |
| CometBFT | `v0.39.4` | `1.25.0` | Exact tag inspected |
| IBC-Go | `v11.2.0` | `1.25.9` | Cosmos SDK `v0.54.0`; CometBFT `v0.39.0` |
| wasmd | `v0.70.5` | `1.25.9` | Cosmos SDK `v0.54.0`; CometBFT `v0.39.0`; wasmvm/v3 `v3.0.9` |

### Exact source links

- Cosmos SDK v0.55.0: https://github.com/cosmos/cosmos-sdk/blob/v0.55.0/go.mod
- Cosmos SDK v0.54.4: https://github.com/cosmos/cosmos-sdk/blob/v0.54.4/go.mod
- Cosmos SDK v0.55.0 release notes: https://github.com/cosmos/cosmos-sdk/blob/v0.55.0/RELEASE_NOTES.md
- Cosmos SDK v0.54.4 release page: https://github.com/cosmos/cosmos-sdk/releases/tag/v0.54.4
- CometBFT v0.39.4: https://github.com/cometbft/cometbft/blob/v0.39.4/go.mod
- IBC-Go v11.2.0: https://github.com/cosmos/ibc-go/blob/v11.2.0/go.mod
- wasmd v0.70.5: https://github.com/CosmWasm/wasmd/blob/v0.70.5/go.mod

## Follow-up official release review — 2026-10-25

The official [Cosmos SDK releases page](https://github.com/cosmos/cosmos-sdk/releases) lists SDK `v0.55.0` and `v0.54.4` among the current release entries reviewed for this decision. The official `v0.55.0` notes state that moving from `v0.54.x` requires a coordinated chain upgrade. The official `v0.54.4` notes describe important security fixes and call the release state-breaking; they include fixes in transaction signature handling, multisig bounds checks, and other areas. See [v0.54.4 release notes](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.54.4) and [SDK upgrade reference](https://github.com/cosmos/cosmos-sdk/blob/main/UPGRADING.md).

This supports treating `v0.54.4` as a serious patch-level candidate, but it does **not** prove it is universally “the most mature” release or automatically suitable for every application. The state-breaking note is a reason to understand upgrade/state implications even for a learning chain. The dependency graph, app build, tests, and scope-specific compatibility remain required validation.

## Initial interpretation

1. The SDK v0.55.0 release uses the CometBFT v0.40.0 line. It is a distinct dependency family from SDK v0.54.4, which declares CometBFT v0.39.4.
2. The inspected IBC-Go v11.2.0 and wasmd v0.70.5 module files declare SDK v0.54.0 and CometBFT v0.39.0. They do **not** by themselves establish compatibility with SDK v0.55.0 / CometBFT v0.40.0.
3. SDK v0.54.4 is a plausible candidate for the single-chain learning baseline and may simplify a later IBC evaluation because IBC-Go v11.2.0 declares SDK v0.54.0 and the SDK v0.54.x line uses CometBFT v0.39.x. This is a **candidate only**; validate official compatibility guidance, module resolution, application build, tests, and release notes before approval.
4. SDK v0.55.0 may be evaluated as a separate SDK-only track. Do not add IBC-Go v11.2.0 or wasmd v0.70.5 to that track unless an officially supported compatible release family is verified.
5. CosmWasm is optional and remains excluded from the initial dependency set unless a concrete learning or product requirement justifies the added runtime and compatibility burden.
6. **IBC is a later-stage topic, not a reason to prematurely choose the initial SDK.** First select and validate a minimal single-chain SDK/CometBFT baseline. Reopen IBC-Go compatibility as its own explicit gate after the single-chain state machine, tests, read model, and reconciliation design have been reviewed.

## Proposed validation procedure (not yet run)

After selecting a candidate baseline:

1. Review official upgrade/compatibility guidance and relevant release notes.
2. Create a minimal disposable application using the selected release family.
3. Record `go version`, `go env`, exact module graph, and checksums.
4. Run `go mod tidy` and inspect unexpected major/minor version selection or replacements.
5. Build the application and run all available unit/integration tests.
6. Run a local-chain smoke test.
7. Record exact commands and observed results; do not convert expected results into claimed observations.
8. Review the dependency decision and receive explicit approval before using it in the product repository.
9. Later, when IBC becomes in-scope, verify the compatible IBC-Go release family and add the necessary tests in a separate decision.

## Learner decision — preferred candidate, validation pending

The learner chose the Cosmos SDK **v0.54.x family**, with **v0.54.4 / CometBFT v0.39.4** as the first candidate to validate, because the inspected dependency declarations appear more aligned with the intended later IBC compatibility path than the SDK v0.55.0 family. This is a scope and validation decision, not a claim that v0.54.4 is the newest or most mature release available as of the research date.

**Status: candidate selected; not yet approved as the project baseline.** Before approval, verify current official release/maintenance status and upgrade notes, inspect the exact tagged module graph, create a disposable minimal application, and run the documented build and tests. Do not add IBC-Go or CosmWasm to the initial dependency graph.

## Decision gate

No dependency lock, starter application, or implementation work is approved by this draft. The next decision is to choose a candidate SDK/CometBFT release family for the **single-chain learning scope**, verify its source and module graph, and validate it in a disposable application. IBC-Go and CosmWasm are not part of the initial dependency graph unless the user explicitly reopens that scope.

## Learner-run tag and Go metadata check — 2026-10-25

Commands were run in Ubuntu 24.04.5 WSL. Shared output:

- `v0.54.4` tag resolves to commit `0019e50dd23e48de820cd9273e84d140334ed2a1`.
- `v0.55.0` tag resolves to commit `69dcb02ab9dbe6340c722266b3c52d32f260ac67`.
- SDK `v0.54.4` declares `go 1.26.5`.
- SDK `v0.54.4` declares `github.com/cometbft/cometbft v0.39.4` and `github.com/cometbft/cometbft-db v0.14.3` (indirect).
- Learner environment reported Go `1.26.10 linux/amd64`, which is newer than the declared Go version.
- The output shows a `replace` section exists, but the replacement targets were **not** captured by the shared output.

This confirms the exact tag references and selected module's declared Go/CometBFT versions. It does not yet establish the full resolved module graph, the effect of replacements, or build/test success.

### Follow-up: SDK replacement directives inspected — 2026-10-25

Learner shared the complete `replace` block from the downloaded SDK `v0.54.4` `go.mod`:

- `github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0` (comment: use Cosmos fork of keyring).
- `github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7` (comment: replace broken goleveldb).

The module file also contains a separate `retract` block for older SDK versions; it is not another replacement block. These are declarations in the SDK's own module file. Their effect on a consuming application's resolved module graph still needs to be inspected after a minimal disposable module is created.
