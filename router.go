package aero

import (
	"context"
	"log"
	"net/http"
	"regexp"
)

//Aero is a router handler
type Aero struct {
	NotFound       http.HandlerFunc
	routesByMethod map[string][]*mux
}

//Router is initialized router
func Router() *Aero {
	aero := Aero{}
	return &aero
}

func (aero *Aero) registerHandler(method string, path string, handlers ...http.HandlerFunc) {
	pathRegex, namedParams := pathRegex(path)

	muxMethod := mux{
		Path:     path,
		Params:   namedParams,
		Regex:    regexp.MustCompile(pathRegex),
		Handlers: handlers,
	}

	if len(aero.routesByMethod[method]) == 0 {
		aero.routesByMethod = map[string][]*mux{
			method: make([]*mux, 0),
		}
	}

	aero.routesByMethod[method] = append(aero.routesByMethod[method], &muxMethod)
}

//Get is to add a new GET route to the router
func (aero *Aero) Get(path string, handlers ...http.HandlerFunc) {
	aero.registerHandler("GET", path, handlers...)
}

//Post is to add a new POST route to the router
func (aero *Aero) Post(path string, handlers ...http.HandlerFunc) {
	aero.registerHandler("POST", path, handlers...)
}

//Put is to add a new PUT route to the router
func (aero *Aero) Put(path string, handlers ...http.HandlerFunc) {
	aero.registerHandler("PUT", path, handlers...)
}

//Delete is to add a new DELETE route to the router
func (aero *Aero) Delete(path string, handlers ...http.HandlerFunc) {
	aero.registerHandler("DELETE", path, handlers...)
}

//Patch is to add a new PATCH route to the router
func (aero *Aero) Patch(path string, handlers ...http.HandlerFunc) {
	aero.registerHandler("PATCH", path, handlers...)
}

//Head is to add a new HEAD route to the router
func (aero *Aero) Head(path string, handlers ...http.HandlerFunc) {
	aero.registerHandler("HEAD", path, handlers...)
}

//Options is to add a new OPTIONS route to the router
func (aero *Aero) Options(path string, handlers ...http.HandlerFunc) {
	aero.registerHandler("OPTIONS", path, handlers...)
}

//ServeHTTP is to implements `http.Handler` interface for this router to serves HTTP requests.
func (aero *Aero) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	match := false

	for _, requestedMethod := range aero.routesByMethod[r.Method] {
		if isMatchRoute, namedParams := requestedMethod.matchingPath(r.URL.Path); isMatchRoute {
			match = isMatchRoute
			if err := r.ParseForm(); err != nil {
				log.Printf("Error parsing form: %s", err)
				return
			}
			currentRequest := 0
			aero.nextWithContext(namedParams, requestedMethod.Handlers[currentRequest], w, r)
			currentRequest++
			break
		}
	}

	if !match {
		aero.notFoundHandler(w, r)
	}
}

func (aero *Aero) nextWithContext(p map[string]string, next http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	for k, v := range p {
		ctx = context.WithValue(ctx, k, v)
	}
	r = r.WithContext(ctx)
	next(w, r)
}

func (aero *Aero) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	if aero.NotFound != nil {
		aero.NotFound(w, r)
	} else {
		http.NotFound(w, r)
	}
}
