# Final Cosmos POC Scope and Completion Plan

**Project:** TeamTask Chain — Cosmos SDK learning POC  
**Repository:** `hari-devloper/cosmos-learning-lab`  
**Scope status:** Baseline plan; implementation evidence must be recorded as each stage is completed.

## 1. Goal and scope contract

Build one small, understandable TeamTask blockchain application and use it as the anchor for learning Cosmos SDK. Cover all 14 learning domains below through the core application and small, isolated labs. **A topic is not complete just because it appears in this document or has been read.** It is complete only when its hands-on exercise, test or experiment, explanation, and evidence are recorded.

Keep the TeamTask business model small: create and assign a task, move it through a controlled lifecycle, review it, and release a reward only when approval rules pass. Advanced areas such as multi-validator networking, IBC, and CosmWasm are separate labs; they do not become unnecessary dependencies of the core task module.

## 2. The 14 required learning domains

### 1. Application foundations
- Cosmos SDK application architecture and module composition
- CometBFT and the ABCI/application boundary
- Deterministic state-machine execution
- Application wiring, genesis, configuration, node startup and shutdown
- **Practical evidence:** start the chain locally and trace one transaction from node entry point to application execution.

### 2. Transaction lifecycle
- Protobuf messages and generated types
- Message signers, keys, signatures and address encoding
- Transaction encoding, signing, broadcast and mempool
- Ante handling, fee deduction, gas accounting and execution
- Events, errors, transaction hash and transaction query
- **Practical evidence:** submit one valid and multiple invalid transactions; document their path, events, gas and errors.

### 3. State and modules
- BaseApp and execution context/cache concepts
- Keepers, stores, codecs and state access boundaries
- KVStore/MultiStore and module manager/order
- Block lifecycle hooks or their version-specific equivalents
- State commitment, proofs and query behavior
- **Practical evidence:** trace a TeamTask state change from message handler through keeper/store and query the resulting state.

### 4. Standard Cosmos modules
- Auth and account model
- Bank balances, sends and module accounts
- Staking, validator delegation and unbonding concepts
- Slashing and distribution/rewards
- Governance, upgrade and evidence modules
- Authz and feegrant where supported by the selected SDK version
- **Practical evidence:** inspect module configuration and run focused commands/experiments for the modules included in the chosen dependency version. Record unsupported/version-dependent features rather than assuming they exist.

### 5. TeamTask business logic
- Create and assign tasks; roles and signer-based authorization
- Valid task state transitions and invalid-transition rejection
- Deadline and reward validation
- Review feedback and rejection/rework path
- Funded reward payment, duplicate-action prevention and exactly-once business semantics
- **Practical evidence:** domain tests plus SDK integration tests for success, wrong signer, invalid state, repeated approval, insufficient funds and reward accounting.

### 6. Client interfaces
- Chain CLI and key management for a local learning environment
- gRPC query and transaction services
- REST gateway when supported/enabled
- Events, transaction confirmation and error interpretation
- Wallet signing/broadcast concepts; a UI only after chain behavior is stable
- **Practical evidence:** perform the complete task lifecycle using CLI, query the state and explain the transaction result.

### 7. Testing and correctness
- Unit, module integration and end-to-end tests
- Invalid inputs, unauthorized callers and duplicate actions
- Insufficient funds and failure-path tests
- Deterministic execution and invariant testing
- Fuzz/property testing for appropriate message and state boundaries
- **Practical evidence:** reproducible test commands, passing output, and a short failure-test report.

### 8. Consensus and networking
- Validator keys, voting power and validator-set changes
- Proposer selection and prevote/precommit/round concepts
- Commit, finality and the relationship between consensus and application state
- Peer networking, node configuration and node failure
- State sync, snapshots and pruning concepts
- **Practical evidence:** run a two- or three-validator local network, submit a transaction, inspect consensus/validator behavior, and document a controlled node failure/recovery experiment.

### 9. Governance and software upgrades
- Governance proposals and voting
- Parameter changes and authorization
- Software-upgrade plan and migration handlers
- Genesis export/import and upgrade rehearsal
- Rollback limitations and recovery planning
- **Practical evidence:** run a disposable local governance or upgrade exercise and document preconditions, result and recovery limitations.

### 10. CosmWasm
- Rust/Wasm contract basics and build artifacts
- Instantiate, execute and query
- Contract storage, permissions and host/chain boundary
- Gas and failure behavior
- Contract migration and version compatibility
- **Practical evidence:** build and deploy one minimal contract on a compatible local chain, call it, query it, test a failure, and document gas/permissions. Keep it separate from TeamTask unless a concrete requirement emerges.

### 11. IBC core
- IBC-Go architecture and client/light-client responsibilities
- Client updates and proof verification concepts
- Connection and channel handshakes
- Packet commitments, acknowledgements, timeouts and relayers
- Trust assumptions, packet failure and recovery
- **Practical evidence:** connect two compatible local chains and trace a packet through send, relay, acknowledgement or timeout. Verify dependency compatibility before implementation.

### 12. Interchain applications
- ICS-20 fungible-token transfer
- Escrow, voucher denomination/traces, balances and refunds
- Timeout/error handling and packet recovery
- ICS-27 Interchain Accounts and controller/host concepts
- Middleware and IBC-enabled contracts where supported
- **Practical evidence:** complete a local ICS-20 transfer, inspect voucher/escrow state, test a failure/refund path, and run a small ICS-27 demonstration if compatible.

### 13. Security and operations
- Threat model, signer/role boundaries and least privilege
- Key custody and safe local-development key handling
- Dependency and software supply-chain review
- Logs, metrics, telemetry and basic troubleshooting
- Backup/restore, state recovery and operational runbooks
- Load/performance tests and profiling
- **Practical evidence:** a threat model, dependency review, observable local run, recovery drill and reproducible performance baseline. Never put real secrets in the repository.

### 14. Interview and architecture mastery
- Explain end-to-end transaction and block lifecycle without notes
- Trace a message through source code into keeper/store changes
- Explain architectural trade-offs and version-specific decisions
- Diagnose deliberately introduced failures using logs/tests/state
- Draw architecture and sequence diagrams and justify security choices
- **Practical evidence:** complete a written and oral-style checkpoint with diagrams, source references, failure diagnosis and recorded gaps.

## 3. Delivery order

Keep the order stable. Do not start every advanced topic at once.

1. **Baseline and toolchain:** record Go, Ignite CLI, Cosmos SDK, CometBFT, IBC-Go and CosmWasm compatibility from the actual generated dependency graph.
2. **Core chain:** scaffold, build, test and start a local single-node chain.
3. **TeamTask module:** implement messages, signer checks, keeper/state, queries and funded bank-module rewards.
4. **Core verification:** transaction lifecycle, CLI/gRPC, negative tests, events, gas and source tracing.
5. **Multi-validator lab:** expand to two or three local validators and test consensus/node behavior.
6. **Governance and upgrades:** run disposable local experiments.
7. **IBC labs:** first IBC core and two-chain setup, then ICS-20, failure/refund tests and ICS-27.
8. **CosmWasm lab:** build and test a minimal contract on a compatible local environment.
9. **Operations and performance:** threat model, telemetry, recovery and baseline performance.
10. **Mastery review:** complete all evidence and interview checkpoints.

The order of the IBC and CosmWasm labs may be adjusted only if official compatibility requirements demand it; record the reason in the toolchain notes before changing sequence.

## 4. Completion tracking

Maintain one row per domain in the project tracker using these statuses:

- **Planned:** scope identified; no hands-on evidence yet.
- **In progress:** implementation or lab started; evidence incomplete.
- **Verified:** hands-on exercise passes, relevant tests/commands and results are recorded, and the concept explanation/source trace is documented.
- **Blocked:** a specific dependency or compatibility issue is recorded with a next action.

For every verified topic, add:
1. Concept explanation suitable for learning and interviews.
2. Architecture or sequence diagram when it adds clarity.
3. Commands and environment/version details.
4. Source-code path/call trace or relevant configuration.
5. Tests/experiment steps and actual results.
6. Failure case and troubleshooting notes.
7. Decision/trade-off and known limitations.

## 5. Definition of done

The overall POC is complete only when:

- All 14 domains have **Verified** evidence or an explicitly documented compatibility blocker with an agreed resolution; a planned item is not counted as covered.
- The core TeamTask chain builds and starts reproducibly from a clean checkout.
- Task creation, assignment, status changes, review/rejection and reward release work end-to-end.
- Authorization, invalid transitions, duplicate actions and insufficient-funds cases are tested.
- Multi-validator, governance/upgrade, IBC/ICS-20/ICS-27, and CosmWasm exercises have reproducible results in compatible local environments.
- Documentation is updated alongside implementation, not after the project is finished.
- The mastery review can explain architecture, source flow, security, failure modes and trade-offs.

## 6. Scope protection

- Do not silently replace TeamTask with a different product idea or change the 14-domain goal.
- Do not add infrastructure or frameworks unless a learning objective requires them.
- Prefer the smallest working experiment for each advanced domain.
- Verify version compatibility from generated dependencies and official release documentation before writing integration code.
- Distinguish **planned**, **implemented**, and **verified** in every progress report. Never report a topic as completed on the strength of this plan alone.

## 7. Progress snapshot at the time this file was written

- **Scope defined:** all 14 domains.
- **Basic Go TeamTask domain model/tests:** prepared; previously passed in the available Go environment.
- **Cosmos SDK chain scaffold and end-to-end chain execution:** not yet verified.
- **Advanced labs (multi-validator, governance/upgrades, CosmWasm, IBC/ICS-20/ICS-27):** planned, not yet verified.

This snapshot must be updated only when there is concrete implementation or test evidence.
