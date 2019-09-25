package main

import "testing"

//TestHello 测试
func TestHello(t *testing.T) {

	//判断是否是正确的字符串
	assertCorrectMessage := func(t *testing.T, got, want string) {
		//测试辅助行数，这样做以后，当测试失败时所报告的行数将在函数调用中而不是辅助函数中
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("sayingj hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
