# Cosmos Learning Roadmap — TeamTask Chain

Status: **Fixed sequence for the TeamTask Chain learning project.** This is a learning POC, not the Institutional Settlement Network. Do not use Codex for this repository.

## Product use case

A team leader creates tasks, assigns a member, defines a deadline and reward, and funds the reward pool. The assigned member moves the task through its permitted workflow and submits it for review. The leader approves or rejects the submission. Approval releases the agreed reward exactly once; rejection records feedback and returns the task to in-progress. Chain-side authorization and state-transition validation are authoritative; frontend controls are not security controls.

The exact task states, messages, reward funding model, and validation rules must be designed and documented before implementation.

## Working cycle and gates

1. Explain the concept and check understanding.
2. Design the smallest required change and record the decisions.
3. Verify versions, APIs, source paths, and commands against the actual checkout.
4. Implement together without Codex.
5. Run tests/labs and record actual evidence, including failures.
6. Trace the implementation through source code and update diagrams/docs.
7. Complete a mastery check before moving on.

Do not skip gates. Do not invent outputs or silently change the stack or use case.

## Phase 0 — Repository and environment baseline

- Inspect the current cosmos-learning-lab checkout and its existing docs.
- Verify OS/WSL environment, Go, Git, protobuf tooling, Node/frontend prerequisites, and installed Cosmos example.
- Record exact versions and distinguish installed facts from planned requirements.
- Use the selected target pair: Cosmos SDK `v0.55.0` and CometBFT `v0.40.0`, based on the dependency pair inspected in the existing official example.
- Preserve the official cosmos/example run as a reference only; implement TeamTask Chain in this repository rather than copying the example.
- Validate dependency resolution, build, tests, and local-chain startup during the real implementation; do not create a separate disposable compatibility experiment.

**Exit gate:** begin the app skeleton with the selected versions, and record actual build/test evidence as it becomes available. Do not claim validation before it runs.

## Phase 1 — Cosmos application foundations

- Understand Cosmos SDK, CometBFT, ABCI/application boundary, deterministic state machines, and module architecture.
- Trace genesis, configuration, app wiring, node startup, block production, and local queries in the selected version.
- Run and inspect a minimal local chain using verified commands.
- Document key directories and source call paths.

**Exit gate:** start the local chain from a clean documented state and explain how the app is wired and started.

## Phase 2 — Transaction lifecycle

- Messages and transactions; protobuf schemas and generated code.
- Account addresses, public/private keys, signatures, encoding, and signers.
- Mempool/broadcast lifecycle, ante handling, fees, gas, execution, events, transaction queries, and errors.
- Trace one successful transaction and deliberately test invalid inputs.

**Exit gate:** explain and verify the path from CLI submission to committed state and queryable result.

## Phase 3 — TeamTask domain and custom module design

- Define task fields, roles, task states, allowed transitions, errors, and events.
- Design the smallest custom module and keeper/store boundaries.
- Define authorization rules and invariants before coding.
- Create state-transition diagrams and message/query specifications.

**Exit gate:** reviewed domain model, transition matrix, authorization matrix, and test cases before implementation.

## Phase 4 — Implement the TeamTask module

- Add task creation and assignment, deadline/reward fields, status transitions, submission, review, and feedback.
- Learn keepers, stores, codecs, protobuf messages, message servers, queries, events, and errors through the actual implementation.
- Test unauthorized actions, invalid transitions, duplicate operations, and deterministic behavior.

**Exit gate:** core task lifecycle works in repeatable tests and the learner can trace each message through validation, keeper, store, and emitted events.

## Phase 5 — Roles, reward pool, and bank integration

- Define leader/member permissions using transaction signer identity.
- Fund a reward pool using explicit, verifiable rules.
- Integrate with the bank module for reward transfer.
- Ensure only approved tasks pay rewards and that a task cannot be rewarded twice.
- Test insufficient funds, wrong signer, rejection, duplicate approval, and transfer failures.

**Exit gate:** authorization and reward invariants are covered by tests; reward issuance is funded and exactly-once at the application state-machine level.

## Phase 6 — Queries and developer interfaces

- Add CLI commands and queries for tasks, assignments, status, review history, and rewards.
- Understand gRPC/query services, REST gateway availability in the chosen version, transaction search, and event inspection.
- Record exact commands and actual responses.

**Exit gate:** users can inspect all required POC state through documented query/CLI paths without directly editing chain state.

## Phase 7 — React frontend and wallet flow

- Build a small UI for leader and member workflows.
- Connect a compatible wallet and understand address selection, signing, broadcasting, pending state, confirmation, and errors.
- Use chain-side authorization as the security boundary.
- Keep frontend scope limited to the working TeamTask lifecycle.

**Exit gate:** demonstrate create → assign → progress → submit → approve/reject → reward, with chain state independently verified.

## Phase 8 — Quality and failure experiments

- Unit, integration, and end-to-end tests.
- Authorization bypass attempts, invalid transitions, duplicate reward attempts, insufficient funds, invalid deadlines, and malformed inputs.
- State invariants, events, error semantics, gas observations, and deterministic execution.
- Document failures, diagnosis, fixes, and regression tests.

**Exit gate:** repeatable test suite and evidence-backed explanation of the main failure paths.

## Phase 9 — Multi-validator network and CometBFT

- Run a small local multi-validator network using version-compatible tooling.
- Study proposal/voting flow, validator power, block commit, finality, node configuration, and node failure behavior.
- Distinguish application execution from consensus responsibilities.

**Exit gate:** capture and explain multi-validator observations and controlled failure/recovery experiments.

## Phase 10 — Governance, parameters, and upgrades

- Study governance proposal lifecycle and parameter changes where supported by the chosen application/version.
- Practice a controlled upgrade and state migration in a disposable local network.
- Understand app versioning, migration handlers, genesis export/import, and rollback limitations.

**Exit gate:** complete a reproducible local upgrade or migration experiment with before/after state checks.

## Phase 11 — CosmWasm

- Learn Rust/Wasm contract architecture, entry points, storage, instantiate/execute/query, events, gas, permissions, and contract migration.
- Add a small contract experiment related to TeamTask only after the SDK application is stable.
- Compare native SDK module logic with contract-based logic and document trade-offs.

**Exit gate:** deploy and test a local contract and explain its lifecycle, execution boundaries, and migration behavior.

## Phase 12 — IBC fundamentals with two chains

- Build/run two compatible local chains and verify exact IBC-Go compatibility before use.
- Study light clients, client updates, proofs, connections, channels, packet commitments, acknowledgements, timeouts, relayers, and recovery.
- Run successful and failed packet experiments.

**Exit gate:** trace a packet end to end and demonstrate acknowledgement/timeout behavior with recorded evidence.

## Phase 13 — Interchain token and application experiments

- Study ICS-20 transfer, escrow, voucher denominations, packet refunds, and reconciliation.
- Explore interchain accounts and IBC-enabled CosmWasm only after the base packet flow is understood and version compatibility is verified.
- Keep interchain experiments in the same repository and document them as optional extensions to TeamTask, not a new product.

**Exit gate:** demonstrate a cross-chain transfer and at least one failure/recovery path across two local chains.

## Phase 14 — Security and operations

- Threat model, key handling, least privilege, supply-chain/dependency review, and safe local development.
- Logging, events, metrics, monitoring, backups, recovery, upgrade runbooks, and performance methodology.
- Fuzz/property testing where appropriate; no production-readiness claim without independent review.

**Exit gate:** documented threat model, operational runbooks, and reproducible test/performance evidence.

## Phase 15 — Interview and mastery review

- Explain architecture and trade-offs without notes.
- Draw transaction, state-transition, consensus, contract, and IBC packet flows.
- Debug selected failures and justify security and reliability decisions.
- Track demonstrated skills and remaining gaps.

**Exit gate:** each topic has an explanation, verified source references, lab/test evidence, failure understanding, and interview questions with answers.

## Scope and honesty

This sequence covers major Cosmos SDK, CometBFT, CosmWasm, and IBC concepts through one integrated learning project plus focused experiments. It does not claim to cover every feature in the Cosmos ecosystem. Any additional topic must be added explicitly to the roadmap with a reason and an acceptance gate.
