# Roadmap

The path ahead (and behind) for the learning Scheme interpreter (Go). High-level
backlog; each entry becomes a Work Plan when picked up. The prime directive is
pedagogical: features earn their place by what they teach about Scheme.

## Core (the chunk currently being scoped)
- [ ] Reader/parser: tokenizer + s-expression reader (the `read` half of the REPL)
- [ ] Tree-walking evaluator with explicit tail-call loop (see ADR-0002)
- [ ] REPL loop (read-eval-print) with a `(load "file.scm")` builtin (REPL-only entry point for now)
- [ ] Core special forms and a starter set of builtin procedures

## Deferred / future
- [ ] Fuller numeric tower: arbitrary-precision integers + rationals (see ADR-0003) (WIP)
- [ ] Macros / `syntax-rules` (hygienic macros) (WIP) — needs a grill-me on approach
- [ ] First-class continuations / `call/cc` (WIP) — hard in a tree-walker, needs grill-me
- [ ] In-Scheme condition system: `raise` / `guard` / `with-exception-handler` (see ADR-0006)
- [ ] `int64` overflow handling decision (wrap / error / promote) — subsumed if numeric tower lands (see ADR-0003) (WIP)
- [ ] CLI / batch mode: run a file and exit (`interp file.scm`), `-i` flag — beyond the REPL-only entry point (WIP)
