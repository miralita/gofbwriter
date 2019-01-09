package gofbwriter

import (
	"testing"
	"time"
)

func testInterface(t *testing.T) {
	var v fb
	v = &annotation{}
	t.Log(v.tag())
	v = &author{}
	t.Log(v.tag())
	v = &binary{}
	t.Log(v.tag())
	v = &body{}
	t.Log(v.tag())
	v = &book{}
	t.Log(v.tag())
	v = &cite{}
	t.Log(v.tag())
	v = &customInfo{}
	t.Log(v.tag())
	tmp := date(time.Now())
	v = &tmp
	t.Log(v.tag())
	v = &description{}
	t.Log(v.tag())
	v = &documentInfo{}
	t.Log(v.tag())
	v = &epigraph{}
	t.Log(v.tag())
	v = &image{}
	t.Log(v.tag())
	v = &empty{}
	t.Log(v.tag())
	v = &p{}
	t.Log(v.tag())
	v = &partShareInstructionType{}
	t.Log(v.tag())
	v = &poem{}
	t.Log(v.tag())
	v = &publishInfo{}
	t.Log(v.tag())
	v = &section{}
	t.Log(v.tag())
	v = &sequence{}
	t.Log(v.tag())
	v = &shareInstructionType{}
	t.Log(v.tag())
	v = &stanza{}
	t.Log(v.tag())
	v = &stylesheet{}
	t.Log(v.tag())
	v = &table{}
	t.Log(v.tag())
	v = &title{}
	t.Log(v.tag())
	v = &titleInfo{}
	t.Log(v.tag())
}
