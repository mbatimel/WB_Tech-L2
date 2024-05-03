package main

import "testing"

func TestSimple(t *testing.T) {
	got, _ := UnpackingString("a4bc2d5e3")
	want := "aaaabccdddddeee"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSimple2(t *testing.T) {
	got, _ := UnpackingString("abcd")
	want := "abcd"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSimple3(t *testing.T) {
	got, _ := UnpackingString("")
	want := ""
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestEscapeSequences(t *testing.T) {
	got, _ := UnpackingString(`qwe\45`)
	want := `qwe44444`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestEscapeSequences2(t *testing.T) {
	got, _ := UnpackingString(`qwe\4\5`)
	want := `qwe45`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestEscapeSequences3(t *testing.T) {
	got, _ := UnpackingString(`qwe\\5`)
	want := `qwe\\\\\`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestError(t *testing.T) {
	_, err := UnpackingString("45aaa")
	if err == nil {
		t.Errorf("Expect error but it isn't")
	}
}
func TestError2(t *testing.T) {
	_, err := UnpackingString("45")
	if err == nil {
		t.Errorf("Expect error but it isn't")
	}
}