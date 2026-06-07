package scheme

import "testing"

// vec is a small helper for building a *Vector in test tables.
func vec(items ...Value) *Vector { return &Vector{Items: items} }

func TestWriteString(t *testing.T) {
	cases := []struct {
		name string
		in   Value
		want string
	}{
		{"int", int64(42), "42"},
		{"negative int", int64(-7), "-7"},
		{"float", float64(1.5), "1.5"},
		{"float whole stays inexact", float64(1), "1."},
		{"true", true, "#t"},
		{"false", false, "#f"},
		{"symbol", Intern("lambda"), "lambda"},
		{"char letter", Char('a'), `#\a`},
		{"char space", Char(' '), `#\space`},
		{"char newline", Char('\n'), `#\newline`},
		{"string escaped", NewString("hi\n"), `"hi\n"`},
		{"string quotes and backslash", NewString(`a"b\c`), `"a\"b\\c"`},
		{"empty list", Empty, "()"},
		{"proper list", List(int64(1), int64(2), int64(3)), "(1 2 3)"},
		{"nested list", List(int64(1), List(int64(2), int64(3))), "(1 (2 3))"},
		{"dotted pair", Cons(int64(1), int64(2)), "(1 . 2)"},
		{"improper tail", Cons(int64(1), Cons(int64(2), int64(3))), "(1 2 . 3)"},
		{"vector", vec(int64(1), int64(2)), "#(1 2)"},
		{
			"acceptance list",
			List(int64(1), Char('a'), NewString("hi\n"), vec(int64(1), int64(2))),
			`(1 #\a "hi\n" #(1 2))`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := WriteString(c.in); got != c.want {
				t.Errorf("WriteString = %q, want %q", got, c.want)
			}
		})
	}
}

func TestDisplayString(t *testing.T) {
	cases := []struct {
		name string
		in   Value
		want string
	}{
		{"char unquoted", Char('a'), "a"},
		{"string unquoted with real newline", NewString("hi\n"), "hi\n"},
		{"int unchanged", int64(42), "42"},
		{"empty list", Empty, "()"},
		{
			"acceptance list",
			List(int64(1), Char('a'), NewString("hi\n"), vec(int64(1), int64(2))),
			"(1 a hi\n #(1 2))",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := DisplayString(c.in); got != c.want {
				t.Errorf("DisplayString = %q, want %q", got, c.want)
			}
		})
	}
}

func TestIntern(t *testing.T) {
	a := Intern("foo")
	b := Intern("foo")
	if a != b {
		t.Errorf("Intern returned different values for the same name: %q vs %q", a, b)
	}
	if a == Intern("bar") {
		t.Error("Intern returned equal values for different names")
	}
}

func TestListToSlice(t *testing.T) {
	items, proper := ListToSlice(List(int64(1), int64(2), int64(3)))
	if !proper {
		t.Error("proper list reported as improper")
	}
	if len(items) != 3 {
		t.Fatalf("got %d items, want 3", len(items))
	}

	if _, proper := ListToSlice(Cons(int64(1), int64(2))); proper {
		t.Error("improper list reported as proper")
	}

	if items, proper := ListToSlice(Empty); !proper || len(items) != 0 {
		t.Errorf("empty list: items=%v proper=%v, want [] true", items, proper)
	}
}
