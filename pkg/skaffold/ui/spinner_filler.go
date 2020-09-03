package ui

import (
	"io"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

type skaffoldFiller struct {
	frames []string
	count  uint
}

// NewSkaffoldFiller constructs mpb.BarFiller, to be used with *Progress.Add(...) *Bar method.
func NewSkaffoldFiller(style []string) mpb.BarFiller {
	if len(style) == 0 {
		style = mpb.DefaultSpinnerStyle
	}
	filler := &skaffoldFiller{
		frames: style,
	}
	return filler
}

func (s *skaffoldFiller) Fill(w io.Writer, reqWidth int, stat decor.Statistics) {
	frame := s.frames[s.count%uint(len(s.frames))]
	// frameWidth := utf8.RuneCountInString(frame)

	io.WriteString(w, frame+" ")
	s.count++
}
