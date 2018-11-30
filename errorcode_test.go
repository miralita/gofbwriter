package go_fbwriter

import "testing"

func TestCaller(t *testing.T) {
	err := makeError(ERR_EMPTY_FIRST_NAME, "test")
	t.Log(err)
}
