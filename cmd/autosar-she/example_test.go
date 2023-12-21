package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	if err := cmdExample("-foo"); err == nil {
		t.Fatal("expected error")
	}
}
