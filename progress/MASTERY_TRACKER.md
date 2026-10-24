# Mastery Tracker

This tracker records evidence, not intentions. Mark an item complete only after the learner has demonstrated the relevant understanding and reproducible lab.

## Status key

- [ ] Not started
- [~] In progress (include evidence and remaining gaps)
- [x] Complete (must include evidence link or a concise evidence record)

## Stage 0 — Environment and compatibility

- [x] Record initial WSL/toolchain observations — `docs/ENVIRONMENT_AUDIT.md`
- [x] Record official-tag dependency evidence and compatibility questions — `docs/TOOLCHAIN_COMPATIBILITY_DRAFT.md`
- [ ] Confirm VS Code integrated terminal is connected to WSL and all required tools resolve to Linux binaries
- [ ] Decide initial Cosmos SDK release-family learning path
- [ ] Validate selected dependency graph in a disposable minimal application
- [ ] Run and record clean build/test evidence for the selected baseline

## Stage 1 — Distributed systems foundations

- [ ] Deterministic state machines
- [ ] Replicated state and consensus
- [ ] Transaction ordering and finality
- [ ] Safety, liveness, and fault assumptions
- [ ] Cryptographic keys, signatures, hashes, and encoding

## Stage 2 — Cosmos stack architecture

- [ ] Cosmos SDK application and module boundaries
- [ ] CometBFT / ABCI boundary
- [ ] Transaction submission and execution lifecycle
- [ ] Ante processing, messages, signers, and authorization
- [ ] State stores, caches, events, and queries
- [ ] Genesis, configuration, upgrades, and local networks


## Learner evidence log

### Module 5 — Cosmos deep architecture (2026-10-24)

- [x] Conceptual checkpoint: 3/3 correct. Evidence: Q1 identified the consensus/application-execution boundary; Q2 identified the security and operational responsibilities of an app-specific chain; Q3 distinguished CheckTx admission from commitment.
- [~] Source-level BaseApp and transaction call graph — next step; exact symbols and release tag still need interactive inspection.
- [ ] Runnable lab and failure experiment — not performed.
- [ ] Selected dependency baseline and application build/test — not approved or performed.
- [ ] Production or performance claims — none established.


### Module 5 source investigation — BaseApp (2026-10-24)

- [~] Source-level lifecycle investigation started against Cosmos SDK `v0.54.4`: inspected `CheckTx`, `RunTx`, `PrepareProposal`, `ProcessProposal`, `FinalizeBlock`, `internalFinalizeBlock`, `executeTxsWithExecutor`, `deliverTx`, `runMsgs`, `Commit`, and `MsgServiceRouter.Handler`. Evidence is documented in `docs/modules/MODULE-05-cosmos-stack-architecture-abci.md`.
- [ ] Complete module wiring and AnteHandler-chain trace.
- [ ] Runnable lab and failure experiment.
- [ ] Clean build/test against an approved dependency baseline.

## Mastery evidence template

For each item, record:

- **Concept ID:**
- **Learner explanation / diagram:**
- **Lab and exact commands:**
- **Observed result and evidence link:**
- **Failure experiment explained:**
- **Interview questions answered:**
- **Remaining gaps:**
- **Review date:**

Do not mark mastery based only on reading, generated code, or a successful compilation.
