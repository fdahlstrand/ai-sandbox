// Package scheme implements a small learning-oriented Scheme interpreter: the
// value model, reader, evaluator, and builtins. This file defines the Go
// representation of every Scheme datum and the helpers for building lists.
package scheme

import "sync"

// Value is the universal type of every Scheme datum. The concrete dynamic types
// stored in a Value are:
//
//	int64      exact integer        (ADR-0003)
//	float64    inexact real         (ADR-0003)
//	bool       #t / #f
//	Symbol     interned identifier
//	Char       character
//	*String    mutable string
//	*Pair      cons cell
//	emptyType  the empty list '()   (the Empty singleton)
//	*Vector    vector
//
// Procedure types (*Lambda, *Builtin) are added by a later work package, which
// also extends the printer to render them.
type Value any

// Symbol is an interned identifier datum such as lambda, x, or +. Construct one
// with Intern so equal names share a single canonical value.
type Symbol string

// Char is a single Scheme character.
type Char rune

// String is a Scheme string. It is backed by a []rune slice so individual
// characters can be mutated later (string-set!), and is deliberately distinct
// from Symbol.
type String struct {
	Runes []rune
}

// NewString builds a String holding the runes of s.
func NewString(s string) *String {
	return &String{Runes: []rune(s)}
}

// Text returns the string's contents as a Go string.
func (s *String) Text() string {
	return string(s.Runes)
}

// Pair is a two-field cons cell. Chains of pairs ending in Empty form lists.
// Pairs are always handled by pointer (*Pair).
type Pair struct {
	Car, Cdr Value
}

// emptyType is the type of the empty list. It has a single instance, Empty.
type emptyType struct{}

// Empty is the empty list '(). It is a distinct singleton — '() is not the same
// as a nil *Pair — so list walkers test against it explicitly.
var Empty = emptyType{}

// Vector is a Scheme vector. Vectors are always handled by pointer (*Vector).
type Vector struct {
	Items []Value
}

// internTable canonicalises symbol names so equal names map to one Symbol value.
var internTable sync.Map // map[string]Symbol

// Intern returns the canonical Symbol for name, interning it on first use. Going
// through Intern keeps symbol identity cheap for eq? in a later package.
func Intern(name string) Symbol {
	if v, ok := internTable.Load(name); ok {
		return v.(Symbol)
	}
	actual, _ := internTable.LoadOrStore(name, Symbol(name))
	return actual.(Symbol)
}

// Cons builds a new pair (car . cdr).
func Cons(car, cdr Value) *Pair {
	return &Pair{Car: car, Cdr: cdr}
}

// List builds a proper list from items, ending in Empty.
func List(items ...Value) Value {
	var result Value = Empty
	for i := len(items) - 1; i >= 0; i-- {
		result = Cons(items[i], result)
	}
	return result
}

// ListToSlice walks v as a list and returns its elements. The second result is
// true when v is a proper list (a chain of pairs ending in Empty) and false for
// an improper list, in which case the returned slice holds the elements gathered
// before the non-pair tail.
func ListToSlice(v Value) ([]Value, bool) {
	var items []Value
	for {
		switch x := v.(type) {
		case emptyType:
			return items, true
		case *Pair:
			items = append(items, x.Car)
			v = x.Cdr
		default:
			return items, false
		}
	}
}
