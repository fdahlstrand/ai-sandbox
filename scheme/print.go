package scheme

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// WriteString renders v in re-readable form — the Scheme write procedure. Strings
// are quoted with escapes and characters use #\ notation, so the output can be
// read back into an equal datum.
func WriteString(v Value) string {
	var b strings.Builder
	writeValue(&b, v, true)
	return b.String()
}

// DisplayString renders v in human-readable form — the Scheme display procedure.
// Strings and characters appear without quotes or escapes; everything else is the
// same as WriteString.
func DisplayString(v Value) string {
	var b strings.Builder
	writeValue(&b, v, false)
	return b.String()
}

// writeValue is the shared printer walker. The write flag selects the re-readable
// rendering (write) versus the human rendering (display); it differs only at
// strings and characters but is threaded through aggregates so nested atoms honour
// it too.
func writeValue(b *strings.Builder, v Value, write bool) {
	switch x := v.(type) {
	case bool:
		if x {
			b.WriteString("#t")
		} else {
			b.WriteString("#f")
		}
	case int64:
		b.WriteString(strconv.FormatInt(x, 10))
	case float64:
		b.WriteString(formatFloat(x))
	case Symbol:
		b.WriteString(string(x))
	case Char:
		if write {
			b.WriteString(writeChar(rune(x)))
		} else {
			b.WriteRune(rune(x))
		}
	case *String:
		if write {
			writeStringLiteral(b, x.Runes)
		} else {
			b.WriteString(string(x.Runes))
		}
	case emptyType:
		b.WriteString("()")
	case *Pair:
		writePair(b, x, write)
	case *Vector:
		b.WriteString("#(")
		for i, item := range x.Items {
			if i > 0 {
				b.WriteByte(' ')
			}
			writeValue(b, item, write)
		}
		b.WriteByte(')')
	default:
		// Unknown value type (e.g. a procedure added later, or a stray Go value):
		// fall back to Go formatting so the printer never panics.
		fmt.Fprintf(b, "%v", v)
	}
}

// writePair renders a pair, collapsing chains of pairs into list notation
// (a b c), ending dotted lists with " . tail)" and proper lists with ")".
func writePair(b *strings.Builder, p *Pair, write bool) {
	b.WriteByte('(')
	writeValue(b, p.Car, write)
	rest := p.Cdr
	for {
		switch x := rest.(type) {
		case *Pair:
			b.WriteByte(' ')
			writeValue(b, x.Car, write)
			rest = x.Cdr
		case emptyType:
			b.WriteByte(')')
			return
		default:
			b.WriteString(" . ")
			writeValue(b, x, write)
			b.WriteByte(')')
			return
		}
	}
}

// formatFloat renders an inexact number so it stays visibly inexact: the result
// always carries a '.', an exponent, or a named special so it reads back as a
// float rather than an integer.
func formatFloat(f float64) string {
	switch {
	case math.IsInf(f, 1):
		return "+inf.0"
	case math.IsInf(f, -1):
		return "-inf.0"
	case math.IsNaN(f):
		return "+nan.0"
	}
	s := strconv.FormatFloat(f, 'g', -1, 64)
	if !strings.ContainsAny(s, ".eE") {
		s += "."
	}
	return s
}

// writeChar renders a character in #\ notation, using the R7RS named forms for
// the common non-printing characters.
func writeChar(r rune) string {
	switch r {
	case ' ':
		return `#\space`
	case '\n':
		return `#\newline`
	case '\t':
		return `#\tab`
	case '\r':
		return `#\return`
	case 0:
		return `#\null`
	}
	return `#\` + string(r)
}

// writeStringLiteral writes runes as a double-quoted string literal with the
// standard escapes so the result is re-readable.
func writeStringLiteral(b *strings.Builder, runes []rune) {
	b.WriteByte('"')
	for _, r := range runes {
		switch r {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\n':
			b.WriteString(`\n`)
		case '\t':
			b.WriteString(`\t`)
		case '\r':
			b.WriteString(`\r`)
		default:
			b.WriteRune(r)
		}
	}
	b.WriteByte('"')
}
