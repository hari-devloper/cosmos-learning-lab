# TeamTask Chain — Fixed Project Plan

## 1. Purpose

Build one small, integrated application from scratch to learn Cosmos SDK and major parts of the Cosmos ecosystem. The goal is deep understanding and interview readiness, not a production product.

This repository is the learning lab. The Institutional Settlement Network is a separate project.

## 2. Use case

A team leader creates tasks, assigns them to team members, sets a deadline and reward, and funds a reward pool. The assigned member progresses the task and submits it for review. The leader approves or rejects the submission. Approval transfers the configured reward once; rejection records feedback and returns the task to in-progress.

### Roles

- **Leader:** creates tasks, assigns members, sets deadlines/rewards, funds the reward pool, and approves/rejects submissions.
- **Assigned member:** views assigned work, performs allowed status transitions, and submits work for review.
- **Other accounts:** cannot mutate tasks unless a future documented rule explicitly allows it.

### Core invariants to design before implementation

- Only the leader can create and assign tasks or make final review decisions.
- Only the assigned member can perform member workflow actions on a task.
- Only valid task-state transitions are accepted.
- Only an approved task can trigger a reward transfer.
- A task can never be rewarded twice.
- Rewards must come from an explicitly funded pool; do not assume unlimited minting.
- Rejected submissions store review feedback and return to the defined in-progress state.
- Chain validation is authoritative; hiding a button in the frontend does not enforce authorization.

The exact state names, deadlines, reward denomination, pool funding, and message/query schemas remain design decisions until reviewed in the relevant milestone.

## 3. Learning rules

1. No Codex in this repository.
2. One milestone at a time; each has a written exit gate.
3. Explain and design before implementation.
4. Verify the actual checkout, version, source paths, APIs, and commands before relying on them.
5. Use the official Cosmos example only as a reference experiment, not as a substitute for our own implementation.
6. Document every meaningful concept, diagram, command, source call path, decision, test, failure, fix, and interview question as the work happens.
7. Clearly label proposed, verified, observed, inferred, and unverified claims.
8. Never invent output, test results, benchmarks, or source locations.
9. Avoid scaffolding future components before the current milestone needs them.
10. Changes to the approved plan require a reason and explanation first.

## 4. Required learning cycle

**Understand → Design → Implement → Test → Trace source → Document → Mastery check.**

A milestone is complete only when the implementation/lab works, actual evidence is recorded, the relevant documentation is current, and the learner can explain the concept and its failure cases.

## 5. Fixed milestone sequence

1. Repository and environment baseline.
2. Cosmos application foundations: SDK, CometBFT, app wiring, genesis, startup.
3. Transaction lifecycle: messages, protobuf, signing, fees, gas, execution, events.
4. TeamTask domain model: roles, states, transitions, authorization and invariants.
5. Custom TeamTask module: keepers, store, messages, queries, events and errors.
6. Reward pool and bank integration; funded, exactly-once reward transfer.
7. CLI/query interfaces, gRPC/REST and transaction inspection.
8. React frontend and wallet signing/broadcasting/status.
9. Tests and controlled failure experiments.
10. Multi-validator network and consensus.
11. Governance, parameters, upgrades and state migration.
12. CosmWasm contract fundamentals and a small related experiment.
13. IBC fundamentals with two compatible local chains.
14. ICS-20 and advanced interchain experiments such as ICA/IBC-enabled contracts.
15. Security, operations, recovery, performance and interview mastery.

Detailed objectives and exit gates are maintained in LEARNING_ROADMAP.md.

## 6. Documentation deliverables

For each concept or milestone, keep the relevant records updated with:

- Concept, motivation, mental model, terminology, and misconceptions.
- Architecture and sequence/state diagrams where useful.
- Exact version/release and verified source-code call path.
- Design decisions and alternatives considered.
- Reproducible setup and commands with expected versus observed results.
- Tests, failure experiments, diagnosis, recovery, and regression coverage.
- Security, correctness, gas/performance, and operational implications where relevant.
- Interview questions, answers, follow-ups, and mastery evidence.
- Open questions and the next concrete step.

## 7. Keep the POC simple

The first release is a single Cosmos chain with one custom task module. Use the standard auth and bank modules; do not introduce a database, microservices, event streaming, a custom token, NFT badges, or a complex frontend.

Build in this order:
1. Run the dependency-free domain tests under `poc/domain/`.
2. Scaffold the actual chain in `chain/` using `docs/POC_BUILD_GUIDE.md`.
3. Implement task messages, state, queries and signer checks in the custom module.
4. Connect funded reward payout to the standard bank keeper.
5. Verify the full flow with CLI transactions and automated tests.
6. Add a UI only after the chain path works.
7. Study multi-validator consensus, governance/upgrades, CosmWasm, IBC/ICS-20 and operations in focused labs without expanding the core application.

The coverage contract for major topics lives in `docs/CONCEPT_COVERAGE_MATRIX.md`. Every concept needs an explanation plus relevant source tracing, lab/test evidence, failure cases and interview practice. This does not mean every topic must become production code in the task module.

## 8. Version decision and current status

**Selected implementation target:** Cosmos SDK `v0.55.0` with CometBFT `v0.40.0`, matching the dependency pair inspected in the existing official Cosmos example reference. The reference is for learning architecture only; TeamTask Chain must be implemented in this repository.

The earlier SDK `v0.54.4` / CometBFT `v0.39.4` candidate is superseded. Validate the chosen pair during the real implementation build, tests, and local-chain smoke test. Do not run a separate disposable compatibility experiment or repeat the version audit unless a concrete issue requires it.

**Current:** the domain model and tests are committed on the POC branch. The chain scaffold, build, local startup and domain-to-module integration remain pending. Follow `docs/POC_BUILD_GUIDE.md`; do not mark the milestone complete until commands have actually run and output is recorded.
