package main

import (
	
	"testing"
)


func Test1(t *testing.T){
	str := []string{"cd", "../dev01"}
	c := CdCommand(str)
	var want error
	want = nil
	if  c != want {
		t.Errorf("got %q, wanted %q", c, want)
	}
}
func Test2(t *testing.T){
	str := []string{"pwd"}
	c , _ := PwdCommand(str)
	want := "/Users/macbook/Golang/internship_go/WB_Tech-L2/develop/dev08" 
	if  c != want {
		t.Errorf("got %q, wanted %q", c, want)
	}
}

func Test3(t *testing.T){
	str := []string{"echo", "Hello,world!"}
	c , _ := EchoCommand(str)
	want := "Hello,world! " 
	if  c != want {
		t.Errorf("got %q, wanted %q", c, want)
	}
}