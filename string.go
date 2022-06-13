package snamer

import (
	"strings"
)

// https://github.com/iancoleman/strcase contains bugs and problems.
// https://github.com/iancoleman/strcase/issues/39
// https://github.com/iancoleman/strcase/issues/38

func pascalStringToCamel(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := strings.ToLower(s[0:1]) + s[1:]
	return result
}
