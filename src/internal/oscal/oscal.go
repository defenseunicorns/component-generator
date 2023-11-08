package oscal

import (
	"fmt"
	"os"

	"github.com/defenseunicorns/component-generator/src/internal/http"
	"github.com/defenseunicorns/component-generator/src/internal/types"
	"gopkg.in/yaml.v2"
)

func ReadFromRemote(repo string, tag string, path string) ([]byte, error) {
	uri, err := http.ConstructURL(repo, tag, path)
	if err != nil {
		return nil, fmt.Errorf("failed to construct git URL: %w", err)
	}
	responseCode, bytes, err := http.FetchFromHTTPResource(uri)
	if err != nil {
		return nil, err
	}
	if responseCode != 200 {
		return nil, fmt.Errorf("unexpected response code when downloading document: %v", responseCode)
	}
	return bytes, nil
}

func GetOscalComponentDocumentFromRepo(repo string, tag string, path string) (oscalDocument types.OscalComponentDocument, err error) {
	uri, err := http.ConstructURL(repo, tag, path)
	if err != nil {
		return oscalDocument, fmt.Errorf("failed to construct git URL: %w", err)
	}
	responseCode, bytes, err := http.FetchFromHTTPResource(uri)
	if err != nil {
		return oscalDocument, err
	}
	if responseCode != 200 {
		return oscalDocument, fmt.Errorf("unexpected response code when downloading document: %v", responseCode)
	}
	err = yaml.Unmarshal(bytes, &oscalDocument)
	if err != nil {
		return oscalDocument, err
	}

	return oscalDocument, nil
}

func ReadFromLocal(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func GetOscalComponentFromLocal(path string) (types.OscalComponentDocument, error) {
	var document types.OscalComponentDocument

	rawDoc, err := os.ReadFile(path)
	if err != nil {
		return document, err
	}

	err = yaml.Unmarshal(rawDoc, &document)
	if err != nil {
		return document, err
	}
	return document, err
}
