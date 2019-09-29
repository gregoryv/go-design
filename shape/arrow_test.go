package shape

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/go-design/style"
)

func TestOneArrow(t *testing.T) {
	it := &OneArrow{
		T:      t,
		assert: asserter.New(t),
		Arrow:  NewArrow(50, 50, 50, 50),
	}
	it.CanPointUpAndRight()
	it.CanPointUpAndLeft()
	it.CanPointDownAndLeft()
	it.CanPointDownAndRight()
	// also
	it.CanPointRight()
	it.CanPointLeft()
	it.CanPointDown()
	it.CanPointUp()
	// when
	it.HasATail()
	it.HasBothTailAndHead()

	it.CanHaveASpecificClass()
	it.CanMove()
	it.IsVisible()
}

type OneArrow struct {
	*testing.T
	assert
	*Arrow
}

func (t *OneArrow) CanPointUpAndRight() {
	t.End.X = t.Start.X + 30
	t.End.Y = t.Start.Y - 30
	t.saveAs("testdata/arrow_points_up_and_right.svg")
}

func (t *OneArrow) CanPointUpAndLeft() {
	t.End.X = t.Start.X - 30
	t.End.Y = t.Start.Y - 30
	t.saveAs("testdata/arrow_points_up_and_left.svg")
}

func (t *OneArrow) CanPointDownAndLeft() {
	t.End.X = t.Start.X - 10
	t.End.Y = t.Start.Y + 30
	dir := t.Direction()
	t.assert(dir == RL).Errorf("Direction not right to left: %v", dir)
	t.saveAs("testdata/arrow_points_down_and_left.svg")
}

func (t *OneArrow) CanPointDownAndRight() {
	t.End.X = t.Start.X + 20
	t.End.Y = t.Start.Y + 30
	dir := t.Direction()
	t.assert(dir == LR).Errorf("Direction not left to right: %v", dir)
	t.saveAs("testdata/arrow_points_down_and_right.svg")
}

func (t *OneArrow) CanPointRight() {
	t.End.X = t.Start.X + 50
	t.End.Y = t.Start.Y
	dir := t.Direction()
	t.assert(dir == LR).Errorf("Direction not left to right: %v", dir)
	t.saveAs("testdata/arrow_points_right.svg")
}

func (t *OneArrow) CanPointLeft() {
	t.End.X = t.Start.X - 40
	t.End.Y = t.Start.Y
	dir := t.Direction()
	t.assert(dir == RL).Errorf("Direction not right to left: %v", dir)
	t.saveAs("testdata/arrow_points_left.svg")
}

func (t *OneArrow) CanPointDown() {
	t.End.X = t.Start.X
	t.End.Y = t.Start.Y + 50
	t.saveAs("testdata/arrow_points_down.svg")
}

func (t *OneArrow) CanPointUp() {
	t.End.X = t.Start.X
	t.End.Y = t.Start.Y - 40
	t.saveAs("testdata/arrow_points_up.svg")
}

func (t *OneArrow) HasATail() {
	t.Tail = true
}

func (t *OneArrow) HasBothTailAndHead() {
	t.saveAs("testdata/arrow_with_tail_and_head.svg")
}

func (t *OneArrow) saveAs(filename string) {
	t.Helper()
	d := &Svg{Width: 100, Height: 100}
	d.Append(t.Arrow)

	fh, err := os.Create(filename)
	if err != nil {
		t.Error(err)
		return
	}
	d.WriteSvg(style.NewStyler(fh))
	fh.Close()
}

func (t *OneArrow) CanHaveASpecificClass() {
	t.Helper()
	t.Class = "special"
	buf := &bytes.Buffer{}
	t.WriteSvg(buf)
	t.assert().Contains(buf.String(), "special")
}

func (t *OneArrow) CanMove() {
	t.Helper()
	x, y := t.Position()
	t.SetX(x + 1)
	t.SetY(y + 1)
	t.assert(x != t.Start.X).Error("start X still the same")
	t.assert(y != t.Start.Y).Error("start Y still the same")
}

func (t *OneArrow) IsVisible() {
	t.Helper()
	h := t.Height()
	w := t.Width()
	t.assert(h > 0 || w > 0).Errorf("%v not visible", t.Arrow)
}

func TestArrowBetweenShapes(t *testing.T) {
	it := &ArrowBetweenShapes{
		T:      t,
		assert: asserter.New(t),
	}

	it.StartsAndEndsAtEdgeOfShapes()
}

type ArrowBetweenShapes struct {
	*testing.T
	assert
}

type A struct{}

type B struct{}

func (t *ArrowBetweenShapes) StartsAndEndsAtEdgeOfShapes() {
	a := NewStructRecord(A{})
	a.SetX(10)
	a.SetY(100)
	b := NewStructRecord(B{})
	b.SetX(80)
	b.SetY(40)

	arrow := NewArrowBetween(a, b)
	svg := newSvg(200, 200, arrow, a, b)
	label := NewLabel(fmt.Sprintf("Angle: %v", arrow.absAngle()))
	label.SetX(100)
	label.SetY(80)

	svg.Append(label)
	writeSvgTo(t.T, "testdata/arrow_between_shapes.svg", svg)
}

func newSvg(width, height int, shapes ...SvgWriterShape) *Svg {
	svg := &Svg{Width: width, Height: height}
	svg.Append(shapes...)
	return svg
}

func writeSvgTo(t *testing.T, filename string, svg *Svg) {
	t.Helper()
	fh, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	svg.WriteSvg(style.NewStyler(fh))
	fh.Close()
}
