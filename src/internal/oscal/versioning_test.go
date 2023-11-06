package oscal

import (
	"testing"
)

const (
	nonExistentVersion     = "1.0.7"
	tooFewVersionNumbers   = "1.0"
	tooManyVersionNumbers  = "1.0.4.1"
	validVersion           = "1.0.4"
	validVersionWithDashes = "1-0-4"
	validVersionWithPrefix = "v1.0.4"
)

func TestOscalVersioning(t *testing.T) {
	t.Run("returns valid version when user version is in proper format", func(t *testing.T) {
		version, err := GetVersion(validVersion)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != validVersion {
			t.Errorf("expected %s, got %s", validVersion, version)
		}
	})

	t.Run("replaces dashes with periods when version given with dashes", func(t *testing.T) {
		version, err := GetVersion(validVersionWithDashes)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != validVersion {
			t.Errorf("expected %s, got %s", validVersion, version)
		}
	})

	t.Run("returns valid version when prefixed with v", func(t *testing.T) {
		version, err := GetVersion(validVersionWithPrefix)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != validVersion {
			t.Errorf("expected %s, got %s", validVersion, version)
		}
	})

	t.Run("uses the default oscal version when version is empty", func(t *testing.T) {
		version, err := GetVersion("")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != DEFAULT_OSCAL_VERSION {
			t.Errorf("expected %s, got %s", DEFAULT_OSCAL_VERSION, version)
		}
	})

	t.Run("throws error with invalid version structure", func(t *testing.T) {
		_, err := GetVersion(tooManyVersionNumbers)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("throws error with too few version numbers", func(t *testing.T) {
		_, err := GetVersion(tooFewVersionNumbers)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("throws error when version is not supported", func(t *testing.T) {
		_, err := GetVersion(nonExistentVersion)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})
}
