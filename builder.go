package htmlgo

import (
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

func (t *TagBuilder[T]) Write(ctx *Context) error {
	return t.HTMLTag.Write(ctx)
}
