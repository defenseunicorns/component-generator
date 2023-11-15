package component

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/defenseunicorns/component-generator/src/internal/oscal"
	"github.com/defenseunicorns/component-generator/src/internal/types"
	"gopkg.in/yaml.v3"
)

func BuildOscalDocument(config types.ComponentsConfig) (string, types.JsonMap, error) {
	var (
		aggregateOscalMap   = types.JsonMap{}
		backMatterResources = []interface{}{}
		components          = []interface{}{}
		rfc3339Time         = time.Now().Format(time.RFC3339)
		documents           = []types.OscalComponentDocumentGeneric{}
	)

	for _, local := range config.Components.Locals {
		document, err := oscal.GetOscalComponentFromLocal(config.BaseDirectory + local.Name)
		if err != nil {
			return "", aggregateOscalMap, err
		}
		documents = append(documents, document)
	}

	for _, remote := range config.Components.Remotes {

		if git := remote.Git; git != "" {
			if !strings.Contains(git, "@") {
				return "", aggregateOscalMap, fmt.Errorf("remote git URL must specify a git ref using the following syntax: 'https://github.com/<org>/<repo>@<git ref>'")
			}
			split := strings.Split(git, "@")
			document, err := oscal.GetOscalComponentDocumentFromRepo(split[0], split[1], remote.Path)
			if err != nil {
				return "", aggregateOscalMap, fmt.Errorf("no OSCAL document was found for %v", git)
			}
			documents = append(documents, document)
		}

	}

	// Collect the components and back-matter fields from component definitions
	for _, doc := range documents {
		components = append(components, doc.ComponentDefinition["components"].([]interface{})...)
		backMatterResources = append(backMatterResources, doc.ComponentDefinition["back-matter"].(types.JsonMap)["resources"].([]interface{})...)
	}

	config.Metadata["last-modified"] = rfc3339Time
	// Populate the aggregated component definition
	aggregateOscalMap = types.JsonMap{
		"component-definition": types.JsonMap{
			"components": components,
			"back-matter": types.JsonMap{
				"resources": backMatterResources,
			},
			"metadata": config.Metadata,
		},
	}

	yamlDocBytes, err := yaml.Marshal(aggregateOscalMap)
	if err != nil {
		return "", aggregateOscalMap, err
	}
	return string(yamlDocBytes), aggregateOscalMap, nil
}

// DiffComponentObjects compares two OSCAL component definitions.
// If they're the same, it returns true.
// If they're different, it returns false.
func DiffComponentObjects(origObj types.JsonMap, newObj types.JsonMap) bool {
	// Compare the metadata structs and the list of components
	// in-scope set LastModified to empty string to remove it from consideration
	if origObj["component-definition"].(types.JsonMap)["metadata"] != nil {
		origObj["component-definition"].(types.JsonMap)["metadata"].(types.JsonMap)["last-modified"] = ""
	}
	if newObj["component-definition"].(types.JsonMap)["metadata"] != nil {
		newObj["component-definition"].(types.JsonMap)["metadata"].(types.JsonMap)["last-modified"] = ""
	}

	metaCompare := reflect.DeepEqual(origObj["component-definition"].(types.JsonMap)["metadata"], newObj["component-definition"].(types.JsonMap)["metadata"])

	childCompare := reflect.DeepEqual(origObj["component-definition"].(types.JsonMap)["components"], newObj["component-definition"].(types.JsonMap)["components"])

	return childCompare && metaCompare
}
