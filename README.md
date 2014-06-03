gowf
====

Go Web Framework

- Go Web Framework created for learning purposes.
- Based on Beego Web Framework tutorial.


## Quick Start
============
 ```go
 package main
 
 import (
 	"github.com/guilhermebr/gowf"
 )
 
 type IndexView struct {
 	gowf.Controller
 }
 
 func (self *IndexView) Get() {
 	self.Ct.WriteString("Hello Gowf(ers)!")
 }
 
 func main() {
	app := gowf.NewApp()
	app.Route("/", &IndexView{})
	app.Run()
 }
 ```
