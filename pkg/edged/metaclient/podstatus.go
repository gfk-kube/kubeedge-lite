package metaclient

import (
	"fmt"

	"edge-core/beehive/adoptions/common/api"
	"edge-core/beehive/pkg/core"
	"edge-core/beehive/pkg/core/context"
	"edge-core/beehive/pkg/core/model"
	"edge-core/pkg/common/message"
)

type PodStatusGetter interface {
	PodStatus(namespace string) PodStatusInterface
}

type PodStatusInterface interface {
	Create(*api.PodStatusRequest) (*api.PodStatusRequest, error)
	Update(rsName string, ps api.PodStatusRequest) error
	Delete(name string) error
	Get(name string) (*api.PodStatusRequest, error)
}

type podStatus struct {
	namespace string
	context   *context.Context
	send      SendInterface
}

func newPodStatus(namespace string, c *context.Context, s SendInterface) *podStatus {
	return &podStatus{
		context:   c,
		send:      s,
		namespace: namespace,
	}
}

func (c *podStatus) Create(ps *api.PodStatusRequest) (*api.PodStatusRequest, error) {
	return nil, nil
}

func (c *podStatus) Update(rsName string, ps api.PodStatusRequest) error {
	podStatusMsg := message.BuildMsg(core.MetaGroup, "", core.EdgedModuleName, c.namespace+"/"+model.ResourceTypePodStatus+"/"+rsName, model.UpdateOperation, ps)
	_, err := c.send.SendSync(podStatusMsg)
	if err != nil {
		return fmt.Errorf("update podstatus failed, err: %v", err)
	}

	return nil
}

func (c *podStatus) Delete(name string) error {
	return nil
}

func (c *podStatus) Get(name string) (*api.PodStatusRequest, error) {
	return nil, nil
}
