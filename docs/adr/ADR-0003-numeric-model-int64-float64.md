# ADR-0003: Numeric model is int64 (exact) + float64 (inexact)

- **Status:** Accepted
- **Date:** 2026-06-07

## Context

R7RS specifies a full numeric tower: exact/inexact, arbitrary-precision integers,
rationals, and complex numbers. Implementing the whole tower is substantial work and
dominated by mechanical edge cases. The learning goal values the *exact/inexact
distinction* (a distinctive Scheme idea) far more than bignum/rational/complex machinery.

A sub-decision left open: **integer overflow behaviour** for int64 (wrap silently /
error / promote to float64). Not yet decided.

## Options Considered

- **float64 only.** Simplest, but discards the exact/inexact distinction — teaches less
  about Scheme's numeric character. Rejected.
- **int64 (exact) + float64 (inexact) (chosen).** Two number types. Supports
  `exact?`, `inexact?`, `exact->inexact` etc. No bignum, rational, or complex. Best
  learning-per-complexity ratio.
- **Fuller tower (bignum via math/big + rationals).** More faithful, noticeably more
  code and edge cases. Deferred to the roadmap rather than rejected.

## Decision

Two number types: **int64 for exact integers, float64 for inexact reals.** No
arbitrary-precision integers, rationals, or complex numbers.

Open sub-decision: int64 overflow behaviour (wrap / error / promote) — to be resolved
during implementation; lean toward erroring or promoting rather than silent wrap.

## Consequences

- The exact/inexact predicates and conversions are meaningful and implementable cheaply.
- Large integer arithmetic is limited by int64 range until/unless the tower is extended.
- Deviates from R7RS-small (which mandates more of the tower); Spec §(numbers) to record
  this. Fuller numeric tower is a roadmap candidate.
