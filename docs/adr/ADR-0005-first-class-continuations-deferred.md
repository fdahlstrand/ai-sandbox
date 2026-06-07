# ADR-0005: First-class continuations (call/cc) — deferred

- **Status:** WIP
- **Date:** 2026-06-07

## Context

R7RS-small includes `call-with-current-continuation` (`call/cc`), giving first-class,
fully re-invocable continuations. In a direct tree-walking interpreter (ADR-0002) this
is hard: capturing a continuation means capturing the entire pending computation, which
a naive Go-stack-based evaluator does not have as a reified value. Realistic approaches
require either converting the evaluator to continuation-passing style (CPS), or some
form of explicit control stack / stack copying.

Note: ADR-0002's tail-call loop handles *tail* calls but does **not** give us first-class
continuations — they are a separate, harder problem. Deferred from the first version.

## Options Considered (to be explored later)

- **CPS evaluator.** Rewrite eval in continuation-passing style; continuations become
  first-class naturally. Large structural change to the core.
- **Explicit control-stack evaluator (defunctionalized CPS).** Reify the control stack so
  continuations capture/restore it. Also a significant rewrite.
- **Escape-only continuations** (e.g. via Go `panic`/`recover`). Supports upward/escape
  use of `call/cc` only — cheap, partial, not re-entrant.
- **No continuations.** Omit entirely.

## Decision

Deferred — open. Not in the first version. Needs a dedicated `grill-me`; the chosen
approach interacts strongly with the evaluator structure decided in ADR-0002.

## Consequences

- First version has no `call/cc`; non-local control flow limited to whatever the chosen
  error model offers (see error-handling ADR).
- Tracked as a `WIP` roadmap candidate.
