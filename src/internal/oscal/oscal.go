package oscal

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/defenseunicorns/component-generator/src/internal/http"
	"github.com/defenseunicorns/component-generator/src/internal/types"
	"gopkg.in/yaml.v2"
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
	err = yaml.Unmarshal(bytes, &document)
	if err != nil {
		return document, err
	}

	return document, nil
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
