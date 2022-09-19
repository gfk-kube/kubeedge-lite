package client

import (
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/kubeedge/beehive/pkg/core/model"
	"github.com/kubeedge/kubeedge/edge/pkg/common/message"
	"github.com/kubeedge/kubeedge/edge/pkg/common/modules"
)

//PodsGetter is interface to get pods
type PodsGetter interface {
	Pods(namespace string) PodsInterface
}

//PodsInterface is pod interface
type PodsInterface interface {
	Create(*corev1.Pod) (*corev1.Pod, error)
	Update(*corev1.Pod) error
	Patch(name string, patchBytes []byte) (*corev1.Pod, error)
	Delete(name, options string) error
	Get(name string) (*corev1.Pod, error)
}

type pods struct {
	namespace string
	send      SendInterface
}

// PodResp represents pod response from the api-server
type PodResp struct {
	Object *corev1.Pod
	Err    *apierrors.StatusError
}

func newPods(namespace string, s SendInterface) *pods {
	return &pods{
		send:      s,
		namespace: namespace,
	}
}

func (c *pods) Create(cm *corev1.Pod) (*corev1.Pod, error) {
	return nil, nil
}

func (c *pods) Update(cm *corev1.Pod) error {
	return nil
}

func (c *pods) Delete(name, options string) error {
	resource := fmt.Sprintf("%s/%s/%s", c.namespace, model.ResourceTypePod, name)
	podDeleteMsg := message.BuildMsg(modules.MetaGroup, "", modules.EdgedModuleName, resource, model.DeleteOperation, options)
	c.send.Send(podDeleteMsg)
	return nil
}

func (c *pods) Get(name string) (*corev1.Pod, error) {
	resource := fmt.Sprintf("%s/%s/%s", c.namespace, model.ResourceTypePod, name)
	podMsg := message.BuildMsg(modules.MetaGroup, "", modules.EdgedModuleName, resource, model.QueryOperation, nil)
	msg, err := c.send.SendSync(podMsg)
	if err != nil {
		return nil, fmt.Errorf("get pod failed, err: %v", err)
	}

	content, err := msg.GetContentData()
	if err != nil {
		return nil, fmt.Errorf("parse message to pod failed, err: %v", err)
	}

	return handlePodFromMetaDB(content)
}

func (c *pods) Patch(name string, patchBytes []byte) (*corev1.Pod, error) {
	resource := fmt.Sprintf("%s/%s/%s", c.namespace, model.ResourceTypePodPatch, name)
	podMsg := message.BuildMsg(modules.MetaGroup, "", modules.EdgedModuleName, resource, model.PatchOperation, patchBytes)
	resp, err := c.send.SendSync(podMsg)
	if err != nil {
		return nil, fmt.Errorf("update pod failed, err: %v", err)
	}

	content, err := resp.GetContentData()
	if err != nil {
		return nil, fmt.Errorf("parse message to pod failed, err: %v", err)
	}

	return handlePodResp(content)
}

func handlePodFromMetaDB(content []byte) (*corev1.Pod, error) {
	var lists []string
	err := json.Unmarshal(content, &lists)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message to pod list from db failed, err: %v", err)
	}

	if len(lists) != 1 {
		return nil, fmt.Errorf("pod length from meta db is %d", len(lists))
	}

	var pod *corev1.Pod
	err = json.Unmarshal([]byte(lists[0]), &pod)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message to pod failed, err: %v", err)
	}
	return pod, nil
}

func handlePodResp(content []byte) (*corev1.Pod, error) {
	var podResp PodResp
	err := json.Unmarshal(content, &podResp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message to pod failed, err: %v", err)
	}

	return podResp.Object, podResp.Err
}
