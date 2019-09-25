package main

import "testing"

//TestHello 测试
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"
	if got != want {
		t.Errorf("got '%s want '%s", got, want)
	}
}
