# TeamTask Chain — POC Build Guide

## Goal

Build one small, real Cosmos SDK chain for the TeamTask workflow. Do not build a production platform or introduce microservices, Kafka, a database, or a separate backend. The chain is the source of truth; the optional UI comes later.

## Smallest useful POC

1. **Domain rules:** dependency-free Go model and tests in `poc/domain/`.
2. **Real chain:** scaffold a Cosmos SDK app with Ignite CLI v30.x and verify the generated dependency versions.
3. **Task module:** implement create, start, submit, approve and reject; enforce signer-based permissions and valid state transitions on-chain.
4. **Rewards:** use the standard bank module and an explicitly funded reward pool; transfer reward only on approval, in the same transaction as the task update.
5. **Read and demonstrate:** use CLI/query APIs first. Add a frontend only after the chain workflow passes.
6. **Advanced labs:** add consensus, upgrades, CosmWasm and IBC one at a time without complicating the core task module.

## Toolchain decision

The Ignite CLI v30 release notes state that its Cosmos tooling scaffolds chains using Cosmos SDK v0.55.x, IBC-Go v11.2.x and CometBFT v0.40.x. The prior compatibility draft incorrectly treated the inspected IBC-Go v11.2.0 / wasmd v0.70.5 dependency declarations as a reason to defer IBC indefinitely. Use the scaffold as the compatibility baseline, but verify the generated `go.mod`, build and tests before claiming success.

Official references:
- Ignite CLI release notes: https://github.com/ignite/cli/blob/main/changelog.md
- Ignite Cosmos scaffold documentation: https://docs.ignite.com/v29/cli-commands/cli-commands
- Cosmos SDK v0.55 upgrade notes: https://github.com/cosmos/cosmos-sdk/blob/main/UPGRADING.md

## WSL steps

First inspect the installed command and its version:

```bash
command -v ignite
ignite version
```

Use the v30 Cosmos namespace. From the repository root, scaffold into a dedicated subdirectory so learning documents remain at the root:

```bash
ignite cosmos scaffold chain github.com/hari-devloper/cosmos-learning-lab/chain --path ./chain --address-prefix teamtask
```

If your installed Ignite does not expose the `cosmos` namespace, stop and check its version instead of guessing commands or mixing CLI generations.

Then inspect the generated app and dependencies:

```bash
cd chain
go version
go list -m github.com/cosmos/cosmos-sdk github.com/cometbft/cometbft github.com/cosmos/ibc-go/v11
go test ./...
```

These are instructions, not recorded test results. Save actual output and resolve any scaffold or dependency errors before proceeding.

## Module design (keep it small)

Task fields: ID, leader, assignee, title, deadline, reward amount, status, review feedback, and reward-paid marker.

State flow:

```text
assigned -> in_progress -> in_review -> completed
                              |
                              +-- rejected with feedback -> in_progress
```

Messages:
- `CreateTask`: leader creates a task with assignee, deadline and reward.
- `StartTask`: only the assigned member can start.
- `SubmitTask`: only the assigned member can submit from in-progress.
- `ReviewTask`: only the leader can approve or reject; rejection requires feedback.
- Approval: call the bank keeper to transfer the funded reward and mark the task completed/rewarded atomically.

Queries:
- `Task(id)`
- `TasksByAssignee(address)`

Avoid extra roles, configurable workflow engines, NFT badges, DAO features, or a separate reward token for the first POC.

## Completion gate for the first real chain milestone

Do not mark the POC complete until:
- the generated chain builds and starts locally;
- task messages mutate committed chain state;
- invalid signer/state-transition attempts fail;
- the reward pool must be funded;
- approval transfers the reward once, while rejection pays nothing;
- automated tests pass and exact output is recorded;
- docs and diagrams reflect the actual source paths and behavior.
