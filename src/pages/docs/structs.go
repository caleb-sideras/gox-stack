package docs

import "github.com/caleb-sideras/gox/src/global"

type DocsData struct {
	ActiveTabId  string
	ActiveDocsId string
	LargeCards   []global.LargeCard
	HoverCards   []global.HoverCard
}
