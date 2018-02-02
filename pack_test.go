package mypack

import "testing"

func Testips(t *testing.T) {
	var v bool
	v = ips("112.256.123.124")
	if v {
		t.Error("Wrong")
	}
}

func Testclean(t *testing.T) {
	var v string
	v = clean("A@@#?aniket")
	if v != "Aaniket" {
		t.Error("Incorrect cleaning")
	}
}
