# TeamTask Chain — Learning Roadmap

**Goal:** build one small, working Cosmos SDK application and use it to learn the major Cosmos concepts. Keep the product simple; cover advanced subjects through focused labs in this repository. Do not use Codex for this learning project.

## Product

A leader creates and assigns a task with a deadline and funded reward. The assigned member starts and submits it. The leader approves or rejects it with feedback. Approval transfers the funded reward once. Chain-side signer checks and state-transition rules are authoritative.

## Rules

- One milestone at a time; do not skip its exit gate.
- Explain the concept and inspect the relevant source before implementing.
- Implement and test together; no fabricated outputs or assumed source paths.
- Update docs, diagrams, commands, tests, failures, and interview notes as work happens.
- Keep the application small. Advanced topics are labs, not extra product features.
- See `docs/FINAL_COSMOS_POC_SCOPE.md` for the full list of major topics and the distinction between core implementation and labs.

## Stage 1 — Small task-domain model

- Run `go test ./...` inside `poc/domain/`.
- Understand task fields, role checks, state transitions, rejection feedback, deadline validation, and one-time logical reward handling.
- Add tests for invalid actors, invalid transitions, and malformed tasks.

**Exit gate:** tests pass locally and the learner can explain each invariant. This model is not the blockchain and does not transfer real coins.

## Stage 2 — Real Cosmos chain

- Scaffold the chain under `chain/` using `docs/POC_BUILD_GUIDE.md`.
- Inspect generated app wiring, genesis/configuration, SDK/CometBFT/IBC versions and source paths.
- Build, test, initialize and start a local node; record observed output.
- Learn SDK architecture, CometBFT, ABCI, deterministic execution, app wiring, module manager and startup.

**Exit gate:** a clean documented local-chain start, passing build/tests, and a source-traced explanation of how the app starts.

## Stage 3 — Transaction path and custom task module

- Trace protobuf, messages, signers, keys/signatures, encoding, mempool/broadcast, ante handling, fees, gas, execution, events, errors and transaction queries.
- Implement only five actions: create, start, submit, approve, reject.
- Use keepers, context/cache, codecs, KV store/MultiStore, message servers, queries and module registration.
- Enforce leader/member authorization and valid state transitions on-chain.

**Exit gate:** successful and invalid transactions are tested; committed state and events are queried; the learner can trace message → validation → keeper → store → result.

## Stage 4 — Funded rewards, interfaces and quality

- Use the standard auth and bank modules; no custom token for the first POC.
- Require an explicitly funded reward pool. Approval and bank transfer must be atomic and a task cannot be paid twice.
- Test insufficient funds, wrong signer, invalid transitions, duplicate approval, malformed inputs and rejection without payout.
- Use CLI/query APIs first; learn gRPC and REST availability in the chosen version.
- Add a minimal wallet/frontend only after the CLI workflow is stable.
- Cover unit/integration/end-to-end tests, invariant/property/fuzz testing where appropriate, gas observations and deterministic behavior.

**Exit gate:** repeatable tests prove the full create → assign → start → submit → approve/reject flow and independently verify reward balances.

## Stage 5 — Consensus and chain operations lab

- Run a small multi-validator local network.
- Study validator keys/power, proposer selection, rounds, prevote/precommit, commit/finality, peer networking and the application/consensus boundary.
- Explore node crash/downtime, state sync, snapshots, pruning, telemetry and recovery.
- Record exact observations; do not imply a local lab proves production readiness.

**Exit gate:** explain and demonstrate one normal consensus round and one controlled failure/recovery scenario.

## Stage 6 — Governance and upgrades lab

- Learn governance proposals and parameters supported by the chosen version.
- Practice software upgrade handlers, state migration, genesis export/import, state compatibility and rollback limitations in a disposable local network.

**Exit gate:** complete a repeatable upgrade/migration experiment with before/after state checks.

## Stage 7 — CosmWasm lab

- Verify compatible wasmd/CosmWasm versions before adding dependencies.
- Learn Rust/Wasm basics, contract storage, instantiate/execute/query, host boundary, gas, permissions and contract migration.
- Build one tiny related contract and compare it with native SDK module logic.

**Exit gate:** deploy, execute and query a local contract; document one failure and migration/security implications.

## Stage 8 — IBC and interchain labs

- Verify the generated SDK, CometBFT and IBC-Go dependency family before starting.
- Run two compatible local chains; learn clients/light clients, proofs, client updates, connection/channel handshakes, packet commitments, relayers, acknowledgements, timeouts and recovery.
- Demonstrate ICS-20 escrow/vouchers/refunds; then study ICS-27 Interchain Accounts, middleware and IBC-enabled CosmWasm where supported.

**Exit gate:** trace a packet end to end and demonstrate a successful transfer plus a failure/refund or timeout path.

## Stage 9 — Security, operations and mastery

- Threat model, key handling, least privilege, dependency/supply-chain review, monitoring, logs, metrics, backups, recovery, upgrade runbooks, profiling and reproducible performance tests.
- For every major topic, document: concept and mental model; architecture/sequence diagrams; verified source paths; design trade-offs; exact commands; expected vs observed output; tests/failures/fixes; security/performance implications; interview questions and mastery evidence.
- Explain the transaction, task, consensus, contract and IBC flows without notes and diagnose selected failures.

**Exit gate:** every row in `docs/FINAL_COSMOS_POC_SCOPE.md` has either verified lab evidence or an explicitly documented limitation. Do not claim mastery from reading alone.

## Scope boundary

This curriculum covers major Cosmos SDK, CometBFT, CosmWasm, IBC/interchain and operational concepts; it does not claim to cover every feature in the ecosystem. Add no topic silently: update the coverage matrix, explain why it matters, and define an acceptance gate first.
