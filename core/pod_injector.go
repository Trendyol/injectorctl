package core

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type PodsInjector struct {
}

func init() {
	podsInjector := &PodsInjector{}
	injectors[podsInjector.Version()] = podsInjector
}

func (p *PodsInjector) Version() metav1.GroupVersionResource {
	return metav1.GroupVersionResource{Version: "v1", Resource: "Pod"}
}

func (p *PodsInjector) Inject(resource interface{}) string {
	resourceAsBytes, err := json.Marshal(resource)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var podAsStruct coreV1.Pod
	if err := json.Unmarshal(resourceAsBytes, &podAsStruct); err != nil {
		log.Fatalf("error: %v", err)
	}

	for key, val := range annotations {
		a := podAsStruct.ObjectMeta.Annotations
		if a != nil {
			a[key] = val
		} else {
			podAsStruct.ObjectMeta.Annotations = map[string]string{}
		}
	}

	if marshal, err := yaml.Marshal(podAsStruct); err != nil {
		log.Fatalf("error: %v", err)
	} else {
		return string(marshal)
	}
	return ""
}
