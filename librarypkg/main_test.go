package librarypkg

import (
	"testing"
)

func TestGetDependencyVersion(t *testing.T) {

	if GetDependencyVersion() == "" {
		t.Error("expected non-empty GetDependencyVersion")
	}
}
