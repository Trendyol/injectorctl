package core

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	appsV1 "k8s.io/api/apps/v1"
	"log"
)

type DeploymentsInjector struct {
}

func (p *DeploymentsInjector) Inject(resource interface{}) string {
	deploymentResourceAsBytes, err := json.Marshal(resource)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var deploymentResourceAsStruct appsV1.Deployment
	if err := json.Unmarshal(deploymentResourceAsBytes, &deploymentResourceAsStruct); err != nil {
		log.Fatalf("error: %v", err)
	}

	for key, val := range annotations {
		a := deploymentResourceAsStruct.Spec.Template.ObjectMeta.Annotations
		if a != nil {
			a[key] = val
		} else {
			deploymentResourceAsStruct.Spec.Template.ObjectMeta.Annotations = map[string]string{}
		}
	}

	if marshal, err := yaml.Marshal(deploymentResourceAsStruct); err != nil {
		log.Fatalf("error: %v", err)
	} else {
		return string(marshal)
	}
	return ""
}
