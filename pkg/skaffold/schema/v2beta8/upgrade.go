/*
Copyright 2020 The Skaffold Authors

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

package v2beta8

import (
	next "github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/util"
	pkgutil "github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
)

// Upgrade upgrades a configuration to the next version.
// Config changes from v2beta8 to v2beta9
// 1. Additions:
//    - steps field to deploy config, where users can specify a list of DeployTypes to run in sequence
// 2. Removals
//    - removed deploy.DeployType in favor of deploy.steps for specifying deployers
// 3. Updates:
//    - sync.auto becomes boolean
func (c *SkaffoldConfig) Upgrade() (util.VersionedConfig, error) {
	var newConfig next.SkaffoldConfig
	pkgutil.CloneThroughJSON(c, &newConfig)
	newConfig.APIVersion = next.Version

	err := util.UpgradePipelines(c, &newConfig, upgradeOnePipeline)
	return &newConfig, err
}

func upgradeOnePipeline(oldPipeline, newPipeline interface{}) error {
	oldDeploy := &oldPipeline.(*Pipeline).Deploy
	newDeploy := &newPipeline.(*next.Pipeline).Deploy
	if oldDeploy.DeployType != (DeployType{}) {
		newDeployStep := next.DeployType{}
		pkgutil.CloneThroughJSON(&oldDeploy.DeployType, &newDeployStep)

		newDeploy.Steps = []next.DeployType{newDeployStep}
	}

	return nil
}

func (a *Auto) MarshalJSON() ([]byte, error) {
	// The presence of an Auto{} means auto-sync is enabled.
	if a != nil {
		return []byte(`true`), nil
	}
	return nil, nil
}
