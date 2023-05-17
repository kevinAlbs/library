package librarypkg

import (
	"github.com/kevinAlbs/dependency/dependencypkg"
)

// GetDependencyVersion returns the version of dependencypkg.
func GetDependencyVersion() string {
	return dependencypkg.GetVersion()
}
