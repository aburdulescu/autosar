package main

import (
	"testing"
)

func Test(t *testing.T) {
	if err := mainErr(); err == nil {
		t.Fatal("expected error")
	}
	if err := mainErr("foo"); err == nil {
		t.Fatal("expected error")
	}
	if err := mainErr("-foo"); err == nil {
		t.Fatal("expected error")
	}

	if err := mainErr("example"); err != nil {
		t.Fatal(err)
	}
	if err := mainErr("example", "-foo"); err == nil {
		t.Fatal("expected error")
	}

	if err := mainErr("encode"); err == nil {
		t.Fatal("expected error")
	}
	if err := mainErr("encode", "-foo"); err == nil {
		t.Fatal("expected error")
	}
}
