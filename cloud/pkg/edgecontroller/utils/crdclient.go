package utils

import (
	v1 "github.com/kubeedge/kubeedge/cloud/pkg/apis/rules/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

// NewCRDClient is used to create a restClient for crd
func NewCRDClient(cfg *rest.Config) (*rest.RESTClient, error) {
	scheme := runtime.NewScheme()
	schemeBuilder := runtime.NewSchemeBuilder(AddRuleCrds)

	err := schemeBuilder.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	config := *cfg
	config.APIPath = "/apis"
	config.GroupVersion = &v1.SchemeGroupVersion
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme).WithoutConversion()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		klog.Errorf("Failed to create REST Client due to error %v", err)
		return nil, err
	}

	return client, nil
}

func AddRuleCrds(scheme *runtime.Scheme) error {
	// Add rule
	scheme.AddKnownTypes(v1.SchemeGroupVersion, &v1.Rule{}, &v1.RuleList{})
	metav1.AddToGroupVersion(scheme, v1.SchemeGroupVersion)
	// Add rule-endpoint
	scheme.AddKnownTypes(v1.SchemeGroupVersion, &v1.RuleEndpoint{}, &v1.RuleEndpointList{})
	metav1.AddToGroupVersion(scheme, v1.SchemeGroupVersion)

	return nil
}
