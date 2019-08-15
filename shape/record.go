package shape

import (
	"fmt"
	"io"
)

type Record struct {
	X, Y         int
	Title        string
	PublicFields []string

	Font Font
	Pad  Padding
}

func (shape *Record) WriteSvg(w io.Writer) error {
	collect := &ErrCollector{}
	collect.Last(fmt.Fprintf(w,
		`<rect x="%v" y="%v" width="%v" height="%v"/>`,
		shape.X, shape.Y, shape.Width(), shape.Height()))

	collect.Err(shape.title().WriteSvg(w))
	return collect.First()
}

func (record *Record) title() *Label {
	return &Label{
		X:    record.X + record.Pad.Left,
		Y:    record.Y + record.Font.Height + record.Pad.Top,
		Text: record.Title,
	}
}

func (record *Record) Height() int {
	lines := 1
	return boxHeight(record.Font, record.Pad, lines) // todo
}

func (record *Record) Width() int {
	return boxWidth(record.Font, record.Pad, record.Title) // todo check widest
}

func (record *Record) Position() (int, int) { return record.X, record.Y }
func (record *Record) SetX(x int)           { record.X = x }
func (record *Record) SetY(y int)           { record.Y = y }