# ADR-0001: Anchor the dialect to a subset of R7RS-small

- **Status:** Accepted
- **Date:** 2026-06-07

## Context

The project is a learning-oriented Scheme interpreter in Go. Its primary purpose is
to build understanding of the Scheme language, not to ship a production interpreter.
That pedagogical goal is the prime directive: decisions should favour clarity and
illustrativeness over completeness, performance, or edge-case fidelity.

A "Scheme interpreter" needs a north star defining what is in-scope and what "done
enough" means. Without an anchor, scope drifts and "is this correct?" has no referent.
The candidate anchors differ in how much surface they imply and how much ceremony they
impose on a learning toy.

The *anchor* is decided here. The precise **subset boundary** within R7RS-small is
still being scoped during this grilling session (which forms, types, and builtins are
in vs out) and will land in `SPECIFICATION.md` / `ROADMAP.md`.

## Options Considered

- **Ad-hoc minimal (Norvig lis.py style).** A tiny, standard-free subset. Maximum
  learning-per-line, minimum ceremony. Rejected: no citable referent for correctness,
  and naming would not align with real Scheme.
- **R7RS-small subset (chosen anchor).** Anchor to the modern small standard and
  implement a chosen subset. Gives a real, citable definition of "done enough" and
  current-Scheme-aligned naming. Risk: spec-completeness can pull focus from learning;
  mitigated by explicitly scoping the subset and deferring heavy features.
- **R5RS subset.** Classic, compact, teaching-tradition (SICP-era). Simpler than R7RS
  (no library system). Rejected in favour of R7RS-small for modern alignment.

## Decision

Anchor to a **subset of R7RS-small**. (Subset boundary still being scoped — see Spec /
Roadmap once written.)

## Consequences

- Naming and semantics follow R7RS-small where implemented.
- Heavy R7RS features (full numeric tower, hygienic macros, full continuations,
  ports/library system) are explicit candidates for *exclusion or deferral* and will be
  tracked as their own ADRs / roadmap entries.
- "Is this correct?" is answerable by reference to the R7RS-small report.
- Spec §(dialect) to be written to record the chosen subset.
