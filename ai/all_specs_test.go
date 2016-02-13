package ai_test

import (
	"testing"

	"github.com/orfjackal/gospec/src/gospec"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(AiSpec)
	r.AddSpec(TermSpec)
	r.AddSpec(ChunkSpec)
	gospec.MainGoTest(r, t)
}
