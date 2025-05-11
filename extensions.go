package revmd

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type mainRendererExt struct {
}

var MainRendererExt = &mainRendererExt{}

func (e *mainRendererExt) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewRenderer(), 1),
	))
}
