package stats_test

import (
  . "gospec"
  "gospec"
  "game/stats"
  "game/base"
)

func StatsSpec(c gospec.Context) {
  b := stats.BaseStats{
    DynamicStats : stats.DynamicStats{
      Health : 10,
      Ap : 10,
    },
    Attack : 10,
    Defense : 10,
    LosDist : 10,
    Atts : []string{
      "basic",
      "tag",
    },
  }
  attmap := map[string]stats.Attributes{
    "basic" : stats.Attributes{
    },
    "tag" : stats.Attributes{
    },
  }
  c.Specify("Dynamic stats stuff", func() {
    stat := stats.MakeStats(b, attmap)
    c.Expect(stat.BaseAp(), Equals, 10)
    c.Expect(stat.BaseHealth(), Equals, 10)
    c.Expect(stat.CurAp(), Equals, 10)
    c.Expect(stat.CurHealth(), Equals, 10)

    // Make sure that the damage is reflected in the current health, not the
    // base health
    stat.DoDamage(3)
    c.Expect(stat.BaseHealth(), Equals, 10)
    c.Expect(stat.CurHealth(), Equals, 7)
    c.Expect(stat.BaseAp(), Equals, 10)
    c.Expect(stat.CurAp(), Equals, 10)

    // Make sure that the health is not regained at the end of a round
    stat.Round()
    c.Expect(stat.BaseHealth(), Equals, 10)
    c.Expect(stat.CurHealth(), Equals, 7)
    c.Expect(stat.BaseAp(), Equals, 10)
    c.Expect(stat.CurAp(), Equals, 10)

    // Make sure that Ap spending is reflected in the cur stat, not the base
    stat.SpendAp(5)
    c.Expect(stat.BaseHealth(), Equals, 10)
    c.Expect(stat.CurHealth(), Equals, 7)
    c.Expect(stat.BaseAp(), Equals, 10)
    c.Expect(stat.CurAp(), Equals, 5)

    // Make sure that Ap is replenished at the end of a round
    stat.Round()
    c.Expect(stat.BaseHealth(), Equals, 10)
    c.Expect(stat.CurHealth(), Equals, 7)
    c.Expect(stat.BaseAp(), Equals, 10)
    c.Expect(stat.CurAp(), Equals, 10)
  })
}

func EffectsSpec(c gospec.Context) {
  b := stats.BaseStats{
    DynamicStats : stats.DynamicStats{
      Health : 10,
      Ap : 10,
    },
    Attack : 10,
    Defense : 10,
    LosDist : 10,
    Atts : []string{
      "basic",
      "tag",
    },
  }
  attmap := map[string]stats.Attributes{
    "basic" : stats.Attributes{
      MoveMods : map[base.Terrain]int{
        "grass" : 0,
        "hills" : 2,
      },
    },
    "tag" : stats.Attributes{
    },
  }
  c.Specify("Movement cost is 1 + whatever is in the MoveMods", func() {
    stat := stats.MakeStats(b, attmap)
    c.Expect(stat.MoveCost("grass"), Equals, 1)
    c.Expect(stat.MoveCost("hills"), Equals, 3)
  })
  c.Specify("Movement can be affected by effects", func() {
    stat := stats.MakeStats(b, attmap)
    stat.AddEffect(&stats.Slow{ TimedEffect : 1 })
    c.Expect(stat.MoveCost("grass"), Equals, 2)
    c.Expect(stat.MoveCost("hills"), Equals, 4)
  })
}
