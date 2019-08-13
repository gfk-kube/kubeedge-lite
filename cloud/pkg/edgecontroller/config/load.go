package config

import (
	"k8s.io/klog"

	"github.com/kubeedge/beehive/pkg/common/config"

	"github.com/kubeedge/kubeedge/common/constants"
)

// UpdatePodStatusWorkers is the count of goroutines of update pod status
var UpdatePodStatusWorkers int

// UpdateNodeStatusWorkers is the count of goroutines of update node status
var UpdateNodeStatusWorkers int

// QueryConfigMapWorkers is the count of goroutines of query configmap
var QueryConfigMapWorkers int

// QuerySecretWorkers is the count of goroutines of query secret
var QuerySecretWorkers int

// QueryServiceWorkers is the count of goroutines of query service
var QueryServiceWorkers int

// QueryEndpointsWorkers is the count of goroutines of query endpoints
var QueryEndpointsWorkers int

func InitLoadConfig() {
	if psw, err := config.CONFIG.GetValue("controller.load.update-pod-status-workers").ToInt(); err != nil {
		UpdatePodStatusWorkers = constants.DefaultUpdatePodStatusWorkers
	} else {
		UpdatePodStatusWorkers = psw
	}
	klog.Infof("update pod status workers: %d", UpdatePodStatusWorkers)

	if nsw, err := config.CONFIG.GetValue("controller.load.update-node-status-workers").ToInt(); err != nil {
		UpdateNodeStatusWorkers = constants.DefaultUpdateNodeStatusWorkers
	} else {
		UpdateNodeStatusWorkers = nsw
	}
	klog.Infof("update node status workers: %d", UpdateNodeStatusWorkers)

	if qcw, err := config.CONFIG.GetValue("controller.load.query-configmap-workers").ToInt(); err != nil {
		QueryConfigMapWorkers = constants.DefaultQueryConfigMapWorkers
	} else {
		QueryConfigMapWorkers = qcw
	}
	klog.Infof("query config map workers: %d", QueryConfigMapWorkers)

	if qsw, err := config.CONFIG.GetValue("controller.load.query-secret-workers").ToInt(); err != nil {
		QuerySecretWorkers = constants.DefaultQuerySecretWorkers
	} else {
		QuerySecretWorkers = qsw
	}
	klog.Infof("query secret workers: %d", QuerySecretWorkers)

	if qsw, err := config.CONFIG.GetValue("controller.load.query-service-workers").ToInt(); err != nil {
		QueryServiceWorkers = constants.DefaultQueryServiceWorkers
	} else {
		QueryServiceWorkers = qsw
	}
	klog.Infof("query service workers: %d", QueryServiceWorkers)

	if qew, err := config.CONFIG.GetValue("controller.load.query-endpoints-workers").ToInt(); err != nil {
		QueryEndpointsWorkers = constants.DefaultQueryEndpointsWorkers
	} else {
		QueryEndpointsWorkers = qew
	}
	klog.Infof("query endpoints workers: %d", QueryEndpointsWorkers)
}
