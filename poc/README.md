# TeamTask Chain — Minimal POC

This folder starts the POC with a **small, dependency-free task-domain model**. It is deliberately not presented as a blockchain. The Cosmos chain will be scaffolded in `chain/` with the selected Ignite/Cosmos toolchain, then this domain behavior will be implemented and tested in a custom SDK module.

## Why this first slice exists

We can prove the task rules with fast, readable Go tests before adding consensus, protobuf generation, wallet UX, and chain wiring. Keep the model small; the chain remains the authoritative place for authorization, state transitions, and reward transfers.

## Run the domain tests

From this directory:

```bash
cd domain
go test ./...
```

## POC workflow

1. Verify the Go domain tests and understand each invariant.
2. Generate the real Cosmos SDK chain in `chain/` using the instructions in `docs/POC_BUILD_GUIDE.md`.
3. Move the domain rules into the chain's custom task module; do not duplicate business logic in the frontend.
4. Add the bank-module reward transfer, CLI/query flow, and end-to-end test.
5. Add advanced topics as small labs in the same repository: multi-validator consensus, governance/upgrades, CosmWasm, IBC/ICS-20, and operations.

**Evidence rule:** this folder's presence is not proof that tests passed. Record actual command output in the mastery tracker after running it in the WSL environment.
