package core

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

var annotations = map[string]string{
	"trendyol.com/config-inject":           "true",
	"trendyol.com/config-server-addr":      "http://consul-consul-server.platform:8500",
	"trendyol.com/config-output-file-path": "/out/output.txt",
	"trendyol.com/vault-server-addr":       "http://vault.platform:8200",
}

var injectors = map[metav1.GroupVersionResource]Injector{}

type Injector interface {
	Version() metav1.GroupVersionResource
	Inject(resource interface{}) string
}

func FromSource2Injector(resource string) Injector {
	if resource == "" {
		log.Fatal("resource can not be empty")
	}
	for _, injector := range injectors {
		if injector.Version().Resource == resource {
			return injector
		}
	}
	return nil
}
