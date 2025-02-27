package engine

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	var uid *uint64
	for range 100 {
		cfg := defaultConfig()

		if cfg.addr != "" {
			t.Fatalf("Expected Empty Addr. got=%q",
				cfg.addr)
		}

		if cfg.kind != "" {
			t.Fatalf("Expected Empty Kind. got=%q",
				cfg.kind)
		}

		if cfg.uid == nil {
			t.Fatalf("Expected Random UID. got nil")
		}

		if uid != nil && uid == cfg.uid {
			t.Logf("Expected unique ID. previous=%d, curr=%d",
				uid, &cfg.kind)
		}
		uid = cfg.uid
	}
}

func TestNewConfig(t *testing.T) {
	var uid *uint64
	tests := []string{
		"foo",
		"bar",
		"baz",
	}

	for _, kind := range tests {
		cfg := NewActorConfig(kind)

		if cfg.addr != "" {
			t.Fatalf("Expected Empty Addr. got=%q",
				cfg.addr)
		}

		if cfg.kind != kind {
			t.Fatalf("NewConfig Kind error. exp=%q, got=%q",
				kind, cfg.kind)
		}

		if cfg.uid == nil {
			t.Fatalf("Expected Random UID. got nil")
		}

		if uid != nil && uid == cfg.uid {
			t.Logf("Expected unique ID. previous=%d, curr=%d",
				uid, &cfg.kind)
		}
		uid = cfg.uid
	}
}

func TestNewConfigWithOpts(t *testing.T) {
	tests := []struct {
		expKind string
		expFrom string
		expAddr string
		expUid  uint64
	}{
		{"foo", "bar", "bar::foo", 100},
		{"bar", "baz", "baz::bar", 200},
		{"baz", "foo::bar", "foo::bar::baz", 300},
	}

	for _, tt := range tests {
		cfg := NewActorConfig(tt.expKind, fromAddr(tt.expFrom), withUid(tt.expUid))

		if cfg.ParentAddress() != tt.expFrom {
			t.Fatalf("Parent Address error. exp=%q got=%q",
				tt.expFrom, cfg.addr)
		}

		if cfg.Kind() != tt.expKind {
			t.Fatalf("Kind error. exp=%q, got=%q",
				tt.expKind, cfg.Kind())
		}

		if cfg.Address() != tt.expAddr {
			t.Fatalf("Address error. exp=%q got=%q",
				tt.expAddr, cfg.Address())
		}

		if cfg.UID() != tt.expUid {
			t.Fatalf("UID error. exp=%d got=%d",
				&tt.expUid, 10)
		}
	}
}
