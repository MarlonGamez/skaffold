package ui

import (
	"sync"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

var (
	current ProgressGroup
)

// ProgressGroup is a simple wrapper for the mpb.Progress type
type ProgressGroup struct {
	*mpb.Progress
}

func NewProgressGroup(wg *sync.WaitGroup) ProgressGroup {
	current = ProgressGroup{mpb.New(
		mpb.WithOutput(out),
		mpb.WithWaitGroup(wg),
	)}
	return current
}

// AddNewSpinner adds a progress spinner to the calling ProgressGroup
func AddNewSpinner(prefix, name string) *mpb.Bar {
	return current.Add(1, NewSkaffoldFiller(mpb.DefaultSpinnerStyle),
		mpb.PrependDecorators(
			decor.Name(prefix),
		),
		mpb.AppendDecorators(
			decor.Name(name+"..."),
		),
		mpb.BarFillerOnComplete("✓"),
	)
}
