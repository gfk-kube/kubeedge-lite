/*
Copyright 2021 The KubeEdge Authors.

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

package config

import (
	"reflect"
	"testing"

	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/cloudcore/v1alpha1"
)

func TestInitConfigure(t *testing.T) {
	type args struct {
		ec            *v1alpha1.EdgeController
		kubeAPIConfig *v1alpha1.KubeAPIConfig
		nodeName      string
		edgesite      bool
	}

	ec := &v1alpha1.EdgeController{}
	kac := &v1alpha1.KubeAPIConfig{}
	nodeName := "NodeA"

	tests := []struct {
		name string
		args args
	}{
		{
			"TestInitCnofigure() Caes 1: init configurae",
			args{
				ec:            ec,
				kubeAPIConfig: kac,
				nodeName:      nodeName,
				edgesite:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConfigure(tt.args.ec, tt.args.kubeAPIConfig, tt.args.nodeName, tt.args.edgesite)
			if !reflect.DeepEqual(*tt.args.ec, Config.EdgeController) || !reflect.DeepEqual(*tt.args.kubeAPIConfig, Config.KubeAPIConfig) || tt.args.nodeName != Config.NodeName || tt.args.edgesite != Config.EdgeSiteEnable {
				t.Errorf("TestInitCnofigure() failed. got: %v/%v/%v/%v want: %v/%v/%v/%v", Config.EdgeController, Config.KubeAPIConfig, Config.NodeName, Config.EdgeSiteEnable, *tt.args.ec, *tt.args.kubeAPIConfig, tt.args.nodeName, tt.args.edgesite)
			}
		})
	}
}
