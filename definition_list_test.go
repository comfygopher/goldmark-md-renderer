package mdrend

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	ext "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func TestDefinitionListRenderer(t *testing.T) {
	markdown := `Term 1
: Definition 1

Term 2
: Definition 2a
: Definition 2b

Term 3
Term 4
: Definition 3-4
`

	// Create our Markdown renderer
	mdRenderer := NewRenderer().(*Renderer)

	// Register the definition list renderers manually
	mdRenderer.nodeRendererFuncMap[ext.KindDefinitionList] = mdRenderer.renderDefinitionList
	mdRenderer.nodeRendererFuncMap[ext.KindDefinitionTerm] = mdRenderer.renderDefinitionTerm
	mdRenderer.nodeRendererFuncMap[ext.KindDefinitionDescription] = mdRenderer.renderDefinitionDescription

	// Create a custom renderer that only uses our Markdown renderer
	customRenderer := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(mdRenderer, 1000),
		),
	)

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.DefinitionList,
		),
		goldmark.WithRenderer(customRenderer),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		t.Fatalf("Failed to convert markdown: %v", err)
	}

	t.Logf("Rendered markdown:\n%s", buf.String())
}
