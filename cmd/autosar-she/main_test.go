package main

import (
	"testing"
)

func TestMainErr(t *testing.T) {
	if err := mainErr(); err == nil {
		t.Fatal("expected error")
	}
	if err := mainErr("foo"); err == nil {
		t.Fatal("expected error")
	}

}
