# Cosmos Learning Lab

A hands-on, from-scratch learning project for understanding the Cosmos ecosystem by building one integrated application: **TeamTask Chain**.

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

- docs/TEAMTASK_CHAIN_PROJECT_PLAN.md — product scope, rules, milestone sequence, and acceptance gates
- docs/LEARNING_ROADMAP.md — ordered Cosmos curriculum mapped to the project
- docs/DOCUMENTATION_STANDARD.md — required documentation and evidence format
- docs/modules/ — existing foundational learning chapters
- docs/diagrams/ — editable diagrams and previews
- docs/ENVIRONMENT_AUDIT.md — environment findings
- docs/TOOLCHAIN_COMPATIBILITY_DRAFT.md — selected SDK/CometBFT target and implementation validation evidence
- progress/ — milestone and mastery tracking

Create implementation directories only when a milestone needs them. Do not scaffold the entire system in advance.

## Version target

- Cosmos SDK `v0.55.0`
- CometBFT `v0.40.0`

This pair matches the dependency versions inspected in the existing official Cosmos example reference. The example is a reference only; TeamTask Chain is implemented here. The selected versions still need to pass the actual project's build, tests, and local-chain smoke test.

## Current project status

**Version decision made; implementation is next.** The existing foundational chapters remain reference material, not proof that the corresponding concepts are mastered. Build the TeamTask Chain application skeleton against the selected versions and record real verification results. Do not claim an implementation or test is complete before it has been run and verified.
