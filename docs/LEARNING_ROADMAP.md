# Cosmos Learning Roadmap

Status: **Draft curriculum — review before locking framework versions or implementation details.**

## Working method

1. GPT explains the concept and checks understanding.
2. We document the concept, diagram, source path, and lab plan.
3. We verify the relevant framework versions and source references.
4. We run a lab or failure experiment and record actual evidence.
5. The learner completes a review / mastery gate.
6. Only then may Codex implement the corresponding approved product change.

## Stage 0 — Environment and compatibility

- Confirm the Linux/WSL development context and toolchain.
- Verify exact compatible Cosmos SDK, CometBFT, IBC-Go, and optional CosmWasm versions from official release metadata and module files.
- Create a reproducible dependency baseline only after compatibility review.
- Keep the product architecture and compatibility choices explicitly provisional until approved.

Gate: a reproducible environment audit and evidence-backed dependency compatibility record.

## Stage 1 — Distributed systems foundations

- Deterministic state machines and replicated state
- Transactions, ordering, consensus, finality, and safety/liveness
- Accounts, keys, signatures, hashes, and encoding
- Byzantine fault tolerance concepts and operational assumptions

Gate: explain the trust model, failure assumptions, and why deterministic execution matters.

## Stage 2 — Cosmos stack architecture

- Cosmos SDK application architecture and module boundaries
- CometBFT responsibilities and ABCI/application boundary
- Transactions, messages, signers, ante processing, and execution lifecycle
- State stores, caches, events, queries, and errors
- Genesis, app configuration, upgrades, and local networks

Gate: trace one transaction from submission through validation, execution, state changes, and query/event observation using a version-pinned source path.

## Stage 3 — Build a minimal Cosmos application

- Scaffold only after versions and design are approved
- Define state, messages, validation, keeper boundaries, and queries
- Write unit, integration, and end-to-end tests
- Test invalid inputs, replay/idempotency assumptions, and deterministic behavior

Gate: explain and demonstrate the full path with tests and recorded evidence.

## Stage 4 — Institutional settlement domain

- Accounts, participants, assets, transfers, and authorization
- State and accounting invariants
- Settlement lifecycle, rejection, duplicate handling, and reconciliation
- Auditability, privacy boundaries, and operational controls

Gate: approve the domain model and invariant tests before product implementation.

## Stage 5 — IBC and cross-chain lifecycle

- IBC architecture, clients, connections, channels, and packets
- Timeout, acknowledgement, refund, recovery, and reconciliation paths
- Version-specific compatibility and local test networks

Gate: demonstrate success and failure paths; do not treat a packet timeout as proof that the remote chain never executed unrelated work.

## Stage 6 — Read models and interfaces

- Events and indexer/read-model design
- PostgreSQL as a derived, rebuildable projection—not the authoritative settlement state
- API contracts, React/TypeScript consoles, pagination, and observability
- Rebuild, replay, and reconciliation tests

Gate: demonstrate that the read model can be rebuilt and reconciled against authoritative chain state.

## Stage 7 — Security, reliability, and performance

- Threat modeling, key management, access control, and secure defaults
- Fuzz/property testing where appropriate
- Load tests with reproducible methodology
- Metrics, logs, alerts, backups, upgrades, and incident runbooks
- Independent security and operational review before any production claim

Gate: evidence-backed risk register and repeatable test suite.

## Stage 8 — Interview and mastery

- Explain architecture and trade-offs without notes
- Draw key execution and packet flows
- Debug deliberately injected failures
- Answer design, security, performance, and operations questions
- Track confidence, evidence, and gaps

## Mastery rule

A topic is not complete merely because code compiles. Completion requires explanation, reproducible evidence, failure understanding, and interview readiness. Actual results must be recorded only after the corresponding work has been performed.
