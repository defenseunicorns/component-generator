package component

import (
	"fmt"
	"os"
	"testing"

	"github.com/defenseunicorns/component-generator/src/internal/types"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestDiffComponentObjects(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		origObj        types.OscalComponentDocument
		newObj         types.OscalComponentDocument
		expectedResult bool
	}{
		{
			name: "No changes to components",
			origObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Components: []types.DefinedComponent{
						{UUID: "1"},
						{UUID: "2"},
					},
				},
			},
			newObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Components: []types.DefinedComponent{
						{UUID: "1"},
						{UUID: "2"},
					},
				},
			},
			expectedResult: true, // No changes to components, so the result should be true
		},
		{
			name: "No changes to metadata",
			origObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Metadata: types.Metadata{
						Version: "0.0.1",
					},
				},
			},
			newObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Metadata: types.Metadata{
						Version: "0.0.1",
					},
				},
			},
			expectedResult: true, // No changes to metadata, so the result should be true
		},
		{
			name: "Changes in components",
			origObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Components: []types.DefinedComponent{
						{UUID: "1"},
						{UUID: "2"},
					},
				},
			},
			newObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Components: []types.DefinedComponent{
						{UUID: "1"},
						{UUID: "3"},
					},
				},
			},
			expectedResult: false, // Changes to components were made, so the result should be false
		},
		{
			name: "Changes in metadata",
			origObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Metadata: types.Metadata{
						Version: "0.0.1",
					},
				},
			},
			newObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Metadata: types.Metadata{
						Version: "0.0.2",
					},
				},
			},
			expectedResult: false, // Changes to metadata were made, so the result should be false
		},
		{
			name: "Changes to 'metadata.LastModified' should be ignored",
			origObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Metadata: types.Metadata{
						LastModified: "2023-06-28T17:19:35-05:00",
					},
				},
			},
			newObj: types.OscalComponentDocument{
				ComponentDefinition: types.ComponentDefinition{
					Metadata: types.Metadata{
						LastModified: "2024-07-29T18:20:36-06:00",
					},
				},
			},
			expectedResult: true, // Changes to 'metadata.LastModified' were made, which shouldn't be detected, so the result should be true
		},
	}

	for _, testCase := range testCases {
		result := DiffComponentObjects(testCase.origObj, testCase.newObj)

		if result != testCase.expectedResult {
			t.Errorf("Test case '%s' failed. Expected: %v, got: %v", testCase.name, testCase.expectedResult, result)
		}
	}
}

// TestBuildOscalDocumentWithValidConfigFile tests that OSCAL component definition files are generated correctly using a valid config file.
func TestBuildOscalDocumentWithValidConfigFile(t *testing.T) {
	t.Parallel()

	// Read in the valid config file as test data
	configFilePath := "../../../testdata/input/valid-components.yaml"
	configFile, err := readConfigFile(t, configFilePath)
	if err != nil {
		t.Fatal(err)
	}

	// Generate OSCAL component definition based on the provided config file
	_, actualComponentDefinition, err := BuildOscalDocument(configFile)
	if err != nil {
		t.Fatal(err)
	}

	// Read in the expected component definition as test data
	expectedComponentDefinition, err := readTestComponentDefinitionFile(t)
	if err != nil {
		t.Fatal(err)
	}

	// Perform a diff of the expected output and actual output
	match := DiffComponentObjects(expectedComponentDefinition, actualComponentDefinition)

	// Fail the test if the actual output doesn't match the expected output
	if !match {
		t.Fatal("the generated OSCAL component definition does not match the expected test output")
	}
}

// TestBuildOscalDocumentWithInvalidConfigFile verifies that the BuildOscalDocument() function returns an error when an invalid config file is provided.
func TestBuildOscalDocumentWithInvalidConfigFile(t *testing.T) {
	t.Parallel()

	// Read in the invalid config file as test data
	configFilePath := "../../../testdata/input/invalid-components.yaml"
	configFile, err := readConfigFile(t, configFilePath)
	if err != nil {
		t.Fatal(err)
	}

	// Check that an error is returned
	_, _, err = BuildOscalDocument(configFile)
	require.Error(t, err)
	require.ErrorContains(t, err, "remote git URL must specify a git ref")
}

func readConfigFile(t *testing.T, filePath string) (configFile types.ComponentsConfig, err error) {
	t.Helper()

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return configFile, fmt.Errorf("failed to read in config file: %v", err)
	}

	if err = yaml.Unmarshal(configBytes, &configFile); err != nil {
		return configFile, fmt.Errorf("failed to unmarshal config file data: %v", err)
	}

	return configFile, nil
}

func readTestComponentDefinitionFile(t *testing.T) (componentDefinition types.OscalComponentDocument, err error) {
	t.Helper()

	testComponentDefinitionPath := "../../../testdata/output/test-data.yaml"

	componentDefinitionBytes, err := os.ReadFile(testComponentDefinitionPath)
	if err != nil {
		return componentDefinition, fmt.Errorf("failed to read in test component definition file: %v", err)
	}

	if err := yaml.Unmarshal(componentDefinitionBytes, &componentDefinition); err != nil {
		return componentDefinition, fmt.Errorf("failed to unmarshal test component definition file: %v", err)
	}

	return componentDefinition, err
}