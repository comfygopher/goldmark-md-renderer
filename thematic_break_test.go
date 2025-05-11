package revmd

import (
	"bytes"
	"github.com/yuin/goldmark/testutil"
	"testing"

	"github.com/yuin/goldmark"
)

func TestThematicBreakRenderer(t *testing.T) {
	markdown := `Before the break

---

After the break

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
