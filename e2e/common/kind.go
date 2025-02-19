// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"bytes"
	"os/exec"

	"github.com/kube-logging/logging-operator/e2e/common/kind"
)

const KindClusterCreationTimeout = "3m"

func KindClusterKubeconfig(name string) ([]byte, error) {
	err := kind.CreateCluster(kind.CreateClusterOptions{
		Name: name,
		Wait: KindClusterCreationTimeout,
	})
	if err != nil && !isClusterAlreadyExistsError(err) {
		return nil, err
	}
	return kind.GetKubeconfig(kind.GetKubeconfigOptions{
		Name: name,
	})
}

func isClusterAlreadyExistsError(err error) bool {
	exitErr, _ := err.(*exec.ExitError)
	return exitErr != nil && bytes.Contains(exitErr.Stderr, []byte("failed to create cluster: node(s) already exist for a cluster with the name"))
}
