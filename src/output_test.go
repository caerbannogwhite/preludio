package preludiocore

import (
	"strings"
	"testing"

	"github.com/caerbannogwhite/enchanter/dataframe"
)

// Every kind of result comes back as a dataframe and renders through
// enchanter's Table. The one-row string case used to panic in enchanter's
// string formatter before v0.5.4, so this walks the REPL's exact path.
func Test_Output_TableRendering(t *testing.T) {
	params := dataframe.NewPPrintParams()

	// A single string value: a one-row, one-column frame.
	res := be.RunSource("tblhello := 'hello'")
	if len(res.Data) != 1 {
		t.Fatalf("expected one result, got %d", len(res.Data))
	}
	out := res.Data[0].Table(params)
	for _, want := range []string{"String", "hello", "1 row, 1 column"} {
		if !strings.Contains(out, want) {
			t.Errorf("table must contain %q, got:\n%s", want, out)
		}
	}

	// A scalar expression renders its value.
	res = be.RunSource("tblcalc := 1 + 2")
	if len(res.Data) != 1 {
		t.Fatalf("expected one result, got %d", len(res.Data))
	}
	if out := res.Data[0].Table(params); !strings.Contains(out, "3") {
		t.Errorf("table must contain the value 3, got:\n%s", out)
	}

	// A null in a string column keeps every box line the same width.
	res = be.RunSource("tblnulls := ['abcdef', na, 'ghijkl']")
	if len(res.Data) != 1 {
		t.Fatalf("expected one result, got %d", len(res.Data))
	}
	out = res.Data[0].Table(params)
	if !strings.Contains(out, "Na") {
		t.Errorf("table must show the null, got:\n%s", out)
	}
	want := -1
	for _, line := range strings.Split(strings.TrimRight(out, "\n"), "\n") {
		if strings.HasPrefix(line, "╭") {
			want = len([]rune(line))
		}
		if want > 0 && strings.HasPrefix(line, "│") && len([]rune(line)) != want {
			t.Errorf("misaligned row (%d runes, want %d): %q", len([]rune(line)), want, line)
		}
	}

	// A dataframe result renders as itself.
	res = be.RunSource("tblframe := new! [a = [1, 2], b = ['x', 'y']]")
	if len(res.Data) != 1 {
		t.Fatalf("expected one result, got %d", len(res.Data))
	}
	out = res.Data[0].Table(params)
	for _, want := range []string{"2 rows, 2 columns", "a", "b", "Int64", "String", "x"} {
		if !strings.Contains(out, want) {
			t.Errorf("table must contain %q, got:\n%s", want, out)
		}
	}

	// An empty frame says so instead of rendering a box.
	res = be.RunSource("tblempty := (from! tblframe | filter! a > 99)")
	if len(res.Data) != 1 {
		t.Fatalf("expected one result, got %d", len(res.Data))
	}
	if out := res.Data[0].Table(params); !strings.Contains(out, "Empty DataFrame") {
		t.Errorf("an empty frame must say so, got %q", out)
	}
}
