package htmlgo_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/go-rvq/htmlgo"
)

/*
Create a simple div, Text will be escaped by html
*/
func ExampleTag() {
	banner := "We write html in Go"
	comp := Div(
		Text("123<h1>"),
		Textf("Hello, %s", banner),
		Br(),
		Text("new line"),
	)
	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <div>123&lt;h1&gt;Hello, We write html in Go<br>new line</div>
}

/*
Create a full html page
*/
func ExampleHTML() {
	comp := HTML(
		Head(
			Meta().Charset("utf8"),
			Title("My test page"),
		),
		Body(
			Div(
				Img("images/firefox-icon.png").Alt("My test image"),
			),
		),
	)
	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <!DOCTYPE html>
	// <html>
	//	<head>
	//		<meta charset='utf8'>
	//		<title>My test page</title>
	//	</head>
	//	<body>
	//		<div><img src='images/firefox-icon.png' alt='My test image'></div>
	//	</body>
	// </html>
}

/*
Use ComponentFunc with RawHTML and Component
*/
func ExampleComponentFunc() {
	userProfile := func(username string, avatarURL string) HTMLComponent {
		return ComponentFunc(func(ctx *Context) (err error) {
			return Div(
				H1(username).Class("profileName"),
				P(Img(avatarURL).Class("profileImage")),
				Svg(RawHTML("the svg")),
			).Class("userProfile").Write(ctx)
		})
	}

	comp := Ul(
		Li(
			userProfile("felix<h1>", "http://image.com/img1.png"),
		),
		Li(
			userProfile("john", "http://image.com/img2.png"),
		),
	)
	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <ul>
	//	<li>
	//		<div class='userProfile'>
	//			<h1 class='profileName'>felix&lt;h1&gt;</h1>
	//			<p><img src='http://image.com/img1.png' class='profileImage'></p>
	//			<Svg>the svg</Svg>
	//		</div>
	//	</li>
	//	<li>
	//		<div class='userProfile'>
	//			<h1 class='profileName'>john</h1>
	//			<p><img src='http://image.com/img2.png' class='profileImage'></p>
	//			<Svg>the svg</Svg>
	//		</div>
	//	</li>
	// </ul>
}

/*
Use ComponentFunc with RawHTML and Component
*/
func ExampleComponentFunc_withRawHTML() {
	userProfile := func(username string, avatarURL string) HTMLComponent {
		return ComponentFunc(func(ctx *Context) (err error) {
			return Div(
				H1(username).Class("profileName"),
				Img(avatarURL).Class("profileImage"),
				P(RawHTML("<svg>complicated svg</svg>")),
				P(RawHTML("<custom-tag>custom 1</custom-tag>"), Span("text"), B("bolded")),
				P(RawHTML("<custom-tag>custom 2</custom-tag>\n\t\t\t\t"), Span("text 2"), B("bolded 2"), RawHTML("<custom-tag>custom 3</custom-tag>")),
				Div(RawHTML("end")),
			).Class("userProfile").Write(ctx)
		})
	}

	comp := Ul(
		Li(
			userProfile("felix<h1>", "http://image.com/img1.png"),
		),
		Li(
			userProfile("john", "http://image.com/img2.png"),
		),
	)

	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <ul>
	//	<li>
	//		<div class='userProfile'>
	//			<h1 class='profileName'>felix&lt;h1&gt;</h1>
	//			<img src='http://image.com/img1.png' class='profileImage'>
	//			<p><svg>complicated svg</svg></p>
	//			<p><custom-tag>custom 1</custom-tag><span>text</span><b>bolded</b></p>
	//			<p><custom-tag>custom 2</custom-tag>
	//				<span>text 2</span><b>bolded 2</b><custom-tag>custom 3</custom-tag></p>
	//			<div>end</div>
	//		</div>
	//	</li>
	//	<li>
	//		<div class='userProfile'>
	//			<h1 class='profileName'>john</h1>
	//			<img src='http://image.com/img2.png' class='profileImage'>
	//			<p><svg>complicated svg</svg></p>
	//			<p><custom-tag>custom 1</custom-tag><span>text</span><b>bolded</b></p>
	//			<p><custom-tag>custom 2</custom-tag>
	//				<span>text 2</span><b>bolded 2</b><custom-tag>custom 3</custom-tag></p>
	//			<div>end</div>
	//		</div>
	//	</li>
	// </ul>
}

type MySelectBuilder struct {
	options  [][]string
	selected string
}

func MySelect() *MySelectBuilder {
	return &MySelectBuilder{}
}

func (b *MySelectBuilder) Options(opts [][]string) (r *MySelectBuilder) {
	b.options = opts
	return b
}

func (b *MySelectBuilder) Selected(selected string) (r *MySelectBuilder) {
	b.selected = selected
	return b
}

func (b *MySelectBuilder) Write(ctx *Context) (err error) {
	opts := HTMLComponents{}
	for _, op := range b.options {
		var opt HTMLComponent
		if op[0] == b.selected {
			opt = Option(op[1]).Value(op[0]).Attr("selected", "true")
		} else {
			opt = Option(op[1]).Value(op[0])
		}
		opts = append(opts, opt)
	}
	return Select(opts...).Write(ctx)
}

/*
More complicated customized component
*/
func ExampleMySelectBuilder() {
	/*
		Define MySelect as follows:

		type MySelectBuilder struct {
			options  [][]string
			selected string
		}

		func MySelect() *MySelectBuilder {
			return &MySelectBuilder{}
		}

		func (b *MySelectBuilder) Options(opts [][]string) (r *MySelectBuilder) {
			b.options = opts
			return b
		}

		func (b *MySelectBuilder) Selected(selected string) (r *MySelectBuilder) {
			b.selected = selected
			return b
		}

		func (b *MySelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
			opts := []HTMLComponent{}
			for _, op := range b.options {
				var opt HTMLComponent
				if op[0] == b.selected {
					opt = Option(op[1]).Value(op[0]).Attr("selected", "true")
				} else {
					opt = Option(op[1]).Value(op[0])
				}
				opts = append(opts, opt)
			}
			return Select(opts...).MarshalHTML(ctx)
		}
	*/

	comp := MySelect().Options([][]string{
		{"1", "label 1"},
		{"2", "label 2"},
		{"3", "label 3"},
	}).Selected("2")

	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <select>
	//	<option value='1'>label 1</option>
	//	<option value='2' selected='true'>label 2</option>
	//	<option value='3'>label 3</option>
	// </select>
}

/*
Write a little bit of JavaScript and stylesheet
*/
func ExampleScript() {
	comp := Div(
		Button("Hello").ID("hello"),
		Style(`
		.container {
			background-color: red;
		}
	`),

		Script(`
		var b = document.getElementById("hello")
		b.onclick = function(e){
			alert("Hello");
		}
	`),
	).Class("container")

	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <div class='container'><button id='hello'>Hello</button>
	//	<style type='text/css'>
	//		.container {
	//			background-color: red;
	//		}
	//	</style>
	//	<script type='text/javascript'>
	//		var b = document.getElementById("hello")
	//		b.onclick = function(e){
	//			alert("Hello");
	//		}
	//	</script>
	// </div>
}

/*
An example about how to integrate into http.Handler, and how to do layout, and how to use context.
*/
func ExampleHTML_httpHandler() {
	type User struct {
		Name string
	}
	type key string

	const currentUserKey key = "currentUser"

	userStatus := func() HTMLComponent {
		return ComponentFunc(func(ctx *Context) (err error) {
			if currentUser, ok := ctx.Value(currentUserKey).(*User); ok {
				return Div(
					Text(currentUser.Name),
				).Class("username").Write(ctx)
			}
			return Div(Text("Login")).Class("login").Write(ctx)
		})
	}

	myHeader := func() HTMLComponent {
		return Div(
			Text("header"),
			userStatus(),
		).Class("header")
	}
	myFooter := func() HTMLComponent {
		return Div(Text("footer")).Class("footer")
	}

	layout := func(in HTMLComponent) (out HTMLComponent) {
		out = HTML(
			Head(
				Meta().Charset("utf8"),
			),
			Body(
				myHeader(),
				in,
				myFooter(),
			),
		)
		return
	}

	getLoginUserFromCookie := func(r *http.Request) *User {
		return &User{Name: "felix"}
	}

	homeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getLoginUserFromCookie(r)
		ctx := context.WithValue(context.TODO(), currentUserKey, user)

		root := Div(
			Text("This is my home page"),
		)

		Fprint(w, layout(root), ctx)
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	homeHandler.ServeHTTP(w, r)

	fmt.Println(w.Body.String())

	// Output:
	// <!DOCTYPE html>
	// <html>
	//	<head>
	//		<meta charset='utf8'>
	//	</head>
	//	<body>
	//		<div class='header'>header
	//			<div class='username'>felix</div>
	//		</div>
	//		<div>This is my home page</div>
	//		<div class='footer'>footer</div>
	//	</body>
	// </html>
}

/*
An example show how to set different type of attributes
*/
func ExampleTag_withMutipleTypeAttrs() {
	type MoreData struct {
		Name  string
		Count int
	}
	comp := Div(
		Input("username").
			Type("checkbox").
			Attr("checked", true).
			Attr("more-data", &MoreData{Name: "felix", Count: 100}).
			Attr("max-length", 10),
		Input("username2").
			Type("checkbox").
			Attr("checked", false),
	)
	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <div><input name='username' type='checkbox' checked more-data='{"Name":"felix","Count":100}' max-length='10'><input name='username2' type='checkbox'></div>
}

/*
An example show how to set styles
*/
func ExampleTag_styles() {
	comp := Div().
		StyleIf("background-color:red; border:1px solid red;", true).
		StyleIf("color:blue", true)
	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <div style='background-color:red; border:1px solid red; color:blue;'></div>
}

/*
An example to use If, `Iff` is for body to passed in as an func for the body depends on if condition not to be nil, `If` is for directly passed in HTMLComponent
*/
func ExampleTag_iff() {
	type Person struct {
		Age int
	}
	var p *Person

	name := "Leon"
	comp := Div(
		Iff(p != nil && p.Age > 18, func() HTMLComponent {
			return Div().Text(name + ": Age > 18")
		}).ElseIf(p == nil, func() HTMLComponent {
			return Div().Text("No person named " + name)
		}).Else(func() HTMLComponent {
			return Div().Text(name + ":Age <= 18")
		}),
	)
	Fprint(os.Stdout, comp, context.TODO())
	// Output:
	// <div>
	//	<div>No person named Leon</div>
	// </div>
}
