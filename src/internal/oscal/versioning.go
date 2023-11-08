package oscal

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/defenseunicorns/component-generator/src/internal/types"
	V104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	V105 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-5"
	V106 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-6"
	V110 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-0"
	V111 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-1"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

const DEFAULT_OSCAL_VERSION = "1.0.4"

var versionRegexp = regexp.MustCompile(`^\d+([-\.]\d+){2}$`)

var supportedVersion = map[string]bool{
	"1.0.4": true,
	"1.0.5": true,
	"1.0.6": true,
	"1.1.0": true,
	"1.1.1": true,
}

type OscalDocumentHandler struct {
	ComponentsConfig types.ComponentsConfig
	version          string
	documents        []interface{}
	components       []interface{}
	resources        []interface{}
}

func NewOscalDocumentHandler() *OscalDocumentHandler {
	handler := &OscalDocumentHandler{
		ComponentsConfig: types.ComponentsConfig{},
		documents:        []interface{}{},
		components:       []interface{}{},
		resources:        []interface{}{},
	}
	handler.SetVersion(DEFAULT_OSCAL_VERSION)
	return handler
}

func (h *OscalDocumentHandler) SetVersion(version string) error {
	v, err := GetVersion(version)
	if err != nil {
		return err
	}
	h.version = v
	return nil
}

func (h *OscalDocumentHandler) GetVersion() string {
	return h.version
}

func (h *OscalDocumentHandler) PopulateComponents() error {
	for _, local := range h.ComponentsConfig.Components.Locals {
		document, err := ReadFromLocal(h.ComponentsConfig.BaseDirectory + local.Name)
		if err != nil {
			return err
		}
		err = h.AddDocument(document)
		if err != nil {
			return err
		}
	}

	for _, remote := range h.ComponentsConfig.Components.Remotes {
		if git := remote.Git; git != "" {
			if !strings.Contains(git, "@") {
				return fmt.Errorf("remote git URL must specify a git ref using the following syntax: 'https://github.com/<org>/<repo>@<git ref>'")
			}
			split := strings.Split(git, "@")
			document, err := ReadFromRemote(split[0], split[1], remote.Path)
			if err != nil {
				return fmt.Errorf("no OSCAL document was found for %v", git)
			}
			err = h.AddDocument(document)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *OscalDocumentHandler) AggregateOscalModel() (interface{}, error) {
	var aggregatedOscalModel interface{}
	h.ComponentsConfig.Metadata.LastModified = time.Now().Format(time.RFC3339)
	switch h.version {
	case "1.0.4":
		var components = []V104.DefinedComponent{}
		var backMatterResources = []V104.Resource{}
		var metadata = V104.Metadata{}
		for _, doc := range h.documents {
			typedDoc := doc.(V104.OscalModels)
			components = append(components, typedDoc.ComponentDefinition.Components...)
			backMatterResources = append(backMatterResources, typedDoc.ComponentDefinition.BackMatter.Resources...)
		}
		metadataBytes, err := json.Marshal(h.ComponentsConfig.Metadata)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(metadataBytes, &metadata)
		if err != nil {
			return nil, err
		}

		aggregatedOscalModel = V104.OscalModels{
			ComponentDefinition: V104.ComponentDefinition{
				UUID:       uuid.NewString(),
				Components: []V104.DefinedComponent(components),
				BackMatter: V104.BackMatter{
					Resources: backMatterResources,
				},
				Metadata: metadata,
			},
		}
	case "1.0.5":
		var components = []V105.DefinedComponent{}
		var backMatterResources = []V105.Resource{}
		var metadata = V105.Metadata{}
		for _, doc := range h.documents {
			typedDoc := doc.(V105.OscalModels)
			components = append(components, typedDoc.ComponentDefinition.Components...)
			backMatterResources = append(backMatterResources, typedDoc.ComponentDefinition.BackMatter.Resources...)
		}
		metadataBytes, err := json.Marshal(h.ComponentsConfig.Metadata)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(metadataBytes, &metadata)
		if err != nil {
			return nil, err
		}
		aggregatedOscalModel = V105.OscalModels{
			ComponentDefinition: V105.ComponentDefinition{
				UUID:       uuid.NewString(),
				Components: []V105.DefinedComponent(components),
				BackMatter: V105.BackMatter{
					Resources: backMatterResources,
				},
				Metadata: metadata,
			},
		}
	case "1.0.6":
		var components = []V106.DefinedComponent{}
		var backMatterResources = []V106.Resource{}
		var metadata = V106.Metadata{}
		for _, doc := range h.documents {
			typedDoc := doc.(V106.OscalModels)
			components = append(components, typedDoc.ComponentDefinition.Components...)
			backMatterResources = append(backMatterResources, typedDoc.ComponentDefinition.BackMatter.Resources...)
		}
		metadataBytes, err := json.Marshal(h.ComponentsConfig.Metadata)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(metadataBytes, &metadata)
		if err != nil {
			return nil, err
		}
		aggregatedOscalModel = V106.OscalModels{
			ComponentDefinition: V106.ComponentDefinition{
				UUID:       uuid.NewString(),
				Components: []V106.DefinedComponent(components),
				BackMatter: V106.BackMatter{
					Resources: backMatterResources,
				},
				Metadata: metadata,
			},
		}
	case "1.1.0":
		var components = []V110.DefinedComponent{}
		var backMatterResources = []V110.Resource{}
		var metadata = V110.Metadata{}
		for _, doc := range h.documents {
			typedDoc := doc.(V110.OscalModels)
			components = append(components, typedDoc.ComponentDefinition.Components...)
			backMatterResources = append(backMatterResources, typedDoc.ComponentDefinition.BackMatter.Resources...)
		}
		metadataBytes, err := json.Marshal(h.ComponentsConfig.Metadata)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(metadataBytes, &metadata)
		if err != nil {
			return nil, err
		}
		aggregatedOscalModel = V110.OscalModels{
			ComponentDefinition: V110.ComponentDefinition{
				UUID:       uuid.NewString(),
				Components: []V110.DefinedComponent(components),
				BackMatter: V110.BackMatter{
					Resources: backMatterResources,
				},
				Metadata: metadata,
			},
		}
	case "1.1.1":
		var components = []V111.DefinedComponent{}
		var backMatterResources = []V111.Resource{}
		var metadata = V111.Metadata{}
		for _, doc := range h.documents {
			typedDoc := doc.(V111.OscalModels)
			components = append(components, typedDoc.ComponentDefinition.Components...)
			backMatterResources = append(backMatterResources, typedDoc.ComponentDefinition.BackMatter.Resources...)
		}
		metadataBytes, err := json.Marshal(h.ComponentsConfig.Metadata)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(metadataBytes, &metadata)
		if err != nil {
			return nil, err
		}
		aggregatedOscalModel = V111.OscalModels{
			ComponentDefinition: V111.ComponentDefinition{
				UUID:       uuid.NewString(),
				Components: []V111.DefinedComponent(components),
				BackMatter: V111.BackMatter{
					Resources: backMatterResources,
				},
				Metadata: metadata,
			},
		}
	}

	return aggregatedOscalModel, nil
}

func (h *OscalDocumentHandler) AddDocument(docBytes []byte) error {
	var (
		err error
		doc interface{}
	)
	switch h.version {
	case "1.0.4":
		doc = V104.OscalModels{}
	case "1.0.5":
		doc = V105.OscalModels{}
	case "1.0.6":
		doc = V106.OscalModels{}
	case "1.1.0":
		doc = V110.OscalModels{}
	case "1.1.1":
		doc = V111.OscalModels{}
	}
	err = yaml.Unmarshal(docBytes, &doc)
	if err != nil {
		return err
	}
	h.documents = append(h.components, doc)
	return nil
}

// GetVersion returns formatted OSCAL version if valid version is passed, returns error if not.
func GetVersion(userVersion string) (string, error) {
	if userVersion == "" {
		return DEFAULT_OSCAL_VERSION, nil
	}
	builtVersion := formatUserVersion(userVersion)

	if !isVersionRegexp(builtVersion) {
		return builtVersion, fmt.Errorf("version %s is not a valid version", userVersion)
	}

	if !supportedVersion[builtVersion] {
		return builtVersion, fmt.Errorf("version %s is not supported", userVersion)
	}

	return builtVersion, nil
}

func isVersionRegexp(v string) bool {
	return versionRegexp.MatchString(v)
}

func formatUserVersion(v string) string {
	builtVersion := v
	if builtVersion[0] == 'v' {
		builtVersion = builtVersion[1:]
	}
	builtVersion = strings.ReplaceAll(builtVersion, "-", ".")
	return builtVersion
}
