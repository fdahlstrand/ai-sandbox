# ADR-0002: Tree-walking evaluator with an explicit tail-call loop

- **Status:** Accepted
- **Date:** 2026-06-07

## Context

The interpreter must evaluate parsed Scheme s-expressions. Two coupled questions:

1. **Evaluation strategy.** Tree-walking interpreter (walk the parsed AST/s-expression
   directly) vs compiling to bytecode for a VM. For a learning project, the tree-walker
   maps most directly onto the language semantics and is the standard pedagogical choice.

2. **Proper tail calls.** R7RS *requires* proper tail calls, and they are a defining
   feature of Scheme (iteration is expressed as tail recursion). Supporting them changes
   the *shape* of `eval`: a naive recursive `eval` grows the Go call stack per Scheme
   call and overflows on deep tail recursion. Proper tail calls require `eval` to be an
   explicit loop that, in tail position, rebinds the current expression+environment and
   continues the loop instead of recursing.

## Options Considered

- **Tree-walking, naive recursion (no TCO).** Simplest possible core. Rejected:
  overflows on deep tail recursion and omits a defining Scheme concept central to the
  learning goal.
- **Tree-walking, explicit tail-call loop (chosen).** `eval` is a `for` loop; non-tail
  subexpressions recurse normally, tail expressions update `(expr, env)` and `continue`.
  Constant stack for tail recursion. Modest extra design effort concentrated in one loop.
- **Bytecode VM.** More machinery, further from the source semantics. Rejected as
  over-engineering for the learning goal (possible far-future roadmap item).

## Decision

Tree-walking evaluator with an **explicit iterative eval loop** that implements
**proper tail calls** from the start.

## Consequences

- `eval` is written as a loop, not as direct Go recursion over every call.
- Tail position must be identified for each special form (e.g. last expr of a body, the
  taken branch of `if`, last expr of `begin`/`let`/`cond`).
- Tail-recursive Scheme runs in constant Go stack.
- Full `call/cc`-style continuations are a separate, harder problem (see future ADR) and
  not implied by this decision.
