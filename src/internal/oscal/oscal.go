package oscal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/brandtkeller/component-generator/src/internal/http"
	"github.com/brandtkeller/component-generator/src/internal/types"
	yaml2 "github.com/ghodss/yaml"
)

func GetOscalComponentDocumentFromRepo(repo string, tag string, path string) (types.OscalComponentDocument, error) {
	var document types.OscalComponentDocument
	repo = strings.Replace(repo, ".git", "", -1)
	rawUrl := fmt.Sprintf("%s/-/raw/%s/%s", repo, tag, path)
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return document, err
	}
	responseCode, bytes, err := http.FetchFromHTTPResource(uri)
	if err != nil {
		return document, err
	}
	if responseCode != 200 {
		return document, fmt.Errorf("unexpected response code when downloading document: %v", responseCode)
	}
	jsonDoc, err := yaml2.YAMLToJSON(bytes)
	if err != nil {
		fmt.Printf("Error converting YAML to JSON: %s\n", err.Error())
	}
	err = json.Unmarshal(jsonDoc, &document)
	if err != nil {
		fmt.Printf("Error converting unmarshalling json: %s\n", err.Error())
	}
	// err = yaml.Unmarshal(bytes, &document)
	// if err != nil {
	// 	return document, err
	// }

	return document, nil
}

func GetOscalComponentFromLocal(path string) (types.OscalComponentDocument, error) {
	var document types.OscalComponentDocument

	rawDoc, err := os.ReadFile(path)
	if err != nil {
		return document, err
	}

	jsonDoc, err := yaml2.YAMLToJSON(rawDoc)
	if err != nil {
		fmt.Printf("Error converting YAML to JSON: %s\n", err.Error())
	}
	err = json.Unmarshal(jsonDoc, &document)
	if err != nil {
		fmt.Printf("Error converting unmarshalling json: %s\n", err.Error())
	}

	// err = yaml.Unmarshal(rawDoc, &document)
	// if err != nil {
	// 	return document, err
	// }
	return document, err
}
