# Module 5 — Architecture diagrams

These diagrams are separated by question. GitHub renders the Mermaid blocks on this page; the editable `.mmd` sources are linked below each diagram.

## 1. Component architecture — who does what?

Follow the main path from client transaction to CometBFT, through ABCI into BaseApp and the module layer, then to committed state. The read side is separate: PostgreSQL and user interfaces are projections, not settlement authorities.

Source: [MODULE-05-cosmos-component-architecture.mmd](MODULE-05-cosmos-component-architecture.mmd)

```mermaid
flowchart TB
  %% COSMOS-ARCH-001 — simplified, transaction-centered component architecture.
  %% Official references:
  %% https://github.com/cosmos/cosmos-sdk-docs/blob/main/docs/learn/beginner/00-app-anatomy.md
  %% https://github.com/cosmos/cosmos-sdk-docs/blob/main/docs/learn/advanced/00-baseapp.md
  %% https://github.com/cosmos/cosmos-sdk-docs/blob/main/docs/learn/advanced/04-store.md
  CLIENT["1. CLIENTS<br/>Institution service · CLI · Wallet"]
  COMET["2. COMETBFT NODE<br/>P2P · Mempool · Proposals · Consensus"]
  ABCI["ABCI / ABCI++ boundary"]
  BASE["3. COSMOS SDK — BASEAPP<br/>Transaction execution · Gas · Context · Routing"]
  ANTE["AnteHandler<br/>Transaction-level checks"]
  ROUTER["Message router"]
  MSG["ISN MsgServer<br/>Settlement rules · Authorization · Invariants"]
  SDKMOD["Enabled SDK modules<br/>Auth · Bank · Governance · Upgrade"]
  KEEP["Keepers<br/>Controlled access to module state"]
  STORE[("4. APPLICATION STATE<br/>CommitMultiStore + module KV stores")]
  COMMIT["Commit successful state changes"]
  APPHASH["Application state commitment / app hash"]
  BLOCKS[("Committed blocks and transaction results")]
  INDEXER["5. DERIVED READ SIDE<br/>Indexer · Deduplication · Checkpoint"]
  PG[("PostgreSQL<br/>Queryable projection")]
  UI["Explorer · Bank Portal · Treasury / Ops Consoles"]
  RECON["Reconciliation<br/>Detect gaps; repair projection from authoritative data"]
  WIRING["Application assembly<br/>cmd / app constructor · codecs · Module Manager<br/>registers services, stores, dependencies, hook order"]

  CLIENT -->|"signed transaction"| COMET
  COMET <--> ABCI
  ABCI <--> BASE
  BASE --> ANTE --> ROUTER
  ROUTER --> MSG
  ROUTER --> SDKMOD
  MSG --> KEEP
  SDKMOD --> KEEP
  KEEP --> STORE --> COMMIT --> APPHASH
  APPHASH -. "reported to consensus protocol" .-> COMET
  COMET --> BLOCKS --> INDEXER --> PG --> UI
  BLOCKS --> RECON
  PG --> RECON
  WIRING -. "constructs and configures" .-> BASE
  WIRING -. "registers modules/services" .-> ROUTER

  classDef protocol fill:#e8efff,stroke:#315ea8,color:#122b50,stroke-width:2px;
  classDef application fill:#e7f5eb,stroke:#287443,color:#12351d,stroke-width:2px;
  classDef storage fill:#fff2d9,stroke:#a45d00,color:#492900,stroke-width:2px;
  classDef derived fill:#f0eaff,stroke:#7550a1,color:#2e1a4a;
  classDef assembly fill:#f3f4f6,stroke:#687386,color:#202938;
  class COMET,ABCI protocol;
  class BASE,ANTE,ROUTER,MSG,SDKMOD,KEEP application;
  class STORE,COMMIT,APPHASH,BLOCKS storage;
  class INDEXER,PG,UI,RECON derived;
  class WIRING assembly;
```

## 2. Transaction lifecycle — what happens to one transaction?

This view focuses on admission, consensus, execution outcome, commit, indexing, and reconciliation. It does not try to explain genesis, upgrades, or node-retention policy.

Source: [MODULE-05-transaction-lifecycle-deep-dive.mmd](MODULE-05-transaction-lifecycle-deep-dive.mmd)

```mermaid
flowchart TB
  CLIENT["Client builds and signs transaction"]
  CHECK["CheckTx / mempool admission"]
  PROPOSE["Proposal processing and consensus"]
  FINALIZE["FinalizeBlock / transaction execution"]
  ANTE["AnteHandler"]
  MSG["Message router and handler"]
  RESULT{"Business message succeeds?"}
  APPLY["Apply message-state writes"]
  FAIL["Do not apply failed message-state writes<br/>Ante effects depend on configured cache/SDK semantics"]
  COMMIT["Commit application state"]
  SOURCE[("Committed block and results")]
  INDEX["Indexer: deduplicate, persist rows and checkpoint"]
  PG[("PostgreSQL projection")]
  RECON["Reconciliation against authoritative chain source"]

  CLIENT --> CHECK --> PROPOSE --> FINALIZE --> ANTE --> MSG --> RESULT
  RESULT -->|yes| APPLY --> COMMIT
  RESULT -->|no| FAIL --> COMMIT
  COMMIT --> SOURCE --> INDEX --> PG --> RECON

  classDef protocol fill:#e8efff,stroke:#315ea8,color:#122b50,stroke-width:2px;
  classDef application fill:#e7f5eb,stroke:#287443,color:#12351d,stroke-width:2px;
  classDef derived fill:#f0eaff,stroke:#7550a1,color:#2e1a4a;
  class CHECK,PROPOSE,FINALIZE protocol;
  class ANTE,MSG,RESULT,APPLY,FAIL,COMMIT application;
  class SOURCE,INDEX,PG,RECON derived;
```

## 3. Chain operations lifecycle — how is the chain initialized, upgraded, and recovered?

This is an operator/lifecycle view, separate from both the component map and the transaction flow. It highlights the dependency between retained history and the ability to rebuild derived data.

Source: [MODULE-05-chain-operations-lifecycle.mmd](MODULE-05-chain-operations-lifecycle.mmd)

```mermaid
flowchart TB
  %% COSMOS-ARCH-001 — chain lifecycle and operator responsibilities.
  %% Exact commands and hooks depend on the selected release and node configuration.
  GENESIS["Genesis document<br/>initial chain configuration and state"]
  INIT["Application startup<br/>construct stores, keepers, modules, services"]
  RUN["Running chain<br/>blocks committed; state transitions applied"]
  UPGRADE["Planned upgrade height<br/>coordinated rollout"]
  MIGRATE["Versioned module migrations<br/>transform persisted state"]
  RESUME["Resume on upgraded binary"]
  SNAPSHOT["Snapshots / state sync<br/>bootstrap or recover a node"]
  PRUNE["Pruning policy<br/>retain selected local history"]
  ARCHIVE[("Archive / retained source<br/>historical queries and replay")]
  INDEXER["Indexer backfill / projection rebuild"]
  RECON["Reconcile rebuilt projection<br/>against authoritative chain source"]
  NOTE["Pruning does not reverse settlement.<br/>A rebuild requires sufficient retained authoritative history.<br/>Test restore, upgrade, and replay procedures before relying on them."]

  GENESIS --> INIT --> RUN
  RUN -->|"approved upgrade plan"| UPGRADE --> MIGRATE --> RESUME --> RUN
  RUN -. "recovery artifacts" .-> SNAPSHOT
  SNAPSHOT -. "restore / bootstrap node" .-> INIT
  RUN --> PRUNE
  RUN --> ARCHIVE
  ARCHIVE --> INDEXER --> RECON
  PRUNE -. "may limit local replay" .-> ARCHIVE
  PRUNE --- NOTE
  RECON --- NOTE

  classDef lifecycle fill:#e8efff,stroke:#315ea8,color:#122b50,stroke-width:2px;
  classDef recovery fill:#e7f5eb,stroke:#287443,color:#12351d,stroke-width:2px;
  classDef caution fill:#fff2d9,stroke:#a45d00,color:#492900,stroke-width:2px;
  class GENESIS,INIT,RUN,UPGRADE,MIGRATE,RESUME lifecycle;
  class SNAPSHOT,ARCHIVE,INDEXER,RECON recovery;
  class NOTE,PRUNE caution;
```

## Authority rules

- **Committed application state and committed chain records are authoritative** for settlement under the chosen chain design.
- **PostgreSQL, indexer checkpoints, and UI views are derived.** Detect discrepancies and repair/rebuild the projection from verified source data; do not alter chain balances to make SQL match.
- **CheckTx acceptance is not proof of settlement.** Admission, inclusion, execution result, commitment, and read-model visibility are separate facts.
- **IBC remains deferred** until single-chain settlement invariants, replay, failure handling, indexing, reconciliation, and recovery have evidence.

## Evidence status

These are conceptual architecture diagrams, not proof that a specific release has been built or tested. The dependency baseline is not yet approved. Review the rendered diagrams in GitHub before treating the layout as final.
