# ADR-0003: Numeric model is int64 (exact) + float64 (inexact)

- **Status:** Accepted
- **Date:** 2026-06-07

## Context

R7RS specifies a full numeric tower: exact/inexact, arbitrary-precision integers,
rationals, and complex numbers. Implementing the whole tower is substantial work and
dominated by mechanical edge cases. The learning goal values the *exact/inexact
distinction* (a distinctive Scheme idea) far more than bignum/rational/complex machinery.

A sub-decision, **integer overflow behaviour** for int64 (wrap silently / error /
promote to float64), was resolved on 2026-06-07 — see the Decision section.

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

int64 overflow behaviour: **raise an error** on `+ - *` overflow rather than wrapping
or promoting. Rationale — principle of least surprise: silent wrap-around gives a
mathematically wrong result, and promoting to float64 silently loses exactness; both are
surprising. Erroring is the honest behaviour given the deliberate absence of bignums.
The arbitrary-precision fix lives in the deferred numeric-tower roadmap item.

## Consequences

- The exact/inexact predicates and conversions are meaningful and implementable cheaply.
- Large integer arithmetic is limited by int64 range until/unless the tower is extended.
- Deviates from R7RS-small (which mandates more of the tower); Spec §(numbers) to record
  this. Fuller numeric tower is a roadmap candidate.
