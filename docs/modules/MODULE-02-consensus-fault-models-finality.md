# Module 2 — Consensus, Fault Models, and Finality

- **Module ID:** COSMOS-FOUND-002
- **Status:** Conceptual learning draft; not a version-pinned implementation guide
- **Prerequisite:** COSMOS-FOUND-001 — Distributed Systems and Deterministic State Machines
- **Project sequence:** Single-chain settlement first; IBC is deferred

## Learning outcomes

By the end of this module, the learner should be able to:
1. Explain why distributed nodes need consensus.
2. Distinguish crash faults, network faults, and Byzantine faults.
3. Explain quorum intuitively and why fault assumptions matter.
4. State the common one-third Byzantine fault threshold for classical partially synchronous BFT consensus, without treating it as universal.
5. Distinguish transaction inclusion, commitment/finality, and application-level business completion.
6. Explain why finality and availability claims must name protocol assumptions.

## 1. Why consensus is needed

Independent nodes may receive transactions at different times and may propose conflicting transactions. They need a protocol to agree on a canonical ordered history, despite delays and certain failures.

Consensus is not simply “all nodes vote yes on every transaction.” A consensus protocol defines proposal, voting, quorum, rounds/timeouts, commit conditions, and behavior under faults. Exact details depend on the protocol and software version.

## 2. Fault models

### Crash fault
A process stops working or becomes unavailable. It does not intentionally send contradictory or malicious messages while faulty.

### Network fault
Messages may be delayed, lost, duplicated, or delivered in different orders. A network partition can leave groups of nodes unable to communicate.

### Byzantine fault
A faulty participant may behave arbitrarily: send conflicting messages, lie, selectively withhold messages, or coordinate with other faulty participants. Byzantine fault tolerance is stronger than crash fault tolerance.

A production design must state which faults it is intended to tolerate. “Fault tolerant” without a fault model is incomplete.

## 3. Quorums and the one-third intuition

Many classical Byzantine fault-tolerant consensus designs use a threshold based on the assumption that fewer than one-third of voting power is Byzantine. In a common equal-weight model, if there are n validators and at most f Byzantine validators, the familiar condition is n >= 3f + 1. Real proof-of-stake protocols may weight votes by voting power, so validator count alone is not sufficient.

Intuition: a protocol needs enough honest voting power to form a decision quorum while preventing two incompatible decisions from both gathering valid quorums under its assumptions. Exact quorum sizes and safety proofs belong to the specific protocol.

Do not assume the one-third rule applies identically to every consensus algorithm, network model, validator weighting scheme, or implementation.

## 4. Safety and liveness revisited

- **Safety:** conflicting histories are not both committed under the protocol's assumptions.
- **Liveness:** the protocol eventually commits new history when its synchrony and participation assumptions hold.

During severe network delays or partitions, a BFT protocol may stop committing rather than violate safety. Liveness is often conditional on the network eventually becoming sufficiently well-behaved and enough voting power participating.

## 5. Inclusion is not the same as finality

These terms should not be casually conflated:
- **Submitted:** a client has sent a transaction.
- **Received/mempool:** a node has received or queued it; this is not a guarantee of inclusion.
- **Included/proposed:** a block or proposal contains it; a proposal alone may not be committed.
- **Committed/finalized:** the protocol's commit condition has been met.
- **Business-complete:** the application's required business process has completed. This may involve additional off-chain steps even after on-chain finality.

The precise meanings and guarantees vary by protocol. Before exposing a status in a UI, define the status semantics and source of truth.

## 6. Institutional settlement implications

For the Institutional Settlement Network:
- Do not treat an API response such as “submitted” as proof of final settlement.
- Define separate transaction lifecycle states and their semantics.
- Preserve transaction identifiers and enough information to reconcile API responses, chain events, and the PostgreSQL read model.
- Do not let a stale indexer status override authoritative chain state.
- Document what the system does when a transaction is delayed, rejected, committed, or when the indexer is behind.

Business-level completion may need a separate definition from consensus finality—for example, whether downstream bank reconciliation or compliance checks are required.

## 7. Failure thought experiments

### Experiment A — Crash fault
One validator process stops. Ask: can the remaining participating voting power still satisfy the protocol's commit conditions? The answer depends on voting power distribution and protocol assumptions.

### Experiment B — Network partition
Two groups of validators cannot communicate. Ask: what conditions allow one side to commit, and what prevents incompatible commits? Answer only with reference to the chosen protocol's assumptions; do not infer behavior from the word “blockchain.”

### Experiment C — Byzantine validator
A validator sends conflicting votes to different peers. Ask: what protocol rules and quorum assumptions protect safety, and what evidence would be needed to diagnose the behavior?

### Experiment D — Submitted but not committed
The API returns that a transaction was received, but it never becomes part of a committed block. Ask: what should the user interface show? It should not label the transaction “settled” merely because submission succeeded.

## 8. Interview questions

1. What problem does consensus solve?
2. How does a Byzantine fault differ from a crash fault?
3. What is a quorum, and why is voting power sometimes more important than validator count?
4. What does the one-third Byzantine fault threshold mean in common BFT designs, and why is it not a universal law?
5. How can a system preserve safety while temporarily losing liveness?
6. Why is a submitted transaction not necessarily a finalized transaction?
7. What is the difference between protocol finality and business completion?
8. What assumptions must be named when describing a consensus protocol's safety and liveness?
9. What should an explorer display if its indexer is behind the chain?
10. Why must the chosen consensus protocol and version be documented before implementation-specific claims?

## 9. Mastery exercise

- [ ] Explain crash, network, and Byzantine faults in your own words.
- [ ] Explain the one-third intuition without claiming it is universal.
- [ ] Draw a transaction lifecycle from submission to commit and business completion.
- [ ] Describe a partition scenario and state what additional protocol information is needed to analyze it.
- [ ] Define the settlement transaction statuses your product might expose.
- [ ] Answer all ten interview questions without copying this document.

## 10. Evidence and implementation status

- **Conceptual explanations:** included.
- **Diagrams:** transaction lifecycle exercise assigned; no rendered diagram claimed in this draft.
- **Runnable lab:** not implemented.
- **Failure experiments:** thought experiments only; no execution results claimed.
- **Cosmos SDK/CometBFT source paths:** intentionally deferred until the initial dependency baseline is verified.
- **Tests/performance results:** none claimed.
- **Next gate:** learner can explain the fault models, conditional safety/liveness guarantees, and transaction lifecycle. Then continue to Module 3: transactions, application state, and invariant design for single-chain settlement.
