package admissioncontroller

import (
	"fmt"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/klog/v2"

	rulesv1 "github.com/kubeedge/kubeedge/pkg/apis/rules/v1"
)

var (
	sourceToTarget = [][2]rulesv1.RuleEndpointTypeDef{
		{rulesv1.RuleEndpointTypeRest, rulesv1.RuleEndpointTypeEventBus},
		{rulesv1.RuleEndpointTypeRest, rulesv1.RuleEndpointTypeServiceBus},
		{rulesv1.RuleEndpointTypeEventBus, rulesv1.RuleEndpointTypeRest},
	}
)

func admitRule(review admissionv1.AdmissionReview) *admissionv1.AdmissionResponse {
	reviewResponse := admissionv1.AdmissionResponse{}

	switch review.Request.Operation {
	case admissionv1.Create:
		raw := review.Request.Object.Raw
		rule := rulesv1.Rule{}
		deserializer := codecs.UniversalDeserializer()
		if _, _, err := deserializer.Decode(raw, nil, &rule); err != nil {
			klog.Errorf("validation failed with error: %v", err)
			return toAdmissionResponse(err)
		}
		err := validateRule(&rule)
		if err != nil {
			return toAdmissionResponse(err)
		}
		reviewResponse.Allowed = true
		return &reviewResponse
	case admissionv1.Delete, admissionv1.Connect:
		//no rule defined for above operations, greenlight for all of above.
		reviewResponse.Allowed = true
		return &reviewResponse
	default:
		err := fmt.Errorf("Unsupported webhook operation %v", review.Request.Operation)
		klog.Errorf("Unsupported webhook operation %v", review.Request.Operation)
		return toAdmissionResponse(err)
	}
}

func validateRule(rule *rulesv1.Rule) error {
	sourceKey := fmt.Sprintf("%s/%s", rule.Namespace, rule.Spec.Source)
	sourceEndpoint, err := controller.getRuleEndpoint(rule.Namespace, rule.Spec.Source)
	if err != nil {
		return fmt.Errorf("cant get source ruleEndpoint %s. Reason: %w", sourceKey, err)
	} else if sourceEndpoint == nil {
		return fmt.Errorf("source ruleEndpoint %s has not been created", sourceKey)
	}
	if err = validateSourceRuleEndpoint(sourceEndpoint, rule.Spec.SourceResource); err != nil {
		return err
	}
	targetKey := fmt.Sprintf("%s/%s", rule.Namespace, rule.Spec.Target)
	targetEndpoint, err := controller.getRuleEndpoint(rule.Namespace, rule.Spec.Target)
	if err != nil {
		return fmt.Errorf("cant get target ruleEndpoint %s. Reason: %w", targetKey, err)
	} else if targetEndpoint == nil {
		return fmt.Errorf("target ruleEndpoint %s has not been created", targetKey)
	}
	if err = validateTargetRuleEndpoint(targetEndpoint, rule.Spec.TargetResource); err != nil {
		return err
	}
	var exist bool
	for _, s2t := range sourceToTarget {
		if s2t[0] == sourceEndpoint.Spec.RuleEndpointType && s2t[1] == targetEndpoint.Spec.RuleEndpointType {
			exist = true
			break
		}
	}
	if !exist {
		return fmt.Errorf("the rule which is from source ruleEndpoint type %s to target ruleEndpoint type %s is not validate ",
			sourceEndpoint.Spec.RuleEndpointType, targetEndpoint.Spec.RuleEndpointType)
	}
	return nil
}
func validateSourceRuleEndpoint(ruleEndpoint *rulesv1.RuleEndpoint, sourceResource map[string]string) error {
	switch ruleEndpoint.Spec.RuleEndpointType {
	case rulesv1.RuleEndpointTypeRest:
		_, exist := sourceResource["path"]
		if !exist {
			return fmt.Errorf("\"path\" property missed in sourceResource when ruleEndpoint is \"rest\"")
		}
		rules, err := controller.listRule(ruleEndpoint.Namespace)
		if err != nil {
			return err
		}
		for _, r := range rules {
			if sourceResource["path"] == r.Spec.SourceResource["path"] {
				return fmt.Errorf("source properties exist in Rule %s/%s. Path: %s", r.Namespace, r.Name, sourceResource["path"])
			}
		}
	case rulesv1.RuleEndpointTypeEventBus:
		_, exist := sourceResource["topic"]
		if !exist {
			return fmt.Errorf("\"topic\" property missed in sourceResource when ruleEndpoint is \"eventbus\"")
		}
		_, exist = sourceResource["node_name"]
		if !exist {
			return fmt.Errorf("\"node_name\" property missed in sourceResource when ruleEndpoint is \"eventbus\"")
		}
		rules, err := controller.listRule(ruleEndpoint.Namespace)
		if err != nil {
			return err
		}
		for _, r := range rules {
			if sourceResource["topic"] == r.Spec.SourceResource["topic"] && sourceResource["node_name"] == r.Spec.SourceResource["node_name"] {
				return fmt.Errorf("source properties exist in Rule %s/%s. Node_name: %s, topic: %s", r.Namespace, r.Name, sourceResource["node_name"], sourceResource["topic"])
			}
		}
	}
	return nil
}

func validateTargetRuleEndpoint(ruleEndpoint *rulesv1.RuleEndpoint, targetResource map[string]string) error {
	switch ruleEndpoint.Spec.RuleEndpointType {
	case rulesv1.RuleEndpointTypeRest:
		_, exist := targetResource["resource"]
		if !exist {
			return fmt.Errorf("\"resource\" property missed in targetResource when ruleEndpoint is \"rest\"")
		}
	case rulesv1.RuleEndpointTypeEventBus:
		_, exist := targetResource["topic"]
		if !exist {
			return fmt.Errorf("\"topic\" property missed in targetResource when ruleEndpoint is \"eventbus\"")
		}
	case rulesv1.RuleEndpointTypeServiceBus:
		_, exist := targetResource["path"]
		if !exist {
			return fmt.Errorf("\"path\" property missed in targetResource when ruleEndpoint is \"servicebus\"")
		}
	}
	return nil
}

func serveRule(w http.ResponseWriter, r *http.Request) {
	serve(w, r, admitRule)
}
