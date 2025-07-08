package htmlgo

// A "a": HTMLAnchorElement;
func A(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("a", children...)
}

// Abbr "abbr": HTMLElement;
func Abbr(text string) (r *HTMLTagBuilder) {
	return Tag("abbr").Text(text)
}

// Address "address": HTMLElement;
func Address(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("address", children...)
}

//     "applet": HTMLAppletElement;
// Not support

// Area "area": HTMLAreaElement;
func Area() (r *HTMLTagBuilder) {
	return Tag("area").OmitEndTag()
}

// Article "article": HTMLElement;
func Article(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("article", children...)
}

// Aside "aside": HTMLElement;
func Aside(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("aside", children...)
}

// Audio "audio": HTMLAudioElement;
func Audio(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("audio", children...)
}

// B "b": HTMLElement;
func B(text string) (r *HTMLTagBuilder) {
	return Tag("b").Text(text)
}

// Base "base": HTMLBaseElement;
func Base() (r *HTMLTagBuilder) {
	return Tag("base").OmitEndTag()
}

//     "basefont": HTMLBaseFontElement;
// Not Support

// Bdi "bdi": HTMLElement;
func Bdi(text string) (r *HTMLTagBuilder) {
	return Tag("bdi").Text(text)
}

// Bdo "bdo": HTMLElement;
func Bdo(text string) (r *HTMLTagBuilder) {
	return Tag("bdo").Text(text)
}

// Blockquote "blockquote": HTMLQuoteElement;
func Blockquote(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("blockquote", children...)
}

// Body "body": HTMLBodyElement;
func Body(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("body", children...)
}

// Br "br": HTMLBRElement;
func Br() (r *HTMLTagBuilder) {
	return Tag("br").OmitEndTag()
}

// Button "button": HTMLButtonElement;
func Button(label string) (r *HTMLTagBuilder) {
	return Tag("button").Text(label)
}

// Canvas "canvas": HTMLCanvasElement;
func Canvas(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("canvas", children...)
}

// Caption "caption": HTMLTableCaptionElement;
func Caption(text string) (r *HTMLTagBuilder) {
	return Tag("caption").Text(text)
}

// Cite "cite": HTMLElement;
func Cite(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("cite", children...)
}

// Code "code": HTMLElement;
func Code(text string) (r *HTMLTagBuilder) {
	return Tag("code").Text(text)
}

// Col "col": HTMLTableColElement;
func Col() (r *HTMLTagBuilder) {
	return Tag("col").OmitEndTag()
}

// Colgroup "colgroup": HTMLTableColElement;
func Colgroup(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("colgroup", children...)
}

// Data "data": HTMLDataElement;
func Data(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("data", children...)
}

// Datalist "datalist": HTMLDataListElement;
func Datalist(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("datalist", children...)
}

// Dd "dd": HTMLElement;
func Dd(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dd", children...)
}

// Del "del": HTMLModElement;
func Del(text string) (r *HTMLTagBuilder) {
	return Tag("del").Text(text)
}

//     "details": HTMLDetailsElement;

func Details(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("details", children...)
}

// Dfn "dfn": HTMLElement;
func Dfn(text string) (r *HTMLTagBuilder) {
	return Tag("dfn").Text(text)
}

// Dialog "dialog": HTMLDialogElement;
func Dialog(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dialog", children...)
}

//     "dir": HTMLDirectoryElement;
// Not Support

// Div "div": HTMLDivElement;
func Div(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("div", children...)
}

// Dl "dl": HTMLDListElement;
func Dl(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dl", children...)
}

// Dt "dt": HTMLElement;
func Dt(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dt", children...)
}

// Em "em": HTMLElement;
func Em(text string) (r *HTMLTagBuilder) {
	return Tag("em").Text(text)
}

// Embed "embed": HTMLEmbedElement;
func Embed() (r *HTMLTagBuilder) {
	return Tag("embed").OmitEndTag()
}

// Fieldset "fieldset": HTMLFieldSetElement;
func Fieldset(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("fieldset", children...)
}

// Figcaption "figcaption": HTMLElement;
func Figcaption(text string) (r *HTMLTagBuilder) {
	return Tag("figcaption").Text(text)
}

// Figure "figure": HTMLElement;
func Figure(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("figure", children...)
}

//     "font": HTMLFontElement;
// Not Support

// Footer "footer": HTMLElement;
func Footer(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("footer", children...)
}

// Form "form": HTMLFormElement;
func Form(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("form", children...)
}

//     "frame": HTMLFrameElement;
// Not Support

//     "frameset": HTMLFrameSetElement;
// Not Support

// H1 "h1": HTMLHeadingElement;
func H1(text string) (r *HTMLTagBuilder) {
	return Tag("h1").Text(text)
}

// H2 "h2": HTMLHeadingElement;
func H2(text string) (r *HTMLTagBuilder) {
	return Tag("h2").Text(text)
}

// H3 "h3": HTMLHeadingElement;
func H3(text string) (r *HTMLTagBuilder) {
	return Tag("h3").Text(text)
}

// H4 "h4": HTMLHeadingElement;
func H4(text string) (r *HTMLTagBuilder) {
	return Tag("h4").Text(text)
}

// H5 "h5": HTMLHeadingElement;
func H5(text string) (r *HTMLTagBuilder) {
	return Tag("h5").Text(text)
}

// H6 "h6": HTMLHeadingElement;
func H6(text string) (r *HTMLTagBuilder) {
	return Tag("h6").Text(text)
}

//     "head": HTMLHeadElement;

func Head(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("head", children...)
}

// Header "header": HTMLElement;
func Header(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("header", children...)
}

// Hgroup "hgroup": HTMLElement;
func Hgroup(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("hgroup", children...)
}

// Hr "hr": HTMLHRElement;
func Hr() (r *HTMLTagBuilder) {
	return Tag("hr").OmitEndTag()
}

// HTML "html": HTMLHtmlElement;
func HTML(children ...HTMLComponent) (r HTMLComponent) {
	return HTMLComponents{
		RawHTML("<!DOCTYPE html>"),
		Tag("html", children...),
	}
}

// I "i": HTMLElement;
func I(text string) (r *HTMLTagBuilder) {
	return Tag("i").Text(text)
}

// Iframe "iframe": HTMLIFrameElement;
func Iframe(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("iframe", children...)
}

// Img "img": HTMLImageElement;
func Img(src string) (r *HTMLTagBuilder) {
	return Tag("img").OmitEndTag().Attr("src", src)
}

//     "input": HTMLInputElement;

func Input(name string) (r *HTMLTagBuilder) {
	return Tag("input").OmitEndTag().Attr("name", name)
}

// Ins "ins": HTMLModElement;
func Ins(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ins", children...)
}

// Kbd "kbd": HTMLElement;
func Kbd(text string) (r *HTMLTagBuilder) {
	return Tag("kbd").Text(text)
}

// Label "label": HTMLLabelElement;
func Label(text string) (r *HTMLTagBuilder) {
	return Tag("label").Text(text)
}

// Legend "legend": HTMLLegendElement;
func Legend(text string) (r *HTMLTagBuilder) {
	return Tag("legend").Text(text)
}

// Li "li": HTMLLIElement;
func Li(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("li", children...)
}

// Link "link": HTMLLinkElement;
func Link(href string) (r *HTMLTagBuilder) {
	return Tag("link").OmitEndTag().Attr("href", href)
}

// Main "main": HTMLElement;
func Main(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("main", children...)
}

// Map "map": HTMLMapElement;
func Map(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("map", children...)
}

// Mark "mark": HTMLElement;
func Mark(text string) (r *HTMLTagBuilder) {
	return Tag("mark").Text(text)
}

//     "marquee": HTMLMarqueeElement;
// Not Support

// Menu "menu": HTMLMenuElement;
func Menu(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("menu", children...)
}

//     "meta": HTMLMetaElement;

func Meta() (r *HTMLTagBuilder) {
	return Tag("meta").OmitEndTag()
}

// Meter "meter": HTMLMeterElement;
func Meter(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("meter", children...)
}

// Nav "nav": HTMLElement;
func Nav(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("nav", children...)
}

// Noscript "noscript": HTMLElement;
func Noscript(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("noscript", children...)
}

// Object "object": HTMLObjectElement;
func Object(data string) (r *HTMLTagBuilder) {
	return Tag("object").Attr("data", data)
}

// Ol "ol": HTMLOListElement;
func Ol(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ol", children...)
}

// Optgroup "optgroup": HTMLOptGroupElement;
func Optgroup(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("optgroup", children...)
}

// Option "option": HTMLOptionElement;
func Option(text string) (r *HTMLTagBuilder) {
	return Tag("option").Text(text)
}

// Output "output": HTMLOutputElement;
func Output(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("output", children...)
}

// P "p": HTMLParagraphElement;
func P(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("p", children...)
}

// Param "param": HTMLParamElement;
func Param(name string) (r *HTMLTagBuilder) {
	return Tag("param").OmitEndTag().Attr("name", name)
}

// Picture "picture": HTMLPictureElement;
func Picture(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("picture", children...)
}

// Pre "pre": HTMLPreElement;
func Pre(text string) (r *HTMLTagBuilder) {
	return Tag("pre").Text(text)
}

// Progress "progress": HTMLProgressElement;
func Progress(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("progress", children...)
}

// Q "q": HTMLQuoteElement;
func Q(text string) (r *HTMLTagBuilder) {
	return Tag("q").Text(text)
}

// Rp "rp": HTMLElement;
func Rp(text string) (r *HTMLTagBuilder) {
	return Tag("rp").Text(text)
}

// Rt "rt": HTMLElement;
func Rt(text string) (r *HTMLTagBuilder) {
	return Tag("rt").Text(text)
}

// Ruby "ruby": HTMLElement;
func Ruby(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ruby", children...)
}

// S "s": HTMLElement;
func S(text string) (r *HTMLTagBuilder) {
	return Tag("s").Text(text)
}

//     "samp": HTMLElement;

func Samp(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("samp", children...)
}

// Script "script": HTMLScriptElement;
func Script(script string) (r *HTMLTagBuilder) {
	return Tag("script").
		Attr("type", "text/javascript").
		Children(RawHTML(script))
}

// Section "section": HTMLElement;
func Section(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("section", children...)
}

// Select "select": HTMLSelectElement;
func Select(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("select", children...)
}

// Slot "slot": HTMLSlotElement;
func Slot(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("slot", children...)
}

// Small "small": HTMLElement;
func Small(text string) (r *HTMLTagBuilder) {
	return Tag("small").Text(text)
}

// Source "source": HTMLSourceElement;
func Source(src string) (r *HTMLTagBuilder) {
	return Tag("source").OmitEndTag().Attr("src", src)
}

// Span "span": HTMLSpanElement;
func Span(text string) (r *HTMLTagBuilder) {
	return Tag("span").Text(text)
}

// Strong "strong": HTMLElement;
func Strong(text string) (r *HTMLTagBuilder) {
	return Tag("strong").Text(text)
}

// Style "style": HTMLStyleElement;
func Style(style string) (r *HTMLTagBuilder) {
	return Tag("style").
		Attr("type", "text/css").
		Children(RawHTML(style))
}

// Sub "sub": HTMLElement;
func Sub(text string) (r *HTMLTagBuilder) {
	return Tag("sub").Text(text)
}

// Summary "summary": HTMLElement;
func Summary(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("summary", children...)
}

// Sup "sup": HTMLElement;
func Sup(text string) (r *HTMLTagBuilder) {
	return Tag("sup").Text(text)
}

// Table "table": HTMLTableElement;
func Table(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("table", children...)
}

// Tbody "tbody": HTMLTableSectionElement;
func Tbody(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tbody", children...)
}

// Td "td": HTMLTableDataCellElement;
func Td(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("td", children...)
}

// Template "template": HTMLTemplateElement;
func Template(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("template", children...)
}

// Textarea "textarea": HTMLTextAreaElement;
func Textarea(text string) (r *HTMLTagBuilder) {
	return Tag("textarea").Text(text)
}

// Tfoot "tfoot": HTMLTableSectionElement;
func Tfoot(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tfoot", children...)
}

// Th "th": HTMLTableHeaderCellElement;
func Th(text string) (r *HTMLTagBuilder) {
	return Tag("th").Text(text)
}

// Thead "thead": HTMLTableSectionElement;
func Thead(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("thead", children...)
}

// Time "time": HTMLTimeElement;
func Time(datetime string) (r *HTMLTagBuilder) {
	return Tag("time").Attr("datetime", datetime)
}

// Title "title": HTMLTitleElement;
func Title(text string) (r *HTMLTagBuilder) {
	return Tag("title").Text(text)
}

// Tr "tr": HTMLTableRowElement;
func Tr(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tr", children...)
}

// Track "track": HTMLTrackElement;
func Track(src string) (r *HTMLTagBuilder) {
	return Tag("track").OmitEndTag().Attr("src", src)
}

// U "u": HTMLElement;
func U(text string) (r *HTMLTagBuilder) {
	return Tag("u").Text(text)
}

// Ul "ul": HTMLUListElement;
func Ul(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ul", children...)
}

// Var "var": HTMLElement;
func Var(text string) (r *HTMLTagBuilder) {
	return Tag("var").Text(text)
}

// Video "video": HTMLVideoElement;
func Video(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("video", children...)
}

// Wbr "wbr": HTMLElement;
func Wbr() (r *HTMLTagBuilder) {
	return Tag("wbr").OmitEndTag()
}

// Svg "svg": HTMLElement;
func Svg(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("Svg", children...)
}
