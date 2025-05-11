package mdrend

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func TestRenderer(t *testing.T) {
	markdown := `# Heading 1

## Heading 2

This is a paragraph with **bold** and *italic* text.

- List item 1
- List item 2
  - Nested list item

1. Ordered list item 1
2. Ordered list item 2

> This is a blockquote

` + "```go" + `
func main() {
	fmt.Println("Hello, world!")
}
` + "```" + `

[Link text](https://example.com "Title")

![Image alt text](https://example.com/image.jpg "Image title")

---

<div>Raw HTML</div>

`

	md := goldmark.New(
		goldmark.WithRenderer(
			renderer.NewRenderer(
				renderer.WithNodeRenderers(
					util.Prioritized(NewRenderer(), 1000),
				),
			),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		t.Fatalf("Failed to convert markdown: %v", err)
	}

	t.Logf("Rendered markdown:\n%s", buf.String())
}
