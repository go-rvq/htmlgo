package htmlgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"strings"
)

type RawHTML string

func (s RawHTML) Write(ctx *Context) (err error) {
	return ctx.WriteString(string(s))
}

func Text(text string) HTMLComponent {
	return RawHTML(html.EscapeString(text))
}

func Textf(format string, a ...interface{}) HTMLComponent {
	return Text(fmt.Sprintf(format, a...))
}

type HTMLComponents []HTMLComponent

func Components(comps ...HTMLComponent) HTMLComponents {
	return comps
}

func (l HTMLComponents) Write(ctx *Context) (err error) {
	var (
		leftSpace        = ctx.LeftSpace()
		last, lastNonRaw HTMLComponent
	)

	_ = lastNonRaw

	err = SimplifyE(l, func(child HTMLComponent) (err error) {
		last = child
		if !IsRaw(child) {
			if !IsInline(child) || (lastNonRaw != nil && !IsInline(lastNonRaw)) {
				if err = ctx.WriteString("\n" + leftSpace); err != nil {
					return
				}
			}
			lastNonRaw = child
		}
		return child.Write(ctx)
	})

	if err != nil {
		return
	}

	if last != nil && !IsInline(last) {
		if err = ctx.WriteString("\n"); err == nil && len(leftSpace) > 1 {
			err = ctx.WriteString(leftSpace[1:])
		}
	}

	return
}

func Fprint(w io.Writer, root HTMLComponent, ctx context.Context) (err error) {
	if root == nil {
		return
	}
	var b []byte
	b, err = Marshall(root, ctx)
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

func IsRaw(c HTMLComponent) (ok bool) {
	_, ok = c.(RawHTML)
	return
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

func JSONString(v interface{}) (r string) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	r = string(b)
	return
}

func EscapeAttr(str string) (r string) {
	r = strings.Replace(str, "'", "&#39;", -1)
	// r = strings.Replace(r, "\n", "", -1)
	return
}
