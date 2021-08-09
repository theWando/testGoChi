package main

import "testing"

func Test_foo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "want foobar", want: "foobar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foo(); got != tt.want {
				t.Errorf("foo() = %v, want %v", got, tt.want)
			}
		})
	}
}
