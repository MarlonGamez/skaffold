/*
Copyright 2021 The Skaffold Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v2

import (
	"fmt"
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/constants"
	proto "github.com/GoogleContainerTools/skaffold/proto/v2"
)

type Logger struct {
	phase constants.Phase
	subtaskId string
	origin string
}

func NewLogger(phase constants.Phase, subtaskId, origin string) io.Writer {
	return Logger {
		phase: phase,
		subtaskId: subtaskId,
		origin: origin,
	}
}

func (l Logger) Write(p []byte) (n int, err error) {
	handler.handleSkaffoldLogEvent(&proto.SkaffoldLogEvent{
		TaskId:               fmt.Sprintf("%s-%d", l.phase, handler.iteration),
		SubtaskId:            l.subtaskId,
		Origin:               l.origin,
		Level:                0,
		Message:              string(p),
	})

	return len(p), nil
}

func (ev *eventHandler) handleSkaffoldLogEvent(e *proto.SkaffoldLogEvent) {
	ev.handle(&proto.Event{
		EventType: &proto.Event_SkaffoldLogEvent{
			SkaffoldLogEvent: e,
		},
	})
}

type Writer struct {
	origWriter io.Writer
	subTaskID string
	origin string
	phase constants.Phase
}

func NewWriter(origWriter io.Writer, phase constants.Phase, subTaskID string, origin string) io.Writer {
	return Writer {
		origWriter: origWriter,
		phase: phase,
		subTaskID: subTaskID,
		origin: origin,
	}
}

func (w Writer) Write(p []byte) (int, error) {
	w.origWriter.Write(p)

	handler.handleSkaffoldLogEvent(&proto.SkaffoldLogEvent{
		TaskId:               fmt.Sprintf("%s-%d", w.phase, handler.iteration),
		SubtaskId:            w.subTaskID,
		Origin:               w.origin,
		Level:                0,
		Message:              string(p),
	})

	return 0, nil
}