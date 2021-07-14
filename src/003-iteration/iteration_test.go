package iteration //nolint: testpackage

import "testing"

//nolint: ifshort
func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}
