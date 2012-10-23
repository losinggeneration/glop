package gin_test

import (
	"github.com/orfjackal/gospec/src/gospec"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(NaturalKeySpec)
	r.AddSpec(DerivedKeySpec)
	r.AddSpec(NestedDerivedKeySpec)
	r.AddSpec(EventSpec)
	r.AddSpec(EventListenerSpec)
	r.AddSpec(AxisSpec)
	gospec.MainGoTest(r, t)
}
