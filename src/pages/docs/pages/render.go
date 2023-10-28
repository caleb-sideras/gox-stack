package pages

import (
	"html/template"

	"github.com/caleb-sideras/gox-website/gox/render"
	"github.com/caleb-sideras/gox-website/src/global"
	"github.com/caleb-sideras/gox-website/src/pages/docs"
)

func Render() render.DynamicT {
	tmpls := []string{"templates/components/nav.html"}

	markdownDocsContent := docs.DocsData{
		ActiveTabId:  "docs",
		ActiveDocsId: "page",
	}

	tmpl := template.Must(global.MarkdownToHTML("pages/docs/_markdown/pages.md"))

	return render.DynamicT{tmpls, markdownDocsContent, tmpl}
}
