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

func NewWriter(w io.Writer) Writer {
	return &WriterImpl{w: w}
}

func WriteToContext(ctx *WriteContext, comp ...HTMLComponent) (err error) {
	return SimplifyE(HTMLComponents(comp), ctx.WriteComponent)
}

func Write(w io.Writer, comp ...HTMLComponent) (err error) {
	return WriteToContext(NewWriteContext(NewWriter(w), context.Background()), comp...)
}

func EnterWriteContext(w io.Writer, f func(ctx *WriteContext) error) (err error) {
	return f(NewWriteContext(NewWriter(w), context.Background()))
}

func WriteSliceToContext[C HTMLComponent](ctx *WriteContext, comp []C) (err error) {
	var ca any
	for _, c := range comp {
		ca = c
		if err = SimplifyE(ca.(HTMLComponent), func(c HTMLComponent) (err error) {
			if err = ctx.WriteComponent(c); err != nil {
				return
			}
			return ctx.WriteByte('\n')
		}); err != nil {
			return
		}
	}
	return
}

func WriteSlice[C HTMLComponent](w io.Writer, comp []C) (err error) {
	return WriteSliceToContext(NewWriteContext(NewWriter(w), context.Background()), comp)
}

func Marshall(c any, ctx context.Context) (_ []byte, err error) {
	var b strings.Builder
	switch t := c.(type) {
	case ComponentContextWriter:
		err = t.Write(NewWriteContext(NewWriter(&b), ctx))
		return []byte(b.String()), err
	case HTMLComponent:
		var b []byte
		if b, err = t.MarshalHTML(ctx); err != nil {
			return
		}
		return b, err
	}
	return
}
