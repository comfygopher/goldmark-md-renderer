package revmd

import (
	"bytes"
	"github.com/yuin/goldmark/testutil"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.DefinitionList,
			DefinitionListExt,
		),
	)

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
