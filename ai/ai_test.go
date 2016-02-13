package ai_test

import (
	"github.com/MobRulesGames/glop/ai"
	"github.com/MobRulesGames/polish"
	"github.com/MobRulesGames/yedparse"
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
)

func AiSpec(c gospec.Context) {
	c.Specify("Load a simple .xgml file.", func() {
		g, err := yed.ParseFromFile("state.xgml")
		c.Assume(err, Equals, nil)
		aig := ai.NewGraph()
		aig.Graph = &g.Graph
		aig.Context = polish.MakeContext()
		polish.AddIntMathContext(aig.Context)

		dist := 0
		dist_func := func() int {
			return dist
		}

		var nearest int = 7
		nearest_func := func() int {
			return nearest
		}

		attacks := 0
		attack_func := func() int {
			attacks++
			return 0
		}

		aig.Context.AddFunc("dist", dist_func)
		aig.Context.AddFunc("nearest", nearest_func)
		aig.Context.AddFunc("move", func() int { nearest--; return 0 })
		aig.Context.AddFunc("wait", func() int { return 0 })
		aig.Context.AddFunc("attack", attack_func)
		aig.Eval(2, func() bool { return true })

		c.Expect(attacks, Equals, 0)
		c.Expect(nearest, Equals, 4)
	})
}

func TermSpec(c gospec.Context) {
	g, err := yed.ParseFromFile("state.xgml")
	c.Assume(err, Equals, nil)
	aig := ai.NewGraph()
	aig.Graph = &g.Graph
	aig.Context = polish.MakeContext()
	polish.AddIntMathContext(aig.Context)
	polish.AddIntMathContext(aig.Context)
	c.Specify("Calling AiGraph.Term() will terminate evaluation early.", func() {
		var nearest int = 7
		nearest_func := func() int {
			return nearest
		}

		dist := 0
		term := true
		dist_func := func() int {
			if nearest == 6 && term {
				aig.Term() <- nil
			}
			return dist
		}

		attacks := 0
		attack_func := func() int {
			attacks++
			return 0
		}

		aig.Context.AddFunc("dist", dist_func)
		aig.Context.AddFunc("nearest", nearest_func)
		aig.Context.AddFunc("move", func() int { nearest--; return 0 })
		aig.Context.AddFunc("wait", func() int { return 0 })
		aig.Context.AddFunc("attack", attack_func)
		aig.Eval(2, func() bool { return true })

		c.Expect(attacks, Equals, 0)
		c.Expect(nearest, Equals, 6)

		term = false
		aig.Eval(2, func() bool { return true })
		c.Expect(nearest, Equals, 4)
	})
}

func ChunkSpec(c gospec.Context) {
	g, err := yed.ParseFromFile("state.xgml")
	c.Assume(err, Equals, nil)
	aig := ai.NewGraph()
	aig.Graph = &g.Graph
	aig.Context = polish.MakeContext()
	polish.AddIntMathContext(aig.Context)
	polish.AddIntMathContext(aig.Context)
	c.Specify("cont() returning false will terminate evaluation early.", func() {
		var nearest int = 7
		nearest_func := func() int {
			return nearest
		}

		dist := 0
		term := true
		dist_func := func() int {
			if nearest == 6 && term {
				aig.Term() <- nil
			}
			return dist
		}

		attacks := 0
		attack_func := func() int {
			attacks++
			return 0
		}

		aig.Context.AddFunc("dist", dist_func)
		aig.Context.AddFunc("nearest", nearest_func)
		aig.Context.AddFunc("move", func() int { nearest--; return 0 })
		aig.Context.AddFunc("wait", func() int { return 0 })
		aig.Context.AddFunc("attack", attack_func)
		_, err := aig.Eval(4, func() bool { return false })
		// Only have time for 1 move before we terminate early
		c.Expect(err, Equals, nil)
		c.Expect(nearest, Equals, 6)
	})
}
