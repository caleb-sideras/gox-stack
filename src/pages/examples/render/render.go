package example_render

import (
	"html/template"

	"github.com/caleb-sideras/gox/.gox/render"
	"github.com/caleb-sideras/gox/src/global"
)

func Markdown_() render.StaticT {
	tmpl := template.Must(global.MarkdownToHTML("pages/examples/_markdown/render_example.md"))

	wrapperTemplateString := `<script src="/static/js/prism.js"></script> <div id="markdown" class="p-4">{{template "code" .}}</div>`
	wrapperTmpl := template.Must(template.New("markdown").Parse(wrapperTemplateString))
	wrapperTmpl.New("code").Parse(tmpl.Tree.Root.String())

	return render.StaticT{wrapperTmpl, nil, ""}
}

func Code_() render.StaticT {
	tmpl := template.Must(global.MarkdownToHTML("pages/examples/_markdown/render.md"))

	wrapperTemplateString := `<script src="/static/js/prism.js"></script> <div id="markdown">{{template "code" .}}</div>`
	wrapperTmpl := template.Must(template.New("markdown").Parse(wrapperTemplateString))
	wrapperTmpl.New("code").Parse(tmpl.Tree.Root.String())

	return render.StaticT{wrapperTmpl, nil, ""}
}
