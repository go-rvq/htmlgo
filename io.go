package htmlgo

import (
	"context"
	"io"
	"strings"
)

type Writer interface {
	WriteString(s string) (err error)
	WriteByte(b byte) (err error)
	Write(b []byte) (err error)
}

type WriterImpl struct {
	w io.Writer
}

func (w *WriterImpl) Writer() io.Writer {
	return w.w
}

func (w *WriterImpl) WriteString(s string) (err error) {
	_, err = w.w.Write([]byte(s))
	return
}

func (w *WriterImpl) WriteByte(b byte) (err error) {
	_, err = w.w.Write([]byte{b})
	return
}

func (w *WriterImpl) Write(b []byte) (err error) {
	_, err = w.w.Write(b)
	return
}

func ToWriter(w io.Writer) Writer {
	var wa any = w
	if wi, _ := wa.(Writer); wi != nil {
		return wi
	}
	return &WriterImpl{w: w}
}

func Write(w io.Writer, comp ...HTMLComponent) (err error) {
	return HTMLComponents(comp).Write(NewContext(ToWriter(w), context.Background()))
}

func Enter(w io.Writer, f func(ctx *Context) error) (err error) {
	return f(NewContext(ToWriter(w), context.Background()))
}

func Marshal(comp HTMLComponent, ctx context.Context) (_ []byte, err error) {
	var b strings.Builder
	err = comp.Write(NewContext(ToWriter(&b), ctx))
	return []byte(b.String()), err
}

func MarshallString(c HTMLComponent, ctx context.Context) (string, error) {
	b, err := Marshal(c, ctx)
	return string(b), err
}
