# ADR-0004: Macro system (syntax-rules) — deferred

- **Status:** WIP
- **Date:** 2026-06-07

## Context

R7RS-small includes `syntax-rules`, a hygienic, pattern-based macro system. Macros are a
high-value learning topic (they're central to Scheme's character) but **hygiene** —
avoiding accidental variable capture between macro-introduced and user identifiers — is
genuinely tricky and is a research-grade chunk. Including it now would dominate the
initial build. Deferred from the first version by decision during scoping.

This ADR records the *open decision* of how to do macros, so it can be picked up later
(a focused `grill-me`).

## Options Considered (to be explored later)

- **`syntax-rules` with full hygiene.** Faithful to R7RS; the hard part is hygiene
  (renaming / marking). Most learning value, most effort.
- **Non-hygienic `defmacro`-style macros.** Much simpler to implement; teaches the
  expansion idea but not hygiene, and can bite with capture. Common pedagogical stepping
  stone.
- **`syntax-rules` patterns without full hygiene first, hygiene later.** Staged.
- **No macros at all.** Keep the interpreter macro-free.

## Decision

Deferred — open. Not in the first version. Needs a dedicated `grill-me` to choose an
approach before it becomes a Work Package.

## Consequences

- First version has no user-defined syntax; everything users want must be a procedure or
  a built-in special form.
- Tracked as a `WIP` roadmap candidate.
