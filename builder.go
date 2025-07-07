package htmlgo

import (
	"bytes"
	"context"
	"fmt"
)

type (
	TagGetter interface {
		GetHTMLTagBuilder() *HTMLTagBuilder
	}

	TagBuilderGetter[T any] interface {
		TagGetter
		GetTagBuilder() *TagBuilder[T]
	}

	TagBuilder[T any] struct {
		DOT     T
		HTMLTag *HTMLTagBuilder
	}
)

func NewTag[T TagBuilderGetter[T]](dot T, name string, children ...HTMLComponent) T {
	b := dot.GetTagBuilder()
	b.DOT = dot
	b.HTMLTag = Tag(name)
	if len(children) > 0 {
		b.HTMLTag.Childs = children
	}
	return dot
}

func (t *TagBuilder[T]) String() string {
	return fmt.Sprintf("<%s of %T />", t.TagName(), t.DOT)
}

func (t *TagBuilder[T]) TagName() string {
	return t.HTMLTag.TagName
}

func (t *TagBuilder[T]) GetHTMLTagBuilder() *HTMLTagBuilder {
	return t.HTMLTag
}

func (t *TagBuilder[T]) GetTagBuilder() *TagBuilder[T] {
	return t
}

func (t *TagBuilder[T]) Dot() T {
	return t.DOT
}

func (t *TagBuilder[T]) MarshalHTML(ctx context.Context) (r []byte, err error) {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()

	var (
		c          = NewWriteContext(NewWriter(buf), ctx)
		dotAny any = t.DOT
		w, _       = dotAny.(interface{ Write(ctx *WriteContext) error })
	)

	if w != nil {
		err = w.Write(c)
	} else {
		err = t.Write(c)
	}

	if err != nil {
		return
	}

	r = make([]byte, buf.Len())
	copy(r, buf.Bytes())
	return
}

func (t *TagBuilder[T]) Write(ctx *WriteContext) error {
	return t.HTMLTag.WriteToContext(ctx)
}
