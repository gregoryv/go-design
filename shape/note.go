package shape

import (
	"fmt"
	"io"

	"github.com/gregoryv/go-design/xy"
)

func NewNote(text string) *Note {
	return &Note{
		Text:  text,
		Font:  DefaultFont,
		Pad:   DefaultPad,
		class: "note",
	}
}

type Note struct {
	Pos  xy.Position
	Text string

	Font
	Pad   Padding
	class string
}

func (n *Note) String() string {
	return fmt.Sprintf("Note at %v", n.Pos)
}
func (note *Note) Position() (int, int) { return note.Pos.XY() }
func (note *Note) Direction() Direction { return LR }
func (note *Note) SetX(x int)           { note.Pos.X = x }
func (note *Note) SetY(y int)           { note.Pos.Y = y }

func (n *Note) Width() int { return n.TextWidth(n.Text) }
func (n *Note) Height() int {
	lines := 1 // todo
	return boxHeight(n.Font, n.Pad, lines)
}
func (n *Note) SetClass(c string) { n.class = c }

func (n *Note) WriteSvg(out io.Writer) error {
	x, y := n.Pos.XY()
	x2 := x + n.Width()
	y2 := y + n.Height()
	flap := 10
	xf := x2 - flap
	yf := y + flap
	w, err := newTagPrinter(out)
	w.printf(`<path class="%s-box" d="M%v,%v `, n.class, x, y)
	w.printf(`L%v,%v %v,%v %v,%v %v,%v %v,%v %v,%v %v,%v M%v,%v L%v,%v" />`,
		x, y, x, y2, x2, y2, x2, yf, xf, y, xf, yf, x2, yf, xf, y, x, y)
	w.printf(`<text class="note" x="%v" y="%v">%s</text>`,
		x, y+n.Font.LineHeight, n.Text)
	return *err
}
