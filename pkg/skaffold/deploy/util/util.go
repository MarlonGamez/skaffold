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

package util

import (
	"fmt"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/config"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/docker"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
)

// ApplyDefaultRepo applies the default repo to a given image tag.
func ApplyDefaultRepo(globalConfig string, defaultRepo *string, tag string) (string, error) {
	repo, err := config.GetDefaultRepo(globalConfig, defaultRepo)
	if err != nil {
		return "", fmt.Errorf("getting default repo: %w", err)
	}

	newTag, err := docker.SubstituteDefaultRepoIntoImage(repo, tag)
	if err != nil {
		return "", fmt.Errorf("applying default repo to %q: %w", tag, err)
	}

	return newTag, nil
}

// ListDeployers returns a list of deployer names being used in the given deploy config.
func ListDeployers(deploy *latest.DeployType) []string {
	results := NewStringSet()
	if deploy.HelmDeploy != nil {
		results.Insert("helm")
	}
	if deploy.KptDeploy != nil {
		results.Insert("kpt")
	}
	if deploy.KubectlDeploy != nil {
		results.Insert("kubectl")
	}
	if deploy.KustomizeDeploy != nil {
		results.Insert("kustomize")
	}

	return results.ToList()
}
