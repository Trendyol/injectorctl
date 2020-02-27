package cmd

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"injectorctl/core"
	"io/ioutil"
	"log"
	"os"
)

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

	var resource map[string]interface{}

	if err := yaml.Unmarshal(sourceYAML, &resource); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Fprint(os.Stdout, core.FromSource2Injector(resource["kind"].(string)).Inject(resource))
}
