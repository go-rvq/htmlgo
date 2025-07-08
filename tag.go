package htmlgo

import (
	"fmt"
	"html"
	"slices"
	"strings"
)

type Attr struct {
	Key   string
	Value interface{}
}

func (a *Attr) Override(f func(old any) any) {
	a.Value = f(a.Value)
}

type Attrs []*Attr

func (a *Attrs) Remove(name string) {
	var s Attrs
	for _, v := range *a {
		if v.Key == name {
			continue
		}
		s = append(s, v)
	}
	*a = s
}

func (a *Attrs) RemoveMany(name ...string) {
	var s Attrs
	for _, v := range *a {
		if slices.Contains(name, v.Key) {
			continue
		}
		s = append(s, v)
	}
	*a = s
}

func (a Attrs) Get(name string) *Attr {
	for _, v := range a {
		if v.Key == name {
			return v
		}
	}
	return nil
}

type HTMLTagBuilder struct {
	TagName      string
	IsOmitEndTag bool
	Attrs        Attrs
	Styles       []string
	ClassNames   []string
	Childs       HTMLComponents
	IsInLine     bool
}

func (b *HTMLTagBuilder) Append(comp ...HTMLComponent) {
	b.Childs = append(b.Childs, comp...)
}

func (b *HTMLTagBuilder) Prepend(comp ...HTMLComponent) {
	b.Childs = append(comp, b.Childs...)
}

func (b *HTMLTagBuilder) GetChildren() HTMLComponents {
	return b.Childs
}

func (b *HTMLTagBuilder) SetChildren(comps HTMLComponents) {
	b.Childs = comps
}

func Tag(tag string, child ...HTMLComponent) *HTMLTagBuilder {
	return &HTMLTagBuilder{
		Childs:   child,
		TagName:  tag,
		IsInLine: InlineTags[tag],
	}
}

func (b *HTMLTagBuilder) Tag(v string) (r *HTMLTagBuilder) {
	b.TagName = v
	if !b.IsInLine {
		b.IsInLine = InlineTags[v]
	}
	return b
}

func (b *HTMLTagBuilder) Inline() (r *HTMLTagBuilder) {
	b.IsInLine = true
	return b
}

func (b *HTMLTagBuilder) OmitEndTag() (r *HTMLTagBuilder) {
	b.IsOmitEndTag = true
	return b
}

func (b *HTMLTagBuilder) Text(v string) (r *HTMLTagBuilder) {
	b.Children(Text(v))
	return b
}

func (b *HTMLTagBuilder) Children(comps ...HTMLComponent) (r *HTMLTagBuilder) {
	b.Childs = comps
	return b
}

func (b *HTMLTagBuilder) HasChilds() bool {
	return len(b.Childs) > 0
}

func (b *HTMLTagBuilder) SetAttr(k string, v interface{}) {
	for _, at := range b.Attrs {
		if at.Key == k {
			at.Value = v
			return
		}
	}
	b.Attrs = append(b.Attrs, &Attr{k, v})
}

func (b *HTMLTagBuilder) Attr(vs ...interface{}) (r *HTMLTagBuilder) {
	if len(vs)%2 != 0 {
		vs = append(vs, "")
	}

	for i := 0; i < len(vs); i = i + 2 {
		if key, ok := vs[i].(string); ok {
			b.SetAttr(key, vs[i+1])
		} else {
			panic(fmt.Sprintf("Attr Key must be string, but was %#+v", vs[i]))
		}
	}
	return b
}

func (b *HTMLTagBuilder) AttrIf(key, value interface{}, add bool) (r *HTMLTagBuilder) {
	if !add {
		return b
	}

	return b.Attr(key, value)
}

func (b *HTMLTagBuilder) GetAttr(key string) *Attr {
	return b.Attrs.Get(key)
}

func (b *HTMLTagBuilder) RemoveAttr(key ...string) *HTMLTagBuilder {
	b.Attrs.RemoveMany(key...)
	return b
}

func (b *HTMLTagBuilder) Class(names ...string) (r *HTMLTagBuilder) {
	b.addClass(names...)
	return b
}

func (b *HTMLTagBuilder) addClass(names ...string) (r *HTMLTagBuilder) {
	for _, n := range names {
		ins := strings.Split(n, " ")
		for _, in := range ins {
			tin := strings.TrimSpace(in)
			if len(tin) > 0 {
				b.ClassNames = append(b.ClassNames, tin)
			}
		}
	}
	return b
}

func (b *HTMLTagBuilder) ClassIf(name string, add bool) (r *HTMLTagBuilder) {
	if !add {
		return b
	}
	b.addClass(name)
	return b
}

func (b *HTMLTagBuilder) Data(vs ...string) (r *HTMLTagBuilder) {
	for i := 0; i < len(vs); i = i + 2 {
		b.Attr(fmt.Sprintf("data-%s", vs[i]), vs[i+1])
	}
	return b
}

func (b *HTMLTagBuilder) ID(v string) (r *HTMLTagBuilder) {
	b.Attr("id", v)
	return b
}

func (b *HTMLTagBuilder) Href(v string) (r *HTMLTagBuilder) {
	b.Attr("href", v)
	return b
}

func (b *HTMLTagBuilder) Rel(v string) (r *HTMLTagBuilder) {
	b.Attr("rel", v)
	return b
}

func (b *HTMLTagBuilder) Title(v string) (r *HTMLTagBuilder) {
	b.Attr("title", html.EscapeString(v))
	return b
}

func (b *HTMLTagBuilder) TabIndex(v int) (r *HTMLTagBuilder) {
	b.Attr("tabindex", v)
	return b
}

func (b *HTMLTagBuilder) Required(v bool) (r *HTMLTagBuilder) {
	b.Attr("required", v)
	return b
}

func (b *HTMLTagBuilder) Readonly(v bool) (r *HTMLTagBuilder) {
	b.Attr("readonly", v)
	return b
}

func (b *HTMLTagBuilder) Role(v string) (r *HTMLTagBuilder) {
	b.Attr("role", v)
	return b
}

func (b *HTMLTagBuilder) Alt(v string) (r *HTMLTagBuilder) {
	b.Attr("alt", v)
	return b
}

func (b *HTMLTagBuilder) Target(v string) (r *HTMLTagBuilder) {
	b.Attr("target", v)
	return b
}

func (b *HTMLTagBuilder) Name(v string) (r *HTMLTagBuilder) {
	b.Attr("name", v)
	return b
}

func (b *HTMLTagBuilder) Value(v string) (r *HTMLTagBuilder) {
	b.Attr("value", v)
	return b
}

func (b *HTMLTagBuilder) For(v string) (r *HTMLTagBuilder) {
	b.Attr("for", v)
	return b
}

func (b *HTMLTagBuilder) Style(v string) (r *HTMLTagBuilder) {
	b.addStyle(strings.Trim(v, ";"))
	return b
}

func (b *HTMLTagBuilder) StyleIf(v string, add bool) (r *HTMLTagBuilder) {
	if !add {
		return b
	}
	b.Style(v)
	return b
}

func (b *HTMLTagBuilder) addStyle(v string) (r *HTMLTagBuilder) {
	if len(v) > 0 {
		b.Styles = append(b.Styles, v)
	}

	return b
}

func (b *HTMLTagBuilder) Type(v string) (r *HTMLTagBuilder) {
	b.Attr("type", v)
	return b
}

func (b *HTMLTagBuilder) Placeholder(v string) (r *HTMLTagBuilder) {
	b.Attr("placeholder", v)
	return b
}

func (b *HTMLTagBuilder) Src(v string) (r *HTMLTagBuilder) {
	b.Attr("src", v)
	return b
}

func (b *HTMLTagBuilder) Property(v string) (r *HTMLTagBuilder) {
	b.Attr("property", v)
	return b
}

func (b *HTMLTagBuilder) Action(v string) (r *HTMLTagBuilder) {
	b.Attr("action", v)
	return b
}

func (b *HTMLTagBuilder) Method(v string) (r *HTMLTagBuilder) {
	b.Attr("method", v)
	return b
}

func (b *HTMLTagBuilder) Content(v string) (r *HTMLTagBuilder) {
	b.Attr("content", v)
	return b
}

func (b *HTMLTagBuilder) Charset(v string) (r *HTMLTagBuilder) {
	b.Attr("charset", v)
	return b
}

func (b *HTMLTagBuilder) Disabled(v bool) (r *HTMLTagBuilder) {
	b.Attr("disabled", v)
	return b
}

func (b *HTMLTagBuilder) Checked(v bool) (r *HTMLTagBuilder) {
	b.Attr("checked", v)
	return b
}

func (b *HTMLTagBuilder) AppendChildren(c ...HTMLComponent) (r *HTMLTagBuilder) {
	b.Childs = append(b.Childs, c...)
	return b
}

func (b *HTMLTagBuilder) PrependChildren(c ...HTMLComponent) (r *HTMLTagBuilder) {
	b.Childs = append(c, b.Childs...)
	return b
}

func (b *HTMLTagBuilder) Write(ctx *Context) (err error) {
	ctx.writed++

	class := strings.TrimSpace(strings.Join(b.ClassNames, " "))
	if len(class) > 0 {
		b.Attr("class", class)
	}

	styles := strings.TrimSpace(strings.Join(b.Styles, "; "))
	if len(styles) > 0 {
		b.Attr("style", styles+";")
	}

	var attrSegs []string
	for _, at := range b.Attrs {
		var val string
		var isBool bool
		var boolVal bool
		switch v := at.Value.(type) {
		case string:
			val = v
		case []byte:
			val = string(v)
		case []rune:
			val = string(v)
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			val = fmt.Sprintf(`%d`, v)
		case float32, float64:
			val = fmt.Sprintf(`%f`, v)
		case bool:
			boolVal = v
			isBool = true
		default:
			val = JSONString(v)
		}

		if len(val) == 0 && !isBool {
			continue
		}

		if isBool && !boolVal {
			continue
		}

		seg := fmt.Sprintf(`%s='%s'`, EscapeAttr(at.Key), EscapeAttr(val))
		if isBool && boolVal {
			seg = EscapeAttr(at.Key)
		}
		attrSegs = append(attrSegs, seg)
	}

	attrStr := ""
	if len(attrSegs) > 0 {
		attrStr = " " + strings.Join(attrSegs, " ")
	}

	if err = ctx.WriteString(fmt.Sprintf("<%s%s>", b.TagName, attrStr)); err != nil {
		return
	}

	if !b.IsOmitEndTag {
		if err = ctx.WriteChildren(b.Childs); err != nil {
			return
		}
		if err = ctx.WriteString(fmt.Sprintf("</%s>", b.TagName)); err != nil {
			return
		}
	}
	return
}

func (b *HTMLTagBuilder) String() string {
	return "<" + b.TagName + ">"
}
