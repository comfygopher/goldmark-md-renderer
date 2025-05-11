package revmd

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Renderer struct {
	nodeRendererFuncMap map[ast.NodeKind]renderer.NodeRendererFunc
}

// NewRenderer initialize Renderer as renderer.NodeRenderer.
func NewRenderer() renderer.NodeRenderer {
	r := &Renderer{
		nodeRendererFuncMap: map[ast.NodeKind]renderer.NodeRendererFunc{},
	}
	return r
}

// RegisterFuncs add AST objects to Renderer.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// blocks
	reg.Register(ast.KindDocument, r.renderDocument)
	r.nodeRendererFuncMap[ast.KindDocument] = r.renderDocument

	reg.Register(ast.KindHeading, r.heading)
	r.nodeRendererFuncMap[ast.KindHeading] = r.heading

	reg.Register(ast.KindBlockquote, r.renderBlockquote)
	r.nodeRendererFuncMap[ast.KindBlockquote] = r.renderBlockquote

	reg.Register(ast.KindCodeBlock, r.renderCodeBlock)
	r.nodeRendererFuncMap[ast.KindCodeBlock] = r.renderCodeBlock

	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
	r.nodeRendererFuncMap[ast.KindFencedCodeBlock] = r.renderFencedCodeBlock

	reg.Register(ast.KindHTMLBlock, r.renderHTMLBlock)
	r.nodeRendererFuncMap[ast.KindHTMLBlock] = r.renderHTMLBlock

	reg.Register(ast.KindList, r.renderList)
	r.nodeRendererFuncMap[ast.KindList] = r.renderList

	reg.Register(ast.KindListItem, r.renderListItem)
	r.nodeRendererFuncMap[ast.KindListItem] = r.renderListItem

	reg.Register(ast.KindParagraph, r.renderParagraph)
	r.nodeRendererFuncMap[ast.KindParagraph] = r.renderParagraph

	reg.Register(ast.KindTextBlock, r.renderTextBlock)
	r.nodeRendererFuncMap[ast.KindTextBlock] = r.renderTextBlock

	reg.Register(ast.KindThematicBreak, r.renderThematicBreak)
	r.nodeRendererFuncMap[ast.KindThematicBreak] = r.renderThematicBreak

	// inlines
	reg.Register(ast.KindAutoLink, r.renderAutoLink)
	r.nodeRendererFuncMap[ast.KindAutoLink] = r.renderAutoLink

	reg.Register(ast.KindCodeSpan, r.renderCodeSpan)
	r.nodeRendererFuncMap[ast.KindCodeSpan] = r.renderCodeSpan

	reg.Register(ast.KindEmphasis, r.renderEmphasis)
	r.nodeRendererFuncMap[ast.KindEmphasis] = r.renderEmphasis

	reg.Register(ast.KindImage, r.renderImage)
	r.nodeRendererFuncMap[ast.KindImage] = r.renderImage

	reg.Register(ast.KindLink, r.renderLink)
	r.nodeRendererFuncMap[ast.KindLink] = r.renderLink

	reg.Register(ast.KindRawHTML, r.renderRawHTML)
	r.nodeRendererFuncMap[ast.KindRawHTML] = r.renderRawHTML

	reg.Register(ast.KindText, r.renderText)
	r.nodeRendererFuncMap[ast.KindText] = r.renderText

	reg.Register(ast.KindString, r.renderString)
	r.nodeRendererFuncMap[ast.KindString] = r.renderString

}

func (r *Renderer) renderDocument(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) heading(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.Heading)
		for i := 0; i < n.Level; i++ {
			_ = w.WriteByte('#')
		}
		_ = w.WriteByte(' ')
	} else {
		_ = w.WriteByte('\n')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderBlockquote(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("> ")
		// Prefix all children with "> " at the beginning of each line
		for c := node.FirstChild(); c != nil; c = c.NextSibling() {
			if c.Kind() == ast.KindParagraph || c.Kind() == ast.KindHeading || c.Kind() == ast.KindBlockquote {
				// These nodes will handle their own line breaks
				continue
			}
			// For other nodes, we need to add the prefix manually
			_ = ast.Walk(c, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
				if entering && n.Kind() == ast.KindText && n.PreviousSibling() != nil {
					// Add blockquote prefix for text nodes that are not the first child
					_, _ = w.WriteString("\n> ")
				}
				return ast.WalkContinue, nil
			})
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.CodeBlock)
	if entering {
		_, _ = w.WriteString("```\n")
		for i := 0; i < n.Lines().Len(); i++ {
			line := n.Lines().At(i)
			_, _ = w.Write(line.Value(source))
		}
	} else {
		_, _ = w.WriteString("```\n\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	if entering {
		_, _ = w.WriteString("```")
		if n.Info != nil {
			_, _ = w.Write(n.Info.Segment.Value(source))
		}
		_ = w.WriteByte('\n')
		for i := 0; i < n.Lines().Len(); i++ {
			line := n.Lines().At(i)
			_, _ = w.Write(line.Value(source))
		}
	} else {
		_, _ = w.WriteString("```\n\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.HTMLBlock)
	if entering {
		for i := 0; i < n.Lines().Len(); i++ {
			line := n.Lines().At(i)
			_, _ = w.Write(line.Value(source))
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		parent := node.Parent()
		if parent != nil && parent.Kind() != ast.KindListItem {
			_, _ = w.WriteString("\n")
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		// Calculate indentation level based on parent lists
		indentLevel := 0
		parent := node.Parent()
		for parent != nil && parent.Kind() == ast.KindList {
			indentLevel++
			parent = parent.Parent().Parent() // Skip the ListItem to get to the next List
		}

		// Apply indentation
		for i := 1; i < indentLevel; i++ {
			_, _ = w.WriteString("  ")
		}

		// Write list marker
		parentList := node.Parent().(*ast.List)
		if parentList.IsOrdered() {
			_, _ = w.WriteString("1. ")
		} else {
			_, _ = w.WriteString("- ")
		}
	} else {
		//if first child is TextBlock then don't add new line:
		firstChild := node.FirstChild()
		if firstChild.Kind() == ast.KindTextBlock && firstChild.NextSibling() != nil && firstChild.FirstChild() != nil {
			return ast.WalkContinue, nil
		}
		_, _ = w.WriteString("\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderParagraph(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString("\n\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
	} else {
		if node.NextSibling() != nil && node.FirstChild() != nil {
			_ = w.WriteByte('\n')
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderThematicBreak(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		// In Markdown, a thematic break (horizontal rule) can be represented by three or more
		// hyphens, asterisks, or underscores, optionally with spaces between them.
		// We'll use three hyphens followed by a newline.
		_, _ = w.WriteString("---\n\n")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderAutoLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.AutoLink)
	_, _ = w.WriteString("<")
	_, _ = w.Write(n.URL(source))
	_, _ = w.WriteString(">")
	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeSpan(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("`")
		for c := node.FirstChild(); c != nil; c = c.NextSibling() {
			if text, ok := c.(*ast.Text); ok {
				_, _ = w.Write(text.Segment.Value(source))
			}
		}
	} else {
		_, _ = w.WriteString("`")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderEmphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Emphasis)
	if entering {
		if n.Level == 1 {
			_, _ = w.WriteString("*")
		} else {
			_, _ = w.WriteString("**")
		}
	} else {
		if n.Level == 1 {
			_, _ = w.WriteString("*")
		} else {
			_, _ = w.WriteString("**")
		}
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*ast.Image)
		_, _ = w.WriteString("![")
		r.renderChildren(w, source, n)
		_, _ = w.WriteString("](")
		_, _ = w.Write(n.Destination)
		if n.Title != nil {
			_, _ = w.WriteString(" \"")
			_, _ = w.Write(n.Title)
			_, _ = w.WriteString("\"")
		}
		_, _ = w.WriteString(")")
		return ast.WalkSkipChildren, nil
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)
	if entering {
		_, _ = w.WriteString("[")
	} else {
		_, _ = w.WriteString("](")
		_, _ = w.Write(n.Destination)
		if n.Title != nil {
			_, _ = w.WriteString(" \"")
			_, _ = w.Write(n.Title)
			_, _ = w.WriteString("\"")
		}
		_, _ = w.WriteString(")")
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderRawHTML(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkSkipChildren, nil
	}
	n := node.(*ast.RawHTML)
	for i := 0; i < n.Segments.Len(); i++ {
		segment := n.Segments.At(i)
		_, _ = w.Write(segment.Value(source))
	}
	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Text)
	segment := n.Segment
	_, _ = w.Write(segment.Value(source))
	if n.HardLineBreak() {
		_, _ = w.WriteString("  \n")
	} else if n.SoftLineBreak() {
		_ = w.WriteByte('\n')
	}
	return ast.WalkContinue, nil
}

func (r *Renderer) renderString(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.String)
	_, _ = w.Write(n.Value)
	return ast.WalkContinue, nil
}

func (r *Renderer) renderChildren(w util.BufWriter, source []byte, node ast.Node) {
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		_ = ast.Walk(c, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
			return r.renderNode(w, source, n, entering)
		})
	}
}

func (r *Renderer) renderNode(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if r, ok := r.nodeRendererFuncMap[node.Kind()]; ok {
		return r(w, source, node, entering)
	}
	// If we don't have a renderer for this node kind, just continue walking
	return ast.WalkContinue, nil
}
