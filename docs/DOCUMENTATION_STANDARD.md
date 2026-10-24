# Documentation Standard

This standard applies to every meaningful concept, lab, experiment, and implementation-related learning task.

## Required concept record

### 1. Metadata
- Stable ID (for example, `COSMOS-001`)
- Title, status, date updated
- Framework/release or commit studied
- Prerequisites and related concept IDs

### 2. Learning objectives
State what the learner should be able to explain, trace, build, test, or diagnose.

### 3. Explanation
Define terminology, explain the mental model, describe why the concept exists, and identify common misconceptions. Separate protocol/framework behavior from application-specific choices.

### 4. Architecture and flow
Include Mermaid or another maintainable source diagram when it improves understanding. Explain the diagram in text; do not rely on the image alone.

### 5. Source-code call path
List the exact source paths and symbols for the release/commit studied. Verify the paths and describe the call sequence. If not yet verified, label it **unverified** and do not present it as fact.

### 6. Lab
Include prerequisites, setup, numbered commands, expected observations, verification checks, cleanup, and reset/re-run instructions. Explain platform-specific differences.

### 7. Tests and evidence
Record exact commands, environment, test conditions, actual outputs or concise summaries, and interpretation. Distinguish expected output from observed output.

### 8. Failure experiments
State the hypothesis, controlled change, expected failure, actual observation, diagnosis, recovery, and lesson. Do not run destructive actions without an explicit safety check.

### 9. Engineering implications
Discuss relevant correctness, security, performance, reliability, observability, and operational trade-offs.

### 10. Interview preparation
Include beginner-to-advanced questions, concise and deep answers, follow-up questions, trade-offs, and common traps.

### 11. Mastery evidence
Record what the learner explained or demonstrated, evidence links, confidence, unresolved questions, and the next revision date. Do not claim mastery without evidence.

## Evidence labels

Use these labels where ambiguity is possible:

- **Verified** — checked against a cited source, command, or test.
- **Observed** — directly produced in the recorded environment.
- **Proposed** — a design recommendation not yet implemented.
- **Inferred** — reasoned conclusion, not direct evidence.
- **Unverified** — requires further checking.

## Documentation/code synchronization

When code changes, update relevant diagrams, call paths, command instructions, tests, and limitations in the same task. Prefer small, reviewable changes. Never fabricate benchmark numbers, outputs, source locations, or security assurances.
