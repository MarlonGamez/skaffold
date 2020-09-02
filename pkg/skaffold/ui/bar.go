package ui

import (
	"sync"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

var (
	// Current is the current progress group being rendered
	Current ProgressGroup
)

// ProgressGroup is a simple wrapper for the mpb.Progress type
type ProgressGroup struct {
	*mpb.Progress
}

func NewProgress(wg *sync.WaitGroup) ProgressGroup {
	Current = ProgressGroup{mpb.New(
		mpb.WithOutput(out),
		mpb.PopCompletedMode(),
		mpb.WithWaitGroup(wg),
	)}
	return Current
}

// AddNewSpinner adds a progress spinner to the calling ProgressGroup
func AddNewSpinner(prefix, name string) *mpb.Bar {
	return Current.AddBar(1,
		mpb.BarStyle("       "),
		mpb.PrependDecorators(
			decor.Name(prefix),
			decor.Spinner(mpb.DefaultSpinnerStyle),
			decor.Name(" "+name+"..."),
		),
		mpb.BarFillerOnComplete("done."),
	)
}
