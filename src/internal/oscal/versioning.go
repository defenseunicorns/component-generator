package oscal

import (
	"fmt"
	"regexp"
	"strings"
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
