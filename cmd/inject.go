package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"io/ioutil"
	corev1 "k8s.io/api/apps/v1"
	"log"
	"os"
)

var annotations = map[string]string{
	"trendyol.com/config-inject":           "true",
	"trendyol.com/config-server-addr":      "http://consul-consul-server.platform:8500",
	"trendyol.com/config-output-file-path": "/out/output.txt",
	"trendyol.com/vault-server-addr":       "http://vault.platform:8200",
}

var (
	injectCmd = &cobra.Command{
		Use:     "inject",
		Aliases: []string{"inj", "ijt"},
		Short:   "This command injects predefined annotations to given valid yaml input",
		Run:     injectFunction,
	}
)

func init() {
	injectCmd.Flags().StringP("file", "f", "", "YAML from file")
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func readYAMLFromStdin() []byte {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return bytes
}

func injectFunction(cmd *cobra.Command, args []string) {
	fileFlagValue := cmd.Flag("file").Value.String()

	var sourceYAML []byte
	if isInputFromPipe() {
		sourceYAML = readYAMLFromStdin()
	} else if fileFlagValue != "" {
		if fileContent, err := ioutil.ReadFile(fileFlagValue); err != nil {
			log.Fatalf("error: %v", err)
		} else {
			sourceYAML = fileContent
		}
	}

	var deploymentResourceAsMap interface{}

	if err := yaml.Unmarshal(sourceYAML, &deploymentResourceAsMap); err != nil {
		log.Fatalf("error: %v", err)
	}

	deploymentResourceAsBytes, err := json.Marshal(deploymentResourceAsMap)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var deploymentResourceAsStruct corev1.Deployment
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
		fmt.Fprint(os.Stdout, string(marshal))
	}
}
