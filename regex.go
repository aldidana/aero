package aero

import "strings"

func pathRegex(path string) (string, []string) {
	var items, namedParams []string
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			name := strings.Trim(part, ":")
			namedParams = append(namedParams, name)
			items = append(items, `([^\/]+)`)
		} else {
			items = append(items, part)
		}
	}
	//Match regex with and without trailing slash
	regexPath := "^" + strings.Join(items, `\/`) + `/?` + "$"
	return regexPath, namedParams
}
