package gofbwriter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInterface(t *testing.T) {
	var v Fb
	v = &Annotation{}
	assert.Equal(t, "annotation", v.tag(), "Should be annotation")
	t.Log(v.tag())
	v = &Author{}
	t.Log(v.tag())
	v = &Binary{}
	t.Log(v.tag())
	v = &Body{}
	t.Log(v.tag())
	v = &Fb2{}
	t.Log(v.tag())
	v = &Cite{}
	t.Log(v.tag())
	v = &CustomInfo{}
	t.Log(v.tag())
	tmp := Date(time.Now())
	v = &tmp
	t.Log(v.tag())
	v = &Description{}
	t.Log(v.tag())
	v = &DocumentInfo{}
	t.Log(v.tag())
	v = &Epigraph{}
	t.Log(v.tag())
	v = &Image{}
	t.Log(v.tag())
	v = &empty{}
	t.Log(v.tag())
	v = &p{}
	t.Log(v.tag())
	v = &partShareInstructionType{}
	t.Log(v.tag())
	v = &Poem{}
	t.Log(v.tag())
	v = &PublishInfo{}
	t.Log(v.tag())
	v = &Section{}
	t.Log(v.tag())
	v = &Sequence{}
	t.Log(v.tag())
	v = &ShareInstruction{}
	t.Log(v.tag())
	v = &Stanza{}
	t.Log(v.tag())
	v = &StyleSheet{}
	t.Log(v.tag())
	v = &Table{}
	t.Log(v.tag())
	v = &Title{}
	t.Log(v.tag())
	v = &TitleInfo{}
	t.Log(v.tag())
}
