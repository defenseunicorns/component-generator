package component

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/defenseunicorns/component-generator/src/internal/oscal"
	"github.com/defenseunicorns/component-generator/src/internal/types"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

func BuildOscalDocument(config types.ComponentsConfig) (string, types.OscalComponentDocument, error) {
	var (
		backMatterResources = []types.Resources{}
		components          = []types.DefinedComponent{}
		rfc3339Time         = time.Now().Format(time.RFC3339)
		documents           = []types.OscalComponentDocument{}
	)

	for _, local := range config.Components.Locals {
		document, err := oscal.GetOscalComponentFromLocal(local.Name)
		if err != nil {
			return "", types.OscalComponentDocument{}, err
		}
		documents = append(documents, document)
	}

	for _, remote := range config.Components.Remotes {
		if git := remote.Git; git != "" {
			if strings.Contains(git, "@") {
				split := strings.Split(git, "@")
				document, err := oscal.GetOscalComponentDocumentFromGitlabRepo(split[0], split[1], remote.Path)
				if err != nil {
					log.Println(fmt.Errorf("No OSCAL document was found for Gitlab %v", git))
				}
				documents = append(documents, document)
			} else if strings.Contains(git, "github") {
				document, err := oscal.GetOscalComponentDocumentFromGitHubRepo(git+"/main/"+remote.Path)
				if err != nil {
					log.Println(fmt.Errorf("No OSCAL document was found for GitHub %v", git))
				}
				documents = append(documents, document)
			} else {
				log.Println(fmt.Errorf("Invalid Git URL format: %v", git))
			}
		}
	}

	// Collect the components and back-matter fields from component definitions
	for _, doc := range documents {
		components = append(components, doc.ComponentDefinition.Components...)
		backMatterResources = append(backMatterResources, doc.ComponentDefinition.BackMatter.Resources...)
	}

	config.Metadata.LastModified = rfc3339Time
	// Populate the aggregated component definition
	aggregateOscalDocument := types.OscalComponentDocument{
		ComponentDefinition: types.ComponentDefinition{
			UUID:       uuid.NewString(),
			Components: components,
			BackMatter: types.BackMatter{
				Resources: backMatterResources,
			},
			Metadata: config.Metadata,
		},
	}

	yamlDocBytes, err := yaml.Marshal(aggregateOscalDocument)
	if err != nil {
		return "", aggregateOscalDocument, err
	}
	return string(yamlDocBytes), aggregateOscalDocument, nil
}

// DiffComponentObjects compares two OSCAL component definitions.
// If they're the same, it returns true.
// If they're different, it returns false.
func DiffComponentObjects(origObj types.OscalComponentDocument, newObj types.OscalComponentDocument) bool {
	// Compare the metadata structs and the list of components
	// in-scope set LastModified to empty string to remove it from consideration
	origObj.ComponentDefinition.Metadata.LastModified = ""
	newObj.ComponentDefinition.Metadata.LastModified = ""

	metaCompare := reflect.DeepEqual(origObj.ComponentDefinition.Metadata, newObj.ComponentDefinition.Metadata)

	childCompare := reflect.DeepEqual(origObj.ComponentDefinition.Components, newObj.ComponentDefinition.Components)

	return childCompare && metaCompare
}
