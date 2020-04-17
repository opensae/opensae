package fc

import (
	"net/http"
)

type FuncSpec struct {
	Name    string
	Service string

	Upstream string
}

func NewFunc(name, service string) *FuncSpec {
	return &FuncSpec{Name: name, Service: service, Upstream: "http://localhost:8081"}
}

type Context struct {
	f *FuncSpec

	req *http.Request
}

func (c *Context) Method() string {
	return c.req.Method
}

func (c *Context) URL() string {
	return c.f.Upstream
}

func NewContext(name, service string) *Context {
	return &Context{
		f: NewFunc(name, service),
	}
}

func (c *Context) WithReq(req *http.Request) *Context {
	c.req = req
	return c
}
