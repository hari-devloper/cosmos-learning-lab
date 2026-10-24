# Module 5 — Cosmos Stack Architecture: Application, Consensus, and ABCI

- **Module ID:** COSMOS-ARCH-001
- **Status:** Conceptual learning draft; not a version-pinned implementation guide
- **Prerequisites:** COSMOS-FOUND-001 through COSMOS-FOUND-004
- **Project sequence:** Single-chain first; IBC is deferred

## Learning outcomes

By the end of this module, the learner should be able to:
1. Explain the distinct responsibilities of a Cosmos SDK application and CometBFT.
2. Describe the purpose of the ABCI boundary without assuming a version-specific call order.
3. Distinguish consensus ordering from application state-transition rules.
4. Explain why deterministic application behavior matters to replicated state.
5. Identify which implementation details must be verified against the selected release.

## 1. The three conceptual layers

### CometBFT

At a high level, CometBFT supplies blockchain-node networking and Byzantine-fault-tolerant consensus machinery. It helps participating validators agree on a sequence of blocks under the protocol's assumptions. Consensus is not a substitute for the application's business rules.

### Cosmos SDK application

The application defines the chain's state and business/protocol behavior. In a Cosmos SDK application, modules and keepers commonly organize state access, message handling, validation, queries, and other application functionality. The exact architecture and APIs depend on the selected SDK release.

For the institutional settlement project, this is where the authoritative transfer rules and state invariants must ultimately be enforced—not in a React UI or a PostgreSQL indexer.

### ABCI boundary

ABCI is the application interface between the consensus/node layer and the application. It lets the node interact with the application through protocol-defined requests and responses. The exact methods, execution phases, and semantics have evolved; do not assume one timeless call sequence across all CometBFT and Cosmos SDK releases.

## 2. Responsibility map

| Concern | Primary conceptual owner | Important limitation |
|---|---|---|
| Validator communication and consensus | CometBFT | Does not define the institutional payment rules |
| Application state and transfer invariants | Cosmos SDK application | Must implement deterministic, correct rules |
| Node/application interaction | ABCI boundary | Exact API and lifecycle are version-specific |
| User interface and API ergonomics | Client/API layer | Cannot be the sole enforcement point for settlement invariants |
| Search, dashboards, and reporting projection | Indexer/read model such as PostgreSQL | Derived data must not override authoritative chain state |

This is a conceptual responsibility map, not a complete deployment diagram.

## 3. Consensus ordering is not business correctness

Imagine validators agree on a block containing a transfer request. Agreement on block history does not mean the transfer is economically valid. The application still needs to enforce authorization, balance safety, value conservation, atomicity, replay protection, and other specified rules.

Conversely, correct application code is not enough if nodes disagree about transaction ordering or committed history. The protocol and application solve different parts of the system's correctness problem.

A useful mental model:

- **Consensus:** which ordered history is committed?
- **Application state machine:** what deterministic state transition follows from the protocol inputs and current state?
- **Business rules:** which transitions are allowed?
- **Read model:** how is committed history projected for convenient queries?

## 4. Why deterministic execution matters

Validators that apply the same committed inputs to equivalent prior state must follow the protocol's deterministic execution requirements. Application behavior must not depend on one validator's local clock, external HTTP response, local file, unseeded randomness, or environment-specific value unless the protocol explicitly handles that input in a deterministic way.

Nondeterministic execution can cause validators to compute different state from the same history, undermining the chain's safety assumptions.

## 5. A version-neutral conceptual flow

1. A client constructs and signs a transaction using the selected transaction format.
2. A node receives the transaction; protocol and application checks occur in the lifecycle defined by the chosen versions.
3. Consensus machinery establishes the committed block history according to protocol rules.
4. The application processes inputs and state transitions according to the selected application's protocol-defined lifecycle.
5. Clients query application state or consume events.
6. An external indexer may project committed data into PostgreSQL and reconcile/rebuild that projection.

**Important:** This list explains responsibilities, not a guaranteed chronological call trace. The actual ordering of checking, proposal processing, execution, commit, and event emission must be traced from exact pinned versions before implementation.

## 6. Questions to answer before selecting a baseline

- Which Cosmos SDK release family will be used for the single-chain learning application?
- Which CometBFT version is required by that SDK release?
- What Go toolchain version is required by the selected module graph?
- Which ABCI methods and lifecycle semantics apply to that exact combination?
- How are transactions encoded, signed, checked, executed, and committed in that release?
- What state store, cache, event, and query abstractions are provided?
- What official sample application or minimal app is maintained for that release?
- Can the selected dependency graph build and pass tests in a clean Linux/WSL environment?

Do not lock versions solely because a version number is newer or because it might support a later feature. IBC is a later project stage, so initial selection should first satisfy the approved single-chain learning scope and be recorded as a deliberate, verified decision.

## 7. Failure thought experiments

### Experiment A — Consensus without business validity
Validators agree on the block history, but a transfer would overdraw an account if applied. Explain why consensus does not remove the need for application-level invariant enforcement.

### Experiment B — Nondeterministic dependency
Application execution queries an external price API during state transition, and validators receive different responses. Explain the risk and propose a deterministic design alternative.

### Experiment C — Layer confusion
A frontend prevents a user from entering a transfer amount above the displayed balance, but a client sends a crafted request directly to a node. Explain which layer must enforce the invariant.

### Experiment D — Version mismatch
A tutorial describes a transaction lifecycle for a different SDK/CometBFT release family. Explain why the names and call order should be verified before relying on it.

## 8. Interview questions

1. What does CometBFT do, and what does it not do?
2. What does a Cosmos SDK application define?
3. What is the role of ABCI?
4. Why is consensus not equivalent to business-rule validation?
5. Why must state-transition execution be deterministic?
6. Why can frontend validation never be the only settlement safeguard?
7. What is the difference between chain state and an indexer read model?
8. Why must source paths and lifecycle claims be pinned to a specific version?
9. What is the risk of making external API calls during deterministic state transition?
10. What evidence is required before approving a dependency baseline?

## 9. Mastery exercise

- [ ] Draw the relationship between a client, CometBFT, the ABCI boundary, the Cosmos SDK application, and PostgreSQL read model.
- [ ] Label consensus responsibilities versus application responsibilities.
- [ ] Explain why an agreed block may still contain a transaction that the application rejects or treats as invalid under the selected protocol.
- [ ] Explain one nondeterministic design and how to replace it with deterministic inputs.
- [ ] Identify at least five version-specific details that must be checked in official source or module files.
- [ ] Answer all ten interview questions in your own words.

## 10. Evidence and implementation status

- **Conceptual explanations:** included.
- **Runnable application lab:** not implemented.
- **Pinned SDK/CometBFT versions:** not selected by this module.
- **Source-code call paths:** deferred until a specific compatible release pair is approved and verified.
- **Tests/performance results:** none claimed.
- **Next gate:** complete the mastery check and verify the development terminal/toolchain context, then perform a source-backed compatibility review before writing application code.
