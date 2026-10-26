# Module 3 — Transactions, Application State, and Invariants

- **Module ID:** COSMOS-FOUND-003
- **Status:** Conceptual learning draft; not a version-pinned implementation guide
- **Prerequisites:** COSMOS-FOUND-001 and COSMOS-FOUND-002
- **Project sequence:** TeamTask single-chain application first; multi-validator, CosmWasm/wasmd, and IBC follow their explicit roadmap phases

## Learning outcomes

By the end of this module, the learner should be able to:
1. Separate transaction submission, validation, state transition, commitment, and read-model projection.
2. Define state invariants before implementation.
3. Explain atomicity and why UI-only validation is insufficient.
4. Distinguish authoritative chain state from frontend, cache, and event-consumer views.
5. Identify missing business rules that must be resolved before building a task-and-reward workflow module.

## 1. Conceptual transaction lifecycle

A useful conceptual lifecycle is:

1. **Construct:** the client prepares a transaction.
2. **Submit:** the client sends it to a node or API.
3. **Validate:** protocol and application checks determine whether it is admissible.
4. **Order/commit:** the consensus protocol establishes committed history according to its rules.
5. **Execute/apply:** deterministic application logic transitions state in the protocol-defined execution model.
6. **Project:** an indexer or event consumer updates a derived read model.
7. **Reconcile:** the system detects and repairs divergence or missed projection updates.

The precise order and naming of these stages vary by implementation. In particular, some systems perform portions of validation or execution before block commitment and may later discard uncommitted work. This diagram is a conceptual lifecycle, not a claim about a selected Cosmos SDK version.

## 2. Authoritative state versus read model

For this project:
- **Authoritative state:** the state maintained by the chain's application and committed history.
- **Client views:** queries, events, caches, and frontend data used to display chain state.
- **Indexer:** the component that consumes committed chain data and projects it into the read model.
- **Reconciliation:** a process that detects missing, duplicated, delayed, or inconsistent projections and repairs them.

A client view or cached response can be stale or absent. It must not independently authorize a task transition or reward; the chain remains authoritative.

## 3. Define invariants before writing code

An invariant is a property that must hold for every valid state (or at every specified committed boundary).

Candidate invariants for a simple account-transfer model:

### Balance safety
A transfer cannot debit more than the sender's spendable balance unless explicit overdraft/credit rules permit it.

### Value conservation
For a transfer without fees or issuance/destruction, the sender's debit equals the recipient's credit. Across accounts, total supply remains constant except for explicitly modeled minting, burning, fees, or other supply changes.

### Atomicity
A failed transfer must not leave only the debit or only the credit applied.

### Replay protection
A transaction identified as already processed must not be applied again. The mechanism and scope must be designed explicitly.

### Authorization
Only an identity meeting the defined authorization policy can cause the requested state transition.

### Read-model provenance
Client views represent chain-derived information and cannot override authoritative application state.

These are candidate product requirements, not a complete financial, legal, or production security specification.

## 4. Why frontend checks are not enough

A UI can show a balance of 100 and allow a user to prepare a transfer of 80. Another transfer of 80 may be submitted simultaneously from a different browser, API client, or institution. Both clients can have stale snapshots.

The authoritative transition logic must validate each transaction against the state used for execution. If the first transfer reduces the balance to 20, the second 80-unit transfer must be rejected unless the domain explicitly supports credit/overdraft behavior.

Validation in the API and UI is useful for user experience and early error reporting, but it is not a substitute for authoritative invariant enforcement.

## 5. Task-and-reward workflow domain questions that must be answered

Before implementation, the project must decide and document:
- What exactly is an account: institution, legal entity, sub-ledger, or wallet?
- What uniquely identifies an account, and who can control it?
- What task fields, reward denomination, and precision rules must be defined?
- Are balances integer smallest-units, and how are fees/rounding represented?
- What does “available balance” mean versus booked balance or reserved balance?
- Can an account be frozen, restricted, or subject to limits?
- How are authorization and institutional roles defined?
- How are duplicate submissions/retries recognized?
- Are transaction IDs globally unique, account-scoped, or derived from signed payloads?
- Which failures reject a transaction, and which create a pending workflow?
- What does committed/finalized mean in the UI?
- Which events are emitted for indexing, and how can the indexer detect gaps?
- How does the read model resume/replay after losing events?
- What is the recovery and reconciliation policy if an indexer has processed duplicate or out-of-order data?

Do not silently assume answers. Record open questions and resolve them with explicit design decisions.

## 6. Read-model reliability patterns

A robust read-model design should consider:
- storing chain identity and block/height information alongside projected records;
- making event processing idempotent where possible;
- using stable transaction/event identifiers;
- tracking the last successfully processed height or cursor;
- detecting gaps and duplicates;
- supporting replay from an agreed point or rebuilding from source history;
- exposing indexer lag so the UI can communicate staleness;
- defining how reorganizations or non-finalized data are handled for the chosen protocol.

Exact schemas and algorithms depend on the selected chain interface and version; no specific implementation is asserted here.

## 7. Failure thought experiments

### Experiment A — Concurrent overspend
Balance is 100. Two 80-unit transfers arrive. Explain why UI checks can both pass and why authoritative sequential state validation must reject at least one under no-overdraft semantics.

### Experiment B — Indexer outage
The chain commits a task/reward transaction while a client is offline. Explain how the client can recover current state without becoming an authority.

### Experiment C — Duplicate event
The indexer processes the same committed event twice. Explain what could happen without idempotency or uniqueness safeguards, and propose an invariant for the projection.

### Experiment D — Retry after timeout
A client times out and resubmits the same logical payment. Explain why transport retries and business-level duplicate prevention are different problems, and list the identity information needed to design safe idempotency.

## 8. Interview questions

1. What is the difference between authoritative chain state and a read model?
2. Why should invariants be written before implementation?
3. Why are frontend balance checks insufficient?
4. What does atomicity mean for a transfer?
5. How does replay protection differ from indexer idempotency?
6. What metadata helps an indexer detect missed or duplicated data?
7. Why is a successful API submission not proof of task/reward?
8. What is the difference between available and booked balance, and why must the domain define it?
9. How should an indexer recover after an outage?
10. Which business rules must be resolved before a TeamTask module can be implemented?

## 9. Mastery exercise

- [ ] Write five invariants for the TeamTask domain.
- [ ] Draw the transaction lifecycle and label which stages are authoritative versus derived.
- [ ] Explain concurrent overspend without relying on frontend checks.
- [ ] Design a high-level indexer recovery plan after a multi-hour outage.
- [ ] Explain how a duplicate API retry differs from a duplicate projected event.
- [ ] List at least eight open task/reward-domain questions and mark which must be answered before implementation.
- [ ] Answer all ten interview questions in your own words.

## 10. Evidence and implementation status

- **Conceptual explanations:** included.
- **Runnable lab:** not implemented.
- **Failure experiments:** thought experiments only; no execution results claimed.
- **Cosmos SDK/CometBFT source paths:** intentionally deferred until the initial dependency baseline is verified.
- **Tests/performance results:** none claimed.
- **Next gate:** review the domain questions and invariants. Then prepare a version-verified technical design and a minimal local lab before application implementation.
