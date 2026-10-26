# Cosmos Concept Coverage Matrix — TeamTask Chain

This matrix is the scope contract. It separates the small application POC from focused labs so broad learning coverage does not make the product itself complicated.

**Status meanings:** Planned = mapped but not implemented; In progress = active with evidence; Verified = code/lab and evidence exist. Reading or writing a plan alone never counts as verified.

| Domain | Concepts to cover | How we learn it | Current status |
|---|---|---|---|
| App foundations | SDK architecture, CometBFT, ABCI boundary, deterministic state machine, app wiring, genesis, config, node startup | Run the generated chain; trace real source | Planned |
| Transaction path | Protobuf, messages, signers, keys/signatures, encoding, mempool, ante handling, fees, gas, execution, events, errors, transaction query | Submit successful and invalid transactions; trace the path | Planned |
| State and modules | BaseApp, context/cache, keepers, KV store/MultiStore, codecs, module manager/order, begin/end block or current equivalents, events, errors, state commitment/proofs | Implement task module; inspect source and state | Planned |
| Standard modules | Auth, bank, staking, slashing, distribution, governance, upgrade, evidence, feegrant/authz where supported | Use only the modules needed by a lab; inspect boundaries and permissions | Planned |
| TeamTask business rules | Task creation/assignment, roles, state transitions, deadline/reward validation, feedback, exactly-once reward | Domain tests then custom SDK module tests | In progress only when tests are run; code currently proposed |
| Interfaces | CLI, gRPC/query services, REST gateway where supported, events, wallet signing/broadcast/confirmation | CLI first; small UI after chain behavior is stable | Planned |
| Testing and correctness | Unit/integration/e2e, invalid inputs, authorization bypass, duplicate actions, insufficient funds, deterministic execution, fuzz/property/invariant testing | Focused test and failure labs | Planned |
| Consensus and networking | Validator keys/power, proposer selection, prevote/precommit/rounds, commit/finality, peer networking, node failure, state sync/snapshots/pruning | Two- or three-validator local lab | Planned |
| Governance and upgrades | Proposals, parameters, software upgrades, migration handlers, genesis export/import, rollback limits | Disposable local upgrade lab | Planned |
| CosmWasm | Rust/Wasm basics, instantiate/execute/query, storage, host boundary, permissions, gas, migration | One small related contract, only after core chain | Planned |
| IBC core | IBC-Go architecture, clients/light clients, proofs, connection/channel handshakes, packet commitments, relayers, acknowledgements, timeouts/recovery | Two compatible local chains | Planned |
| Interchain apps | ICS-20 escrow/vouchers/refunds, ICS-27 Interchain Accounts, middleware and IBC-enabled contracts where supported | Small transfer plus failure/refund experiment | Planned |
| Security and operations | Key management, threat model, dependency/supply-chain review, telemetry, logs, metrics, backup/recovery, upgrade runbook, performance/profile | Reproducible operations labs | Planned |
| Interview mastery | Architecture and trade-offs, source tracing, diagrams, failure diagnosis, design decisions | Oral/written checkpoint tied to evidence | Planned |

## Scope rule

- The **core POC** only implements task workflow, signer-based authorization, queries, and funded bank-module rewards.
- **Advanced Cosmos topics** are small labs in the same repository. They do not become dependencies of the task module unless a concrete requirement calls for them.
- IBC and CosmWasm release compatibility must be verified from the generated dependency graph and official release guidance before those labs begin.
- Do not claim to cover every Cosmos ecosystem feature. This matrix covers the major SDK, consensus, contract, interchain and operations concepts appropriate for a serious learning project.
