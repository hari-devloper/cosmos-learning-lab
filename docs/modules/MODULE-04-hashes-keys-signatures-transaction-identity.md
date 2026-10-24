# Module 4 — Hashes, Keys, Signatures, and Transaction Identity

- **Module ID:** COSMOS-FOUND-004
- **Status:** Conceptual learning draft; not a version-pinned Cosmos implementation guide
- **Prerequisites:** COSMOS-FOUND-001, COSMOS-FOUND-002, COSMOS-FOUND-003
- **Project sequence:** Single-chain first; IBC is deferred

## Learning outcomes

By the end of this module, the learner should be able to:
1. Explain the role of cryptographic hashes, key pairs, digital signatures, and addresses.
2. Distinguish authentication/authorization from transaction validity and settlement.
3. Explain why the exact bytes being signed matter.
4. Recognize replay-protection and key-management design questions.
5. Avoid treating a transaction hash as proof that a transaction succeeded.

## 1. Hashes: fingerprints of bytes

A cryptographic hash maps input bytes to a fixed-size digest. A secure hash function is designed so that finding collisions or recovering a useful original input is computationally infeasible under its security assumptions.

Useful properties include:
- **Determinism:** the same bytes produce the same digest.
- **Fixed-size output:** digest size does not depend on input length.
- **Preimage resistance:** given a digest, it should be infeasible to find an input that hashes to it.
- **Collision resistance:** it should be infeasible to find two different inputs with the same digest.

A hash does not encrypt data and does not prove who created it. Hashing a transfer request is not a substitute for authorization, validation, or consensus.

## 2. Public/private key pairs and digital signatures

A private key is secret signing material. A public key is used to verify signatures produced by the corresponding private key. A digital signature can establish that the signer had access to the private key for the signed message, assuming a sound scheme and correct implementation.

A signature alone does not establish:
- that the signer is authorized to debit a particular institutional account;
- that the account has sufficient spendable balance;
- that the transaction is not a replay;
- that the transaction has been committed;
- that a business process is complete.

The application must connect the verified signer to an explicit authorization policy and validate the requested state transition.

## 3. Addresses are identifiers, not secrets

A blockchain address is commonly derived from or associated with public-key material, using chain-specific rules. The exact derivation and encoding are implementation-specific. An address should not be treated as a private key, and knowing an address does not enable signing.

For the project, do not invent address derivation rules. Pin them to the chosen Cosmos SDK and key/address configuration once the dependency baseline is approved.

## 4. Why signed bytes and encoding matter

A signature verifies a particular message representation. If two systems serialize logically similar objects into different byte sequences, signatures and hashes may differ. Ambiguous encoding can also create security bugs if different components interpret the same bytes differently.

A transaction format must define:
- which fields are signed;
- canonical or otherwise unambiguous encoding rules;
- how chain/domain context is included;
- how sequence numbers, account numbers, expirations, or nonces are represented where applicable;
- how signature algorithms and public keys are identified.

These details must be taken from the selected protocol version, not guessed from a generic blockchain example.

## 5. Authentication, authorization, validity, and settlement

Keep these questions separate:

1. **Signature verification:** Is the signature valid for the exact bytes and public key?
2. **Identity/account mapping:** Which account or participant does this signer represent?
3. **Authorization:** Is this signer permitted to request this operation?
4. **Transaction validity:** Are required fields, sequence/replay rules, limits, and domain rules satisfied?
5. **State-transition validity:** Does applying the operation preserve all invariants?
6. **Commitment/finality:** Has the transaction become part of the committed chain history under the selected protocol?
7. **Business completion:** Have all required business and operational steps completed?

Passing one stage does not imply passing the later stages.

## 6. Replay protection is not one universal mechanism

A valid signed message can potentially be resubmitted. Systems commonly use sequence numbers, nonces, unique operation IDs, expiry windows, or protocol-specific transaction rules to prevent repeated application. The exact mechanism and scope vary.

For an institutional payment, distinguish:
- **Transport retry:** a client resends because it did not receive a response.
- **Duplicate business request:** the same logical payment is submitted again, perhaps with a new transaction envelope.
- **Consensus/protocol replay:** a transaction or signed payload is replayed within or across contexts.
- **Indexer replay:** an already committed event is delivered to the projection consumer again.

These cases need related but distinct controls. Indexer idempotency cannot replace on-chain replay protection.

## 7. Key management and institutional controls

Key security is part of system design, not merely a cryptography library choice. Questions include:
- Where are private keys generated and stored?
- Which operators or services can use them?
- Are keys held in an HSM or managed signing service?
- How are key rotation, revocation, backup, and recovery handled?
- Can one signer move funds alone, or is threshold/multisig approval required?
- How are signing requests audited and rate-limited?
- What happens if a key is compromised?

Do not place private keys in source control, frontend code, logs, screenshots, or ordinary configuration files. A local learning key is not suitable for real assets or production use.

## 8. Common misconceptions

- **“A transaction hash means it succeeded.”** No; it is an identifier/fingerprint, not a success verdict.
- **“A valid signature means the payment is authorized.”** Not necessarily; authorization policy is a separate check.
- **“A signature means the transaction is final.”** No; commitment/finality is a protocol state.
- **“The address is the secret.”** No; signing requires private-key material.
- **“If the same payment is retried, the transaction hash will always be the same.”** Not necessarily; envelope fields, sequence/nonce, signatures, or encoding can differ. Business-level idempotency may need a stable payment identifier.
- **“A hash hides the original data.”** Hashing is not encryption and may leak equality or enable guessing of low-entropy inputs.

## 9. Failure thought experiments

### Experiment A — Valid signature, unauthorized account
A signer produces a valid signature but is not permitted to debit the requested institutional account. Explain which checks must reject the request.

### Experiment B — Hash confusion
A UI shows a transaction hash immediately after submission. Explain why this does not prove commitment or settlement.

### Experiment C — Client retry
A client times out and resends the same logical payment with a different transaction envelope. Explain why protocol-level duplicate handling and business-level idempotency must both be considered.

### Experiment D — Ambiguous serialization
Two components interpret or encode a transaction differently. Explain why signing and verification must agree on the exact bytes and domain context.

## 10. Interview questions

1. What does a cryptographic hash prove, and what does it not prove?
2. What is the difference between a public key and a private key?
3. What does a digital signature establish?
4. Why is a valid signature not equivalent to authorization?
5. Why does canonical/unambiguous encoding matter?
6. Why is a transaction hash not proof of successful settlement?
7. What is the difference between a protocol nonce/sequence and a business payment ID?
8. How does indexer replay differ from transaction replay?
9. What operational controls should protect institutional signing keys?
10. Why must address derivation and transaction signing details be pinned to the chosen framework version?

## 11. Mastery exercise

- [ ] Explain the difference between hash, signature, address, and transaction ID.
- [ ] Draw the path from signature verification through authorization and state-transition validation.
- [ ] Give an example of a validly signed but invalid or unauthorized transfer.
- [ ] Explain why a transaction hash displayed by a UI does not prove settlement.
- [ ] Propose separate protections for transport retry, business duplicate submission, protocol replay, and indexer replay.
- [ ] List at least five key-management controls.
- [ ] Answer all ten interview questions in your own words.

## 12. Evidence and implementation status

- **Conceptual explanations:** included.
- **Runnable cryptography lab:** not implemented.
- **Cosmos SDK source paths:** intentionally deferred until the dependency baseline is selected and verified.
- **Tests/performance results:** none claimed.
- **Next step:** complete the mastery check, then continue to version-pinned Cosmos SDK/CometBFT architecture and transaction lifecycle after the relevant compatibility gate.
