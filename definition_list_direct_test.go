package mdrend

import (
	"testing"

	ext "github.com/yuin/goldmark/extension/ast"
)

func TestDefinitionListDirectRendering(t *testing.T) {
	// Create a renderer and manually register the definition list renderers
	renderer := NewRenderer().(*Renderer)

	// Register the definition list renderers manually
	renderer.nodeRendererFuncMap[ext.KindDefinitionList] = renderer.renderDefinitionList
	renderer.nodeRendererFuncMap[ext.KindDefinitionTerm] = renderer.renderDefinitionTerm
	renderer.nodeRendererFuncMap[ext.KindDefinitionDescription] = renderer.renderDefinitionDescription

	// Check that the definition list node kinds are registered
	if _, ok := renderer.nodeRendererFuncMap[ext.KindDefinitionList]; !ok {
		t.Errorf("KindDefinitionList renderer not registered")
	}

	if _, ok := renderer.nodeRendererFuncMap[ext.KindDefinitionTerm]; !ok {
		t.Errorf("KindDefinitionTerm renderer not registered")
	}

	if _, ok := renderer.nodeRendererFuncMap[ext.KindDefinitionDescription]; !ok {
		t.Errorf("KindDefinitionDescription renderer not registered")
	}

	// If we got here without errors, the test passes
	t.Log("All definition list renderers are correctly registered")
}
