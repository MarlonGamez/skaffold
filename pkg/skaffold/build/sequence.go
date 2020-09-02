/*
Copyright 2019 The Skaffold Authors

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

package build

import (
	"context"
	"fmt"
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build/tag"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/event"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/ui"
)

// InSequence builds a list of artifacts in sequence.
func InSequence(ctx context.Context, out io.Writer, tags tag.ImageTags, artifacts []*latest.Artifact, buildArtifact ArtifactBuilder) ([]Artifact, error) {
	var builds []Artifact

	for _, artifact := range artifacts {
		bar := ui.AddNewSpinner("  ", artifact.ImageName)

		color.Default.Fprintf(out, "Building [%s]...\n", artifact.ImageName)

		event.BuildInProgress(artifact.ImageName)

		tag, present := tags[artifact.ImageName]
		if !present {
			return nil, fmt.Errorf("unable to find tag for image %s", artifact.ImageName)
		}

		finalTag, err := buildArtifact(ctx, out, artifact, tag)
		if err != nil {
			event.BuildFailed(artifact.ImageName, err)
			return nil, fmt.Errorf("couldn't build %q: %w", artifact.ImageName, err)
		}

		bar.Increment()
		event.BuildComplete(artifact.ImageName)

		builds = append(builds, Artifact{
			ImageName: artifact.ImageName,
			Tag:       finalTag,
		})
	}

	return builds, nil
}
