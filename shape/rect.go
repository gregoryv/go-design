package shape

import (
	"fmt"
	"io"
	"math"

	"github.com/gregoryv/go-design/xy"
)

func NewRect(title string) *Rect {
	return &Rect{
		Title: title,
		Font:  DefaultFont,
		Pad:   DefaultTextPad,
		class: "rect",
	}
}

type Rect struct {
	X, Y  int
	Title string

	Font  Font
	Pad   Padding
	class string
}

func (r *Rect) String() string {
	return fmt.Sprintf("R %q", r.Title)
}

func (r *Rect) Position() (int, int) { return r.X, r.Y }
func (r *Rect) SetX(x int)           { r.X = x }
func (r *Rect) SetY(y int)           { r.Y = y }
func (r *Rect) Direction() Direction { return LR }
func (r *Rect) SetClass(c string)    { r.class = c }

func (r *Rect) WriteSvg(out io.Writer) error {
	w, err := newTagPrinter(out)
	w.printf(
		`<rect class="%s" x="%v" y="%v" width="%v" height="%v"/>`,
		r.class, r.X, r.Y, r.Width(), r.Height())
	w.printf("\n")
	r.title().WriteSvg(w)
	return *err
}

func (r *Rect) title() *Label {
	return &Label{
		Pos: xy.Position{
			r.X + r.Pad.Left,
			r.Y + r.Pad.Top/2,
		},
		Font:  r.Font,
		Text:  r.Title,
		class: "record-title",
	}
}

func (r *Rect) SetFont(f Font)         { r.Font = f }
func (r *Rect) SetTextPad(pad Padding) { r.Pad = pad }

func (r *Rect) Height() int {
	return boxHeight(r.Font, r.Pad, 1)
}

func (r *Rect) Width() int {
	return boxWidth(r.Font, r.Pad, r.Title)
}

// Edge returns xy position of a line starting at start and
// pointing to the records center. The returned position would
// be at the crosspoint.
func (r *Rect) Edge(start xy.Position) xy.Position {
	center := xy.Position{
		r.X + r.Width()/2,
		r.Y + r.Height()/2,
	}
	l1 := xy.Line{start, center}

	var (
		d      float64 = math.MaxFloat64
		pos    xy.Position
		lowY   = r.Y + r.Height()
		rightX = r.X + r.Width()
		top    = xy.NewLine(r.X, r.Y, rightX, r.Y)
		left   = xy.NewLine(r.X, r.Y, r.X, lowY)
		right  = xy.NewLine(rightX, r.Y, rightX, lowY)
		bottom = xy.NewLine(r.X, lowY, rightX, lowY)
	)

	for _, side := range []*xy.Line{top, left, right, bottom} {
		p, err := l1.IntersectSegment(side)
		if err != nil {
			continue
		}
		dist := start.Distance(p)
		if dist < d {
			pos = p
			d = dist
		}
	}
	return pos
}
