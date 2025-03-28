package core

import (
	"sync"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Method string
	Path   string
	Name   string
	Group  string
}

type routeRegistry struct {
	echo   *echo.Echo
	routes []Route
	lookup map[string]Route
	mux    sync.RWMutex
}

var registry *routeRegistry = &routeRegistry{
	lookup: make(map[string]Route),
}

// InitRouter sets the echo instance for routing.
func InitRouter(e *echo.Echo) {
	registry.echo = e
}

// Route exposes static access to route registry methods.
var Route routeStatic

type routeStatic struct{}

// Register adds a new route to Echo and stores metadata.
func (r routeStatic) Register(method, path string, handler echo.HandlerFunc, name, group string) {
	registry.mux.Lock()
	defer registry.mux.Unlock()

	rt := Route{
		Method: method,
		Path:   path,
		Name:   name,
		Group:  group,
	}
	registry.routes = append(registry.routes, rt)
	if name != "" {
		registry.lookup[name] = rt
	}
	registry.echo.Add(method, path, handler)
}

// All returns all registered routes.
func (r routeStatic) All() []Route {
	registry.mux.RLock()
	defer registry.mux.RUnlock()
	return registry.routes
}

// Get returns a route by its name.
func (r routeStatic) Get(name string) (Route, bool) {
	registry.mux.RLock()
	defer registry.mux.RUnlock()
	rt, ok := registry.lookup[name]
	return rt, ok
}

// PathFor returns the path for a named route, with optional path params.
func (r routeStatic) PathFor(name string, params ...string) string {
	rt, ok := r.Get(name)
	if !ok {
		return ""
	}

	path := rt.Path

	if len(params)%2 != 0 {
		return path // ignore if params are malformed
	}

	// Replace :param in path
	for i := 0; i < len(params); i += 2 {
		key := ":" + params[i]
		val := params[i+1]
		path = strings.Replace(path, key, val, 1)
	}

	return path
}

