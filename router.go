package gowf

import (
	"net/http"
	"reflect"
)

type routerHandler struct {
	pattern        string
	controllerType reflect.Type
}

type Router struct {
	routers []*routerHandler
}

func (app *App) Route(path string, c ControllerInterface) {
	app.Handlers.Add(path, c)
}

func NewRouter() *Router {
	return &Router{routers: make([]*routerHandler, 0)}
}

func (p *Router) Add(pattern string, c ControllerInterface) {
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &routerHandler{}
	route.pattern = pattern
	route.controllerType = t

	p.routers = append(p.routers, route)
}

func (p *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	requestPath := r.URL.Path

	for _, route := range p.routers {

		if route.pattern != requestPath {
			continue
		}

		//Invoke the request handler
		vc := reflect.New(route.controllerType)
		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)

		ct := &Context{ResponseWriter: w, Request: r}

		in[0] = reflect.ValueOf(ct)
		in[1] = reflect.ValueOf(route.controllerType.Name())
		init.Call(in)
		in = make([]reflect.Value, 0)

		if r.Method == "GET" {
			method := vc.MethodByName("Get")
			method.Call(in)

		} else if r.Method == "POST" {
			method := vc.MethodByName("Post")
			method.Call(in)

		} else if r.Method == "DELETE" {
			method := vc.MethodByName("Delete")
			method.Call(in)

		} else if r.Method == "PUT" {
			method := vc.MethodByName("Put")
			method.Call(in)
		}

		break
	}

}
