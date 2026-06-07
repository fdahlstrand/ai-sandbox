# ADR-0006: Error model — explicit Go errors to a boundary handler

- **Status:** Accepted
- **Date:** 2026-06-07

## Context

Runtime errors (unbound symbol, `car` of a non-pair, arity mismatch, type errors in
builtins) need a propagation strategy. The choice pervades the evaluator's signature and
is awkward to change later. There is also the question of whether errors are catchable
*within Scheme* (a condition system) or only at the host boundary.

## Options Considered

- **Explicit Go errors to a boundary (chosen).** `eval` returns `(Value, error)`; errors
  bubble up and are caught at the REPL/file boundary. REPL prints and continues; file
  mode aborts with a message. Idiomatic Go, explicit control flow.
- **panic / recover.** Keeps `eval` signature clean (`Value`), raises via `panic`,
  recovers at the boundary. Less plumbing but uses panic for ordinary control flow.
  Rejected for now in favour of explicit, teachable control flow.
- **Scheme condition system now (raise / guard / with-exception-handler).** Most faithful
  to R7RS; sizable. Deferred to the roadmap.

## Decision

`eval` and builtins return `(Value, error)`. Errors propagate explicitly to a boundary
handler. At the **REPL** the error is printed and the loop continues with the next datum.
In **file mode** the error message is printed and the process exits non-zero.

## Consequences

- The eval loop and every builtin thread an `error` return.
- No in-Scheme exception handling initially (`guard`/`raise` are a roadmap candidate).
- Error messages are a host concern (Go strings), not Scheme condition objects, for now.
