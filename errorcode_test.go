package gofbwriter

import "testing"

func TestCaller(t *testing.T) {
	err := makeError(ErrEmptyFirstName, "test")
	t.Log(err)
}
