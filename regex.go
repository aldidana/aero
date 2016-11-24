package aero

import "strings"

func pathRegex(path string) (string, []string) {
	var items, namedParams []string
	var regexPath string

	parts := strings.Split(path, "/")
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			name := strings.Trim(part, ":")
			namedParams = append(namedParams, name)
			items = append(items, `([^\/]+)`)
		} else {
			items = append(items, part)
		}
	}
	regexPath = "^" + strings.Join(items, `\/`) + `/?` + "$"
	return regexPath, namedParams
}
