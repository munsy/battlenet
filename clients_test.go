package battlenet

import (
	"testing"
)

func TestNewAccountClient(t *testing.T) {
	args := "test"
	c, err := NewAccountClient(args)

	if c == nil {
		t.Fatal("nil client")
	}
	if nil != err {
		t.Fatal(err.Error())
	}
}

func TestNewD3Client(t *testing.T) {
	args := "test"
	c, err := NewD3Client(args)

	if c == nil {
		t.Fatal("nil client")
	}
	if nil != err {
		t.Fatal(err.Error())
	}
}

func TestNewSC2Client(t *testing.T) {
	args := "test"
	c, err := NewSC2Client(args)

	if c == nil {
		t.Fatal("nil client")
	}
	if nil != err {
		t.Fatal(err.Error())
	}
}

func TestNewWoWClient(t *testing.T) {
	args := "test"
	c, err := NewWoWClient(args)

	if c == nil {
		t.Fatal("nil client")
	}
	if nil != err {
		t.Fatal(err.Error())
	}
}
