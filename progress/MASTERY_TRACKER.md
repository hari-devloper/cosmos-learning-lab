# Mastery Tracker — TeamTask Chain

This tracker records demonstrated understanding and reproducible evidence, not intentions. Do not mark an item complete just because documentation was read or code compiled.

## Status key

- [ ] Not started
- [~] In progress; include evidence and remaining gaps
- [x] Complete; include evidence and a mastery demonstration

## Project reset status

The learning project is now TeamTask Chain. The Institutional Settlement Network is separate. No Codex is used in this repository. Existing foundational modules remain useful study material, but they do not automatically satisfy a TeamTask milestone or prove mastery.

## Phase 0 — Repository and environment baseline

- [x] Initial WSL/toolchain observations recorded — docs/ENVIRONMENT_AUDIT.md
- [x] Selected target pair recorded: Cosmos SDK `v0.55.0` + CometBFT `v0.40.0` — docs/TOOLCHAIN_COMPATIBILITY_DRAFT.md
- [ ] Validate dependency resolution during actual TeamTask Chain implementation
- [ ] Build the application and run available tests; record actual results
- [ ] Run a local-chain smoke test and record observed output
- [ ] Update the environment audit only if implementation exposes a concrete prerequisite gap

## Fixed milestone tracker

### 1. Cosmos application foundations
- [ ] Explain SDK, CometBFT, ABCI/application boundary, deterministic state
- [ ] Verify app wiring, genesis/configuration, and node startup source paths
- [ ] Start and inspect the local chain from documented steps
- [ ] Pass explanation and source-tracing checkpoint

### 2. Transaction lifecycle
- [ ] Explain messages, protobuf, signer, signature, encoding, fees and gas
- [ ] Trace submission through execution and committed state
- [ ] Verify transaction result, events, query, and error behavior
- [ ] Complete invalid-input/failure experiment

### 3. TeamTask domain design
- [ ] Approve roles and permissions matrix
- [ ] Approve task state-transition table
- [ ] Define task fields, reward-pool rules, errors and events
- [ ] Define invariants and tests before implementation

### 4. Custom TeamTask module
- [ ] Implement task creation and assignment
- [ ] Implement authorized member workflow and submission
- [ ] Implement leader review, rejection feedback and state changes
- [ ] Trace message → validation → keeper → store → event/query
- [ ] Pass authorization and invalid-transition tests

### 5. Rewards and bank integration
- [ ] Define and test funded reward pool
- [ ] Transfer reward only after approval
- [ ] Prevent duplicate reward
- [ ] Test insufficient funds, wrong signer and repeated approval

### 6. Query and developer interfaces
- [ ] Task and assignment queries
- [ ] CLI and transaction/event inspection
- [ ] Verify gRPC/REST capabilities for the selected release

### 7. React frontend and wallet
- [ ] Connect wallet and show selected address
- [ ] Sign/broadcast transactions and display pending/success/failure states
- [ ] Demonstrate full workflow and independently verify chain state

### 8. Quality and failure experiments
- [ ] Unit, integration and end-to-end tests
- [ ] Authorization, duplicate operation, invalid transition and insufficient-funds tests
- [ ] Record gas observations and test methodology without unsupported claims

### 9. Multi-validator consensus
- [ ] Run compatible local multi-validator network
- [ ] Explain validator power, voting, block commit and finality
- [ ] Record controlled failure/recovery experiment

### 10. Governance and upgrades
- [ ] Study proposal/parameter flow supported by selected version
- [ ] Perform disposable local upgrade or migration
- [ ] Verify state before and after migration

### 11. CosmWasm
- [ ] Explain Rust/Wasm contract lifecycle and execution boundary
- [ ] Deploy and test a local contract experiment
- [ ] Demonstrate query/execute and contract migration where supported

### 12. IBC fundamentals
- [ ] Run two compatible local chains
- [ ] Trace clients, proofs, connections, channels and packet flow
- [ ] Demonstrate acknowledgement and timeout/recovery paths

### 13. Interchain extensions
- [ ] Demonstrate ICS-20 transfer and voucher/escrow behavior
- [ ] Test a failure/refund path
- [ ] Explore ICA or IBC-enabled CosmWasm after fundamentals pass

### 14. Security and operations
- [ ] Threat model and key-handling notes
- [ ] Monitoring, backup, recovery and upgrade runbooks
- [ ] Reproducible performance methodology and results

### 15. Interview mastery
- [ ] Explain major concepts without notes
- [ ] Draw transaction, task lifecycle, consensus, contract and IBC flows
- [ ] Explain at least one failure and trade-off per major phase

## Existing foundational evidence (historical; not a project completion claim)

Previous study includes conceptual checkpoints and source investigations around deterministic state machines, consensus/finality, transaction lifecycle, hashes/keys/signatures, BaseApp, ABCI, AnteHandler and module wiring. Detailed notes remain in docs/modules/. Some prior source investigation used Cosmos SDK v0.54.4; that is not an approved dependency choice for TeamTask Chain.

## Evidence log

Append each completed experiment with date, exact version, command, actual output or evidence link, interpretation, failure case, and remaining gaps.

## Mastery evidence template

- **Concept/milestone:**
- **Version/commit:**
- **Learner explanation / diagram:**
- **Lab and exact commands:**
- **Observed result and evidence link:**
- **Failure experiment and diagnosis:**
- **Interview questions answered:**
- **Remaining gaps:**
- **Review date:**
