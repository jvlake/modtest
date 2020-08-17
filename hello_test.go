package modtest

import "testing"

func TestFoo(t *testing.T){
	t.Parallel()

	h:= NewHello()
	s := h.ToJsonString()
	p:=FromJsonString(s)
	if p.Name != h.Name{
		t.Fatalf("failed: %v", p.Name)
	}
}
