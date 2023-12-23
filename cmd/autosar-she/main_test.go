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
	if err := mainErr("-foo"); err == nil {
		t.Fatal("expected error")
	}
	if err := mainErr("example"); err != nil {
		t.Fatal(err)
	}
	if err := mainErr("encode"); err == nil {
		t.Fatal("expected error")
	}
	if err := mainErr("decode"); err == nil {
		t.Fatal("expected error")
	}
	if err := mainErr("-version"); err != nil {
		t.Fatal(err)
	}
}
