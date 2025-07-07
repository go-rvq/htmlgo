package htmlgo

import (
	"context"
	"fmt"
	"strings"
)

type WriteContext struct {
	Writer
	Context context.Context
	Depth   int
	writed  int
}

func NewWriteContext(w Writer, ctx context.Context) *WriteContext {
	return &WriteContext{Writer: w, Context: ctx}
}

func (c *WriteContext) LeftSpace() string {
	return strings.Repeat("\t", c.Depth)
}

func (c *WriteContext) WriteLeftSpace() error {
	return c.WriteString(c.LeftSpace())
}

func (c *WriteContext) Sub() func() {
	c.Depth++
	return func() {
		c.Depth--
	}
}

func (c *WriteContext) Enter(f func() (err error)) error {
	c.Depth++
	defer func() {
		c.Depth--
	}()
	return f()
}

func (c *WriteContext) WriteChildren(cs ...HTMLComponent) (err error) {
	c.Depth++
	defer func() {
		c.Depth--
	}()

	var (
		leftSpace = c.LeftSpace()
		last      HTMLComponent
	)

	err = SimplifyE(HTMLComponents(cs), func(child HTMLComponent) (err error) {
		last = child
		if !IsInline(child) {
			c.WriteString("\n" + leftSpace)
		}
		return c.WriteComponent(child)
	})

	if err != nil {
		return
	}

	if last != nil && !IsInline(last) {
		c.WriteString("\n" + leftSpace[1:])
	}

	return
}

func (c *WriteContext) WriteAny(comp any) error {
	switch t := comp.(type) {
	case ComponentContextWriter:
		return t.Write(c)
	case *HTMLTagBuilder:
		return t.WriteToContext(c)
	case HTMLComponent:
		var (
			b   []byte
			err error
		)
		if b, err = t.MarshalHTML(c.Context); err != nil {
			return err
		}
		return c.Write(b)
	}
	return fmt.Errorf("WriteContext.WriteAny: unknown component type %T", comp)
}

func (c *WriteContext) WriteComponent(comp HTMLComponent) error {
	return c.WriteAny(comp)
}

type ComponentContextWriter interface {
	Write(w *WriteContext) (err error)
}
