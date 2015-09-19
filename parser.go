package pqarray

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	unquotedChar  = `[^",\\{}\s(NULL)]`
	unquotedValue = fmt.Sprintf("(%s)+", unquotedChar)
	quotedChar    = `[^"\\]|\\"|\\\\`
	quotedValue   = fmt.Sprintf("\"(%s)*\"", quotedChar)
	arrayValue    = fmt.Sprintf("(?P<value>(%s|%s))", unquotedValue, quotedValue)
	arrayExp      = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", arrayValue))
	valueIndex    int
)

func init() {
	for i, subexp := range arrayExp.SubexpNames() {
		if subexp == "value" {
			valueIndex = i
			break
		}
	}
}

func parse(array string) []string {
	results := make([]string, 0)
	matches := arrayExp.FindAllStringSubmatch(array, -1)
	for _, match := range matches {
		s := match[valueIndex]
		s = strings.Trim(s, "\"")
		results = append(results, s)
	}
	return results
}
