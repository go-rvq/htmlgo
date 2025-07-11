package htmlgo_test

import (
	"context"
	"strings"
	"testing"

	. "github.com/go-rvq/htmlgo"
	"github.com/theplant/testingutils"
)

var htmltagCases = []struct {
	name     string
	tag      *HTMLTagBuilder
	expected string
}{
	{
		name: "case 1",
		tag: Div(
			Div().Text("Hello"),
		),
		expected: `
<div>
	<div>Hello</div>
</div>
`,
	},
	{
		name: "case 2",
		tag: Div(
			Div().Text("Hello").
				Attr("class", "menu",
					"id", "the-menu",
					"style").
				Attr("id", "menu-id"),
		),
		expected: `
<div>
	<div class='menu' id='menu-id'>Hello</div>
</div>
`,
	},
	{
		name: "escape 1",
		tag: Div(
			Div().Text("Hello").
				Attr("class", "menu",
					"id", "the><&\"'-menu",
					"style"),
		),
		expected: `
<div>
	<div class='menu' id='the><&"&#39;-menu'>Hello</div>
</div>
`,
	},
}

func TestHtmlTag(t *testing.T) {
	for _, c := range htmltagCases {
		r, err := Marshal(c.tag, context.TODO())
		if err != nil {
			panic(err)
		}
		diff := testingutils.PrettyJsonDiff(strings.TrimSpace(c.expected), string(r))
		if len(diff) > 0 {
			t.Error(c.name, diff)
		}
	}
}
