# Agent and contributor instructions

## Learning before implementation

- Teach and discuss the relevant concept before proposing implementation work.
- Do not use Codex to bypass the learning and design gates.
- Before a meaningful implementation task, state the concept, assumptions, acceptance criteria, and verification plan.
- Wait for explicit user approval when a task crosses a design or implementation gate.

## Documentation is mandatory

For each concept used or learned, add or update documentation in this repository. A useful concept record should cover:

1. Learning objectives and prerequisites
2. Precise definitions and mental model
3. Architecture/context diagram where helpful
4. Execution or source-code call path pinned to the studied release
5. A small example and a reproducible lab
6. Commands and expected observations
7. Tests and evidence actually obtained
8. Failure experiments and troubleshooting
9. Security, performance, and operational implications
10. Interview questions, answers, common misconceptions, and revision prompts
11. Mastery evidence and remaining questions

Adapt the template to the concept, but do not silently omit relevant sections. Mark non-applicable sections with a reason.

## Evidence and accuracy

- Never fabricate source paths, outputs, tests, benchmarks, or mastery claims.
- Label statements as verified, inferred, proposed, or unverified when the distinction matters.
- Pin source-code references to a tag or commit and verify that the referenced path exists.
- Record operating system, tool versions, commands, and test conditions for reproducibility.
- Keep documentation synchronized with code changes in the same work unit.

## Repository boundary

This repository is the learning handbook and lab environment. The product repository is `institutional-settlement-network`. Do not copy or share source directories, build outputs, or dependency graphs between repositories. Cross-reference concepts by stable IDs such as `COSMOS-001`.

## Quality gate

Do not mark a module complete until the learner can explain the concept, reproduce its lab, interpret the relevant test or failure, and answer its review questions without relying entirely on notes.
