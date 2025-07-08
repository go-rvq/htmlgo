package htmlgo

func (t *TagBuilder[T]) GetChildren() []HTMLComponent {
	return t.HTMLTag.Childs
}

func (t *TagBuilder[T]) Children(c ...HTMLComponent) T {
	t.HTMLTag.Children(c...)
	return t.DOT
}

func (t *TagBuilder[T]) AppendChild(c ...HTMLComponent) T {
	t.HTMLTag.AppendChildren(c...)
	return t.DOT
}

func (t *TagBuilder[T]) PrependChild(c ...HTMLComponent) T {
	t.HTMLTag.PrependChildren(c...)
	return t.DOT
}

func (t *TagBuilder[T]) GetAttr(key string) *Attr {
	return t.HTMLTag.Attrs.Get(key)
}

func (t *TagBuilder[T]) RemoveAttr(key ...string) T {
	t.HTMLTag.Attrs.RemoveMany(key...)
	return t.DOT
}

func (t *TagBuilder[T]) SetTag(v string) T {
	t.HTMLTag.Tag(v)
	return t.DOT
}

func (t *TagBuilder[T]) OmitEndTag() T {
	t.HTMLTag.OmitEndTag()
	return t.DOT
}

func (t *TagBuilder[T]) Text(v string) T {
	t.HTMLTag.Text(v)
	return t.DOT
}

func (t *TagBuilder[T]) SetAttr(k string, v interface{}) T {
	t.HTMLTag.SetAttr(k, v)
	return t.DOT
}
func (t *TagBuilder[T]) Attr(vs ...interface{}) T {
	t.HTMLTag.Attr(vs...)
	return t.DOT
}
func (t *TagBuilder[T]) AttrIf(key, value interface{}, add bool) T {
	t.HTMLTag.AttrIf(key, value, add)
	return t.DOT
}

func (t *TagBuilder[T]) Class(names ...string) T {
	t.HTMLTag.Class(names...)
	return t.DOT
}

func (t *TagBuilder[T]) ClassIf(name string, add bool) T {
	t.HTMLTag.ClassIf(name, add)
	return t.DOT
}

func (t *TagBuilder[T]) Data(vs ...string) T {
	t.HTMLTag.Data(vs...)
	return t.DOT
}

func (t *TagBuilder[T]) ID(v string) T {
	t.HTMLTag.ID(v)
	return t.DOT
}

func (t *TagBuilder[T]) Href(v string) T {
	t.HTMLTag.Href(v)
	return t.DOT
}

func (t *TagBuilder[T]) Rel(v string) T {
	t.HTMLTag.Rel(v)
	return t.DOT
}

func (t *TagBuilder[T]) Title(v string) T {
	t.HTMLTag.Title(v)
	return t.DOT
}

func (t *TagBuilder[T]) TabIndex(v int) T {
	t.HTMLTag.TabIndex(v)
	return t.DOT
}

func (t *TagBuilder[T]) Required(v bool) T {
	t.HTMLTag.Required(v)
	return t.DOT
}

func (t *TagBuilder[T]) Readonly(v bool) T {
	t.HTMLTag.Readonly(v)
	return t.DOT
}

func (t *TagBuilder[T]) Role(v string) T {
	t.HTMLTag.Role(v)
	return t.DOT
}

func (t *TagBuilder[T]) Alt(v string) T {
	t.HTMLTag.Alt(v)
	return t.DOT
}

func (t *TagBuilder[T]) Target(v string) T {
	t.HTMLTag.Target(v)
	return t.DOT
}

func (t *TagBuilder[T]) Name(v string) T {
	t.HTMLTag.Name(v)
	return t.DOT
}

func (t *TagBuilder[T]) Value(v string) T {
	t.HTMLTag.Value(v)
	return t.DOT
}

func (t *TagBuilder[T]) For(v string) T {
	t.HTMLTag.For(v)
	return t.DOT
}

func (t *TagBuilder[T]) Style(v string) T {
	t.HTMLTag.Style(v)
	return t.DOT
}

func (t *TagBuilder[T]) StyleIf(v string, add bool) T {
	t.HTMLTag.StyleIf(v, add)
	return t.DOT
}

func (t *TagBuilder[T]) Type(v string) T {
	t.HTMLTag.Type(v)
	return t.DOT
}

func (t *TagBuilder[T]) Placeholder(v string) T {
	t.HTMLTag.Placeholder(v)
	return t.DOT
}

func (t *TagBuilder[T]) Src(v string) T {
	t.HTMLTag.Src(v)
	return t.DOT
}

func (t *TagBuilder[T]) Property(v string) T {
	t.HTMLTag.Property(v)
	return t.DOT
}

func (t *TagBuilder[T]) Action(v string) T {
	t.HTMLTag.Action(v)
	return t.DOT
}

func (t *TagBuilder[T]) Method(v string) T {
	t.HTMLTag.Method(v)
	return t.DOT
}

func (t *TagBuilder[T]) Content(v string) T {
	t.HTMLTag.Content(v)
	return t.DOT
}

func (t *TagBuilder[T]) Charset(v string) T {
	t.HTMLTag.Charset(v)
	return t.DOT
}

func (t *TagBuilder[T]) Disabled(v bool) T {
	t.HTMLTag.Disabled(v)
	return t.DOT
}

func (t *TagBuilder[T]) Checked(v bool) T {
	t.HTMLTag.Checked(v)
	return t.DOT
}
