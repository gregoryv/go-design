package design

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/go-design/shape"
)

func TestOneDiagram(t *testing.T) {
	it := NewOneDiagram(t)
	it.AdaptsInSize()
	it.CanHaveFixedSize()
}

func NewOneDiagram(t *testing.T) *OneDiagram {
	return &OneDiagram{T: t}
}

type OneDiagram struct {
	*testing.T
	Diagram
}

func (t *OneDiagram) AdaptsInSize() {
	t.Log("Adapts in size")
	t.Place(shape.NewLine(0, 0, 100, 100))
	w, h := t.AdaptSize()
	assert := asserter.New(t)
	assert(w == 100).Errorf("width did not adapt: %v", w)
	assert(h == 100).Errorf("height did not adapt: %v", h)
}

func (t *OneDiagram) CanHaveFixedSize() {
	t.Log("Can have fixed size")
	t.Place(shape.NewLine(0, 0, 100, 100))
	adjusted := &bytes.Buffer{}
	t.WriteSvg(adjusted)

	t.SetWidth(5)
	t.SetHeight(10)
	fixed := &bytes.Buffer{}
	t.WriteSvg(fixed)

	assert := asserter.New(t)
	assert(adjusted.String() != fixed.String())
}