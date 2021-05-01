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

package manager

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"

	"github.com/kubeedge/kubeedge/cloud/pkg/edgecontroller/config"
	"github.com/kubeedge/kubeedge/cloud/pkg/edgecontroller/utils"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/cloudcore/v1alpha1"
)

func TestConfigMapManager_Events(t *testing.T) {
	type fields struct {
		events chan watch.Event
	}

	ch := make(chan watch.Event, 1)
	tests := []struct {
		name   string
		fields fields
		want   chan watch.Event
	}{
		{
			"TestConfigMapManager_Events(): Case 1",
			fields{
				events: ch,
			},
			ch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmm := &ConfigMapManager{
				events: tt.fields.events,
			}
			if got := cmm.Events(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigMapManager.Events() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfigMapManager(t *testing.T) {
	type args struct {
		kubeClient *kubernetes.Clientset
		namespace  string
	}

	config.Config.KubeAPIConfig = v1alpha1.KubeAPIConfig{
		KubeConfig:  fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")),
		QPS:         100,
		Burst:       200,
		ContentType: "application/vnd.kubernetes.protobuf",
	}
	config.Config.Buffer = &v1alpha1.EdgeControllerBuffer{
		ConfigMapEvent: 1024,
	}

	cli, err := utils.KubeClient()
	if err != nil {
		t.Skip("No k8s cluster config file in $HOME/.kube/config, skip it.")
		return
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"TestNewConfigMapManager(): Case 1",
			args{
				cli,
				v1.NamespaceAll,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewConfigMapManager(tt.args.kubeClient, tt.args.namespace)
		})
	}
}
