package snamer

import (
	"regexp"
	"strings"
)

// We don't want to depend on external project.
// https://github.com/iancoleman/strcase

func pascalStringToCamel(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := strings.ToLower(s[0:1]) + s[1:]
	return result
}

func pascalStringToSnake(s string) string {
	if len(s) == 0 {
		return ""
	}
	re := regexp.MustCompile(`([A-Z])`)
	return strings.ToLower(s[0:1] + re.ReplaceAllString(s[1:], `_$1`))
}
