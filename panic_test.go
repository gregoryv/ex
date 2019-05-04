package ex

import (
	"fmt"
	"testing"
)

func Test_panics(t *testing.T) {
	err := fmt.Errorf("bad stuff")
	cases := []func(){
		func() { StatusOf(nil, err) },
		func() { JsonOf(nil, err) },
		func() { HeadersOf(nil, err) },
		func() { BodyOf(nil, err) },
	}
	for _, fn := range cases {
		shouldPanic(t, fn)
	}
}

func shouldPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		t.Helper()
		e := recover()
		if e == nil {
			t.Error("Expected a panic")
		}
	}()
	fn()
}
