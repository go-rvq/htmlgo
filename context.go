package htmlgo

import (
	"context"
	"strings"
)

type Context struct {
	Writer
	Context context.Context
	depth   int
	writed  int
}

func (ctx *Context) Value(name any) interface{} {
	return ctx.Context.Value(name)
}

func NewContext(w Writer, ctx context.Context) *Context {
	return &Context{Writer: w, Context: ctx}
}

func (ctx *Context) Depth() int {
	return ctx.depth
}

func (ctx *Context) LeftSpace() string {
	return strings.Repeat("\t", ctx.depth)
}

func (ctx *Context) WriteLeftSpace() error {
	return ctx.WriteString(ctx.LeftSpace())
}

func (ctx *Context) Leave() {
	ctx.depth--
}

func (ctx *Context) Enter(f func() (err error)) error {
	defer ctx.EnterLeave()
	return f()
}

func (ctx *Context) EnterLeave() func() {
	ctx.depth++
	return func() {
		ctx.depth--
	}
}

func (ctx *Context) WriteChildren(l HTMLComponents) (err error) {
	defer ctx.EnterLeave()()
	return l.Write(ctx)
}

func (ctx *Context) Write(comp HTMLComponent) (err error) {
	return comp.Write(ctx)
}
