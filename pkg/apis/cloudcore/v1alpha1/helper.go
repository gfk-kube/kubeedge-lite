/*
Copyright 2019 The KubeEdge Authors.

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

package v1alpha1

import (
	"io/ioutil"

	"k8s.io/klog"
	"sigs.k8s.io/yaml"
)

func (c *CloudCoreConfig) Parse(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		klog.Errorf("Failed to read configfile %s: %v", filename, err)
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		klog.Errorf("Failed to unmarshal configfile %s: %v", filename, err)
		return err
	}
	return nil
}
