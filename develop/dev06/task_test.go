package main
import "testing"


func Test1(t *testing.T){
	c := NewCustomCut("1,2", " ", true)
	got:= c.CustomCutFunc("asd asd asd asd asd asd asd")
	want := "asd asd "
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func Test2(t *testing.T){
	c := NewCustomCut("1,2", " ", true)
	got:= c.CustomCutFunc("asd_asd_asd_asd_asd_asd_asd")
	want := ""
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func Test3(t *testing.T){
	c := NewCustomCut("2", "\t", true)
	got:= c.CustomCutFunc("a	d	f	c")
	want := "d "
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}