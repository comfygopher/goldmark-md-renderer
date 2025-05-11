package revmd

import (
	"bytes"
	"github.com/yuin/goldmark/testutil"
	"testing"

	"github.com/yuin/goldmark"
)

func TestRenderer(t *testing.T) {
	markdown := `# Heading 1
## Heading 2
This is a paragraph with **bold** and *italic* text.

- List item 1
- List item 2
  - Nested list item
    - Nested 2

1. Ordered list item 1
1. Ordered list item 2

> This is a blockquote

` + "```go" + `
func main() {
    fmt.Println("Hello, world!")
}
` + "```" + `

[Link text](https://example.com "Title")

![Image alt text](https://example.com/image.jpg "Image title")

---

<div>Raw *HTML*</div>
`

	md := goldmark.New(goldmark.WithExtensions(MainRendererExt))

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		t.Fatalf("Failed to convert markdown: %v", err)
	}

	if buf.String() != markdown {
		t.Errorf(
			"Rendered markdown does not match input:\n%s",
			testutil.DiffPretty([]byte(markdown), buf.Bytes()),
		)
	}
}
