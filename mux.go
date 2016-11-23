package aero

import (
	"net/http"
	"regexp"
)

type mux struct {
	Path     string
	Params   []string
	Regex    *regexp.Regexp
	Handlers []http.HandlerFunc
}

func (m *mux) matchingPath(path string) (bool, map[string]string) {
	routeParams := make(map[string]string)
	isMatch := false
	matchingRegex := m.Regex.FindAllStringSubmatch(path, -1)

	if isMatch = len(matchingRegex) != 0; isMatch {
		for i, param := range m.Params {
			routeParams[param] = matchingRegex[0][i+1]
		}
	}
	return isMatch, routeParams
}
