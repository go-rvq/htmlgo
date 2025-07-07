package htmlgo

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"io"
)

type RawHTML string

func (s RawHTML) MarshalHTML(ctx context.Context) (r []byte, err error) {
	r = []byte(s)
	return
}

func Text(text string) (r HTMLComponent) {
	return RawHTML(html.EscapeString(text))
}

func Textf(format string, a ...interface{}) (r HTMLComponent) {
	return Text(fmt.Sprintf(format, a...))
}

type HTMLComponents []HTMLComponent

func Components(comps ...HTMLComponent) HTMLComponents {
	return HTMLComponents(comps)
}

func (hcs HTMLComponents) MarshalHTML(ctx context.Context) (r []byte, err error) {
	buf := bytes.NewBuffer(nil)
	for _, h := range hcs {
		if h == nil {
			continue
		}
		var b []byte
		b, err = h.MarshalHTML(ctx)
		if err != nil {
			return
		}
		buf.Write(b)
	}
	r = buf.Bytes()
	return
}

func Fprint(w io.Writer, root HTMLComponent, ctx context.Context) (err error) {
	if root == nil {
		return
	}
	var b []byte
	b, err = root.MarshalHTML(ctx)
	if err != nil {
		return
	}
	_, err = fmt.Fprint(w, string(b))
	return
}

func MustString(root HTMLComponent, ctx context.Context) string {
	b := bytes.NewBuffer(nil)
	err := Fprint(b, root, ctx)
	if err != nil {
		panic(err)
	}
	return b.String()
}

// Simplify Simplifies components walking over nested HTMLComponents and calls cb if
// component not is nil
func Simplify(c HTMLComponent, cb func(HTMLComponent)) {
	if c == nil {
		return
	}
	switch t := c.(type) {
	case HTMLComponents:
		for _, c := range t {
			Simplify(c, cb)
		}
	default:
		cb(c)
	}
}

func SimplifyE(c HTMLComponent, cb func(HTMLComponent) (err error)) (err error) {
	if c == nil {
		return
	}
	switch t := c.(type) {
	case HTMLComponents:
		for _, c := range t {
			if err = SimplifyE(c, cb); err != nil {
				return
			}
		}
	default:
		return cb(c)
	}
	return
}

func SimplifyItems(c HTMLComponent) (r HTMLComponents) {
	Simplify(c, func(t HTMLComponent) {
		r = append(r, t)
	})
	return
}

func SimplifyComponent(c HTMLComponent) HTMLComponent {
	var r HTMLComponents
	Simplify(c, func(t HTMLComponent) {
		r = append(r, t)
	})
	switch len(r) {
	case 0:
		return nil
	case 1:
		return r[0]
	default:
		return r
	}
}

func IsInline(c HTMLComponent) bool {
	switch child := c.(type) {
	case RawHTML:
		return true
	case *HTMLTagBuilder:
		if child.IsInLine {
			return true
		}
	case TagGetter:
		if child.GetHTMLTagBuilder().IsInLine {
			return true
		}
	}
	return false
}
