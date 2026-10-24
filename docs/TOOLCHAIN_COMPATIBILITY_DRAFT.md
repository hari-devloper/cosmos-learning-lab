# Cosmos Toolchain Compatibility — Draft

Status: **Evidence collected; no final dependency lock approved.**

Research date: 2026-10-24. This is a compatibility investigation, not a claim that a complete application build has passed.

## Why the version choice matters

The SDK, CometBFT, IBC-Go, and optional CosmWasm dependencies must be selected as a coherent release family. Do not force a newer SDK into an older IBC integration with ad hoc `replace` directives. A module's declared dependency is useful evidence, but a complete application build and relevant tests are still required.

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

## Initial interpretation

1. The SDK v0.55.0 release uses the CometBFT v0.40.0 line. It is a distinct dependency family from SDK v0.54.4, which declares CometBFT v0.39.4.
2. The inspected IBC-Go v11.2.0 and wasmd v0.70.5 module files declare SDK v0.54.0 and CometBFT v0.39.0. They do **not** by themselves establish compatibility with SDK v0.55.0 / CometBFT v0.40.0.
3. SDK v0.54.4 is a plausible candidate for an IBC-capable learning baseline because IBC-Go v11.2.0 declares SDK v0.54.0 and the SDK v0.54.x line uses CometBFT v0.39.x. This is a **candidate only**; validate module resolution, application build, tests, and official compatibility guidance before approval.
4. SDK v0.55.0 may be evaluated as a separate SDK-only learning track. Do not add IBC-Go v11.2.0 or wasmd v0.70.5 to that track unless an officially supported compatible release family is verified.
5. CosmWasm is optional and should remain excluded from the initial dependency set unless a concrete learning or product requirement justifies the additional runtime and compatibility burden.

## Proposed validation procedure (not yet run)

After selecting a candidate baseline:

1. Review official upgrade/compatibility guidance and relevant release notes.
2. Create a minimal disposable application using the selected release family.
3. Record `go version`, `go env`, exact module graph, and checksums.
4. Run `go mod tidy` and inspect unexpected major/minor version selection or replacements.
5. Build the application and run all available unit/integration tests.
6. Run a local-chain smoke test and any applicable IBC/wasm tests if those components are included.
7. Record exact commands and observed results; do not convert expected results into claimed observations.
8. Review the dependency decision and receive explicit approval before using it in the product repository.

## Decision gate

No dependency lock, starter application, or implementation work is approved by this draft. The next decision is whether the initial learning path prioritizes an IBC-capable SDK v0.54.x release family or studies SDK v0.55.x separately first.
