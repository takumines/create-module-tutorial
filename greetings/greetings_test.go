package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName Hello関数から返ってくるメッセージに指定した名前が含まれていること
func TestHelloName(t *testing.T) {
	name := "takumines"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloNames Hellos関数から返ってくるメッセージに指定した名前が含まれていること
func TestHelloNames(t *testing.T) {
	names := []string{
		"takumines",
		"taro",
	}
	want := regexp.MustCompile(`\b` + names[0] + `\b|\b` + names[1] + `\b`)
	msgs, err := Hellos(names)
	if !want.MatchString(msgs["takumines"]) || !want.MatchString(msgs["taro"]) || err != nil {
		t.Fatalf(`Hellos(%v) = %q, %v, want match for %#q, nil`, names, msgs, err, want)
	}
}

// TestHelloEmptyName 名前が空文字の場合はエラーが発生すること
func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
