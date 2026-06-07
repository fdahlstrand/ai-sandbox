# Scheme Interpreter

A learning-oriented Scheme interpreter written in Go. This glossary pins the Scheme
domain vocabulary so code, comments, and docs use one consistent word per concept. The
domain *is* the Scheme language, so the canonical terms here track R7RS-small usage.

## Language

**Datum**:
A single piece of Scheme data produced by reading source text — a number, symbol,
string, char, boolean, pair, or vector. The output of the reader, before evaluation.
_Avoid_: token (a token is a lexical fragment, not a datum), value

**S-expression**:
The parenthesized list/atom notation that is both Scheme's source syntax and its data.
The textual/structural form a datum takes.
_Avoid_: sexp, form (reserve "form" for the evaluation sense below)

**Atom**:
A datum that is not a pair — a number, symbol, string, char, or boolean.

**Pair**:
A two-field cons cell (car, cdr). Chains of pairs ending in the empty list form lists.
_Avoid_: cons cell, cons

**Symbol**:
An interned identifier datum (e.g. `lambda`, `x`, `+`). Used as variable names and as
quoted data.

**Special form**:
A built-in syntactic construct evaluated by its own rule rather than as a procedure call
(e.g. `if`, `define`, `lambda`, `quote`, `let`, `cond`). Distinct from a procedure
because its operands are not all evaluated first.
_Avoid_: keyword, syntax (reserve "syntax" for the future macro layer)

**Builtin procedure**:
A procedure implemented in Go and pre-bound in the global environment (e.g. `car`, `+`,
`cons`). Contrast with procedures defined in Scheme via `lambda`.
_Avoid_: primitive, native function, intrinsic

**Closure**:
A procedure created by `lambda`, pairing its parameter list and body with the
environment in which it was defined (giving lexical scope).
_Avoid_: lambda (the keyword) — use "closure" for the runtime value

**Environment**:
The mapping from symbols to values in effect during evaluation, as a chain of frames
(lexical scoping; inner frames shadow outer).
_Avoid_: scope, context, frame (a frame is one link of the environment chain)

**Exact / Inexact**:
The two numeric flavours. Exact numbers are int64 integers; inexact numbers are float64
reals. (See ADR-0003 for the numeric model.)

**Tail position** (WIP):
An expression position whose value is the value of the enclosing procedure call, so a
call there need not grow the stack. Definition firm; exact set of tail positions per
special form to be enumerated in the spec.

**REPL**:
The read-eval-print loop: read one datum, evaluate it, print the result, repeat. Also the
default interactive entry point of the program.
