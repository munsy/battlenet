package battlenet

import (
	"testing"
)

// make these better

func TestNewAccountClient(t *testing.T) {
	args := "test"
	c, err := NewAccountClient(args)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

func TestNewD3Client(t *testing.T) {
	args := "test"
	c, err := NewD3Client(args)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

func TestNewSC2Client(t *testing.T) {
	args := "test"
	c, err := NewSC2Client(args)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

func TestNewWoWClient(t *testing.T) {
	args := "test"
	c, err := NewWoWClient(args)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}
