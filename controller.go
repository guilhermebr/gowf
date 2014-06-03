package gowf

import (
	"net/http"
	"path"
	"text/template"
)

var TemplatePath = "templates"

type Controller struct {
	Ct        *Context
	Tpl       *template.Template
	Data      map[interface{}]interface{}
	ChildName string
	TplNames  string
	Layout    []string
	TplExt    string
}

type ControllerInterface interface {
	Init(ct *Context, cn string) //Initialize the context and subclass name
	Get()                        //method = GET processing
	Post()                       //method = POST processing
	Delete()                     //method = DELETE processing
	Put()                        //method = PUT handling
	Render() error               //method executed after the corresponding method to render the page
}

func (c *Controller) Init(ct *Context, cn string) {
	c.Data = make(map[interface{}]interface{})
	c.Layout = make([]string, 0)
	c.TplNames = ""
	c.ChildName = cn
	c.Ct = ct
	c.TplExt = "tpl"
}

func (c *Controller) Get() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Render() error {
	t, err := template.ParseFiles(path.Join(TemplatePath, c.TplNames))
	err = t.Execute(c.Ct.ResponseWriter, c.Data)
	return err

}
