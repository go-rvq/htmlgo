/*
## htmlgo

Type safe and modularize way to generate html on server side.
Download the package with `go get -v github.com/go-rvq/htmlgo` and import the package with `.` gives you simpler code:

	import (
		. "github.com/go-rvq/htmlgo"
	)

also checkout full API documentation at: https://godoc.org/github.com/go-rvq/htmlgo
*/
package htmlgo

type HTMLComponent interface {
	Write(w *Context) (err error)
}

type ComponentFunc func(ctx *Context) (err error)

func (f ComponentFunc) Write(ctx *Context) error {
	return f(ctx)
}

type MutableAttrHTMLComponent interface {
	HTMLComponent
	SetAttr(k string, v interface{})
}

type ContainerHTMLComponent interface {
	HTMLComponent
	HasChilds() bool
	Append(...HTMLComponent)
	Prepend(...HTMLComponent)
	GetChildren() HTMLComponents
	SetChildren(HTMLComponents)
}
