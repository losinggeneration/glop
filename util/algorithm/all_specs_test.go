package algorithm_test

import (
	"testing"

	"github.com/orfjackal/gospec/src/gospec"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.AddSpec(DijkstraSpec)
	r.AddSpec(ReachableSpec)
	r.AddSpec(ReachableDestinationsSpec)
	r.AddSpec(ChooserSpec)
	r.AddSpec(Chooser2Spec)
	r.AddSpec(MapperSpec)
	r.AddSpec(Mapper2Spec)
	r.AddSpec(TopoSpec)
	gospec.MainGoTest(r, t)
}
