package mdrend

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func TestThematicBreakRenderer(t *testing.T) {
	markdown := `Before the break

---

After the break`

	// Create a custom renderer that only uses our Markdown renderer
	customRenderer := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(NewRenderer(), 1000),
		),
	)

	md := goldmark.New(
		goldmark.WithRenderer(customRenderer),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		t.Fatalf("Failed to convert markdown: %v", err)
	}

	output := buf.String()
	t.Logf("Rendered markdown:\n%s", output)

	// Check that the thematic break is rendered correctly
	if !bytes.Contains(buf.Bytes(), []byte("---")) {
		t.Errorf("Expected output to contain '---', but got: %s", output)
	}
}
