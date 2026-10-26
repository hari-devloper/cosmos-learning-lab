# Cosmos Learning Lab

A hands-on learning project for understanding the Cosmos ecosystem by building one integrated application: **TeamTask Chain**.

## Fixed learning project

TeamTask Chain is a team task-management application in which the chain enforces task ownership, role permissions, workflow transitions, leader review, and one-time token rewards. The frontend makes the workflow usable; it is not the security boundary.

This repository is for learning and interview mastery. It is separate from the Institutional Settlement Network project.

## Non-negotiable working rules

1. **No Codex for this learning project.** We learn, design, write, test, and debug the implementation together.
2. **One milestone at a time.** Do not start the next milestone until the current one passes its stated exit gate.
3. **Explain before implementing.** Understand the concept, design, relevant source path, and expected behavior first.
4. **Verify before instructing.** Check the actual repository, installed toolchain, framework version, command, and source path. Do not guess.
5. **Keep one evolving application.** Add capabilities to TeamTask Chain rather than creating unrelated demos or changing the use case.
6. **Document continuously.** Update explanations, diagrams, call paths, commands, tests, failures, troubleshooting, and interview notes alongside the work.
7. **No fabricated evidence.** Mark expected outcomes separately from observed output. Never claim a test passed unless it was run.
8. **No unnecessary complexity.** Introduce a component only when the current milestone needs it and its purpose is explained.
9. **Changes to the plan require an explanation first.** Keep the approved sequence stable unless a real dependency requires a change.

## Learning cycle

**Understand → Design → Implement → Test → Trace source → Document → Mastery check.**

A milestone is complete only when its code or lab works, the learner can explain the concepts, evidence is recorded, and the exit gate is met.

## Repository map

- docs/TEAMTASK_CHAIN_PROJECT_PLAN.md — product scope, rules, and acceptance gates
- docs/LEARNING_ROADMAP.md — simple staged curriculum mapped to the project
- docs/POC_BUILD_GUIDE.md — concrete commands and the first-chain implementation path
- docs/CONCEPT_COVERAGE_MATRIX.md — explicit coverage contract for major Cosmos concepts
- docs/DOCUMENTATION_STANDARD.md — required documentation and evidence format
- docs/modules/ — existing foundational learning chapters
- docs/diagrams/ — editable diagrams and previews
- docs/ENVIRONMENT_AUDIT.md — environment findings
- docs/TOOLCHAIN_COMPATIBILITY_DRAFT.md — selected SDK/CometBFT target and implementation validation evidence
- poc/domain/ — small dependency-free Go task model and unit tests; not a blockchain by itself
- progress/ — milestone and mastery tracking

Keep the core task application small. Advanced topics are separate focused labs in this same repository; do not add them as dependencies to the core task workflow.

## Version target

Use Ignite CLI v30.x Cosmos tooling as the scaffold baseline for SDK v0.55.x, IBC-Go v11.2.x, and CometBFT v0.40.x. Verify the generated go.mod and actual build/tests; version notes are not proof that the application builds.

## Current project status

**POC started on branch `poc/teamtask-minimal`.** A small Go domain model and tests have been added as the first slice. The real Cosmos chain has not yet been scaffolded or tested. Follow `docs/POC_BUILD_GUIDE.md`, then record only observed build/test results. See `docs/CONCEPT_COVERAGE_MATRIX.md` for the agreed major-topic coverage.
