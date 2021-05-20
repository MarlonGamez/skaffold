package output

import (
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/constants"
	eventV2 "github.com/GoogleContainerTools/skaffold/pkg/skaffold/event/v2"
)

type SkaffoldWriter struct {
	MainWriter  io.Writer
	EventWriter io.Writer
}

func (s SkaffoldWriter) Write(p []byte) (int, error) {
	for _, w := range []io.Writer{s.MainWriter, s.EventWriter} {
		n, err := w.Write(p)
		if err != nil {
			return n, err
		}
		if n < len(p) {
			return n, io.ErrShortWrite
		}
	}

	return len(p), nil
}

func SetupOutput(out io.Writer, defaultColor int, forceColors bool) io.Writer {
	return SkaffoldWriter{
		MainWriter:  SetupColors(out, defaultColor, forceColors),
		EventWriter: io.Discard,
	}
}

func EventContext(out io.Writer, phase constants.Phase, subtaskId, origin string) io.Writer {
	sw, ok := out.(SkaffoldWriter)
	if !ok {
		return out
	}

	return SkaffoldWriter{
		MainWriter: sw.MainWriter,
		EventWriter: eventV2.NewLogger(phase, subtaskId, origin),
	}
}
