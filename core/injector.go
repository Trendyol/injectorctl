package core

import (
	"log"
)

var annotations = map[string]string{
	"trendyol.com/config-inject":           "true",
	"trendyol.com/config-server-addr":      "http://consul-consul-server.platform:8500",
	"trendyol.com/config-output-file-path": "/out/output.txt",
	"trendyol.com/vault-server-addr":       "http://vault.platform:8200",
}

const (
	POD        = "Pod"
	DEPLOYMENT = "Deployment"
)

type injector interface {
	Inject(resource interface{}) string
}

func FromSource2Injector(resource string) injector {
	if resource == "" {
		log.Fatal("resource can not be empty")
	}
	switch resource {
	case POD:
		return new(PodsInjector)
	case DEPLOYMENT:
		return new(DeploymentsInjector)
	default:
		return nil
	}
}
