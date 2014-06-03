package gowf

import "net/http"

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func (ctx *Context) WriteString(content string) {
	ctx.ResponseWriter.Write([]byte(content))
}
