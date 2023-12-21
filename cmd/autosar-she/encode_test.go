package main

import (
	"testing"
)

func TestEncode(t *testing.T) {
	if err := cmdEncode("-foo"); err == nil {
		t.Fatal("expected error")
	}
}
