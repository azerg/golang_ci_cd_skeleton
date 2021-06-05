package main

import "testing"

func TestFoo(t *testing.T) {
	actual := foo(2)
	if actual != 4 {
		t.Error("eek")
	}
}
