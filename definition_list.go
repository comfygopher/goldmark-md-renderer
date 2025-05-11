package revmd

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	ext "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type DefinitionListRenderer struct {
}

func NewDefinitionListRenderer() renderer.NodeRenderer {
	return &DefinitionListRenderer{}
}

func (r *DefinitionListRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ext.KindDefinitionList, r.renderDefinitionList)
	reg.Register(ext.KindDefinitionTerm, r.renderDefinitionTerm)
	reg.Register(ext.KindDefinitionDescription, r.renderDefinitionDescription)
}

func (r *DefinitionListRenderer) renderDefinitionList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// For definition lists, we don't need to add any special formatting at the list level
	// The terms and descriptions will handle their own formatting
	if !entering {
		_, _ = w.WriteString("\n")
	}
	return ast.WalkContinue, nil
}

func (r *DefinitionListRenderer) renderDefinitionTerm(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// For definition terms, we just render the content and add a newline at the end
	if entering {
		if node.PreviousSibling() != nil && node.PreviousSibling().Kind() == ext.KindDefinitionDescription {
			_, _ = w.WriteString("\n")
		}
	} else {
		_, _ = w.WriteString("\n")
	}
	return ast.WalkContinue, nil
}

func (r *DefinitionListRenderer) renderDefinitionDescription(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// For definition descriptions, we add the colon and space at the beginning
	if entering {
		_, _ = w.WriteString(": ")
	} else {
		_, _ = w.WriteString("\n")
	}
	return ast.WalkContinue, nil
}

type definitionListExt struct {
}

// DefinitionListExt is an mainRendererExt that allow you to use PHP Markdown Extra Definition lists.
var DefinitionListExt = &definitionListExt{}

func (e *definitionListExt) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewDefinitionListRenderer(), 250),
	))
}
