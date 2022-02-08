package controller

import (
	"testing"
)

func TestEvalFunctionPositiveCase(t *testing.T) {
	ok, err := EvalFunction("something is online", []string{"offline"})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	if !ok {
		t.Error("got false from eval function")
		t.Fail()
	}
}

func TestEvalFunctionNegativeCase(t *testing.T) {
	ok, err := EvalFunction("something is online", []string{"online"})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	if ok {
		t.Error("this value should be false")
		t.Fail()
	}
}
