package gin

import (
  "fmt"
)


type Key interface {
  // Human readable name
  String() string

  // Unique Id
  Id() KeyId

  // Sets the instantaneous press amount for this key at a specific time and returns the
  // event generated, if any
  SetPressAmt(amt float64, ms int64, cause Event) Event

  // Returns the Cursor associated with this key, or nil if it has no such association.
  Cursor() Cursor

  Think(ms int64)

  subAggregator
}
type subAggregator interface {
  IsDown() bool
  FramePressCount() int
  FrameReleaseCount() int
  FramePressAmt() float64
  FramePressSum() float64
  FramePressAvg() float64
  CurPressCount() int
  CurReleaseCount() int
  CurPressAmt() float64
  CurPressSum() float64
}
type aggregator interface {
  subAggregator
  Think(ms int64)
  SetPressAmt(amt float64, ms int64, event_type EventType)
}
// Simple struct that aggregates presses and press_amts during a frame so they can be viewed
// between Think()s
type keyStats struct {
  press_count   int
  release_count int
  press_amt     float64
  press_sum     float64
  press_avg     float64
}

type baseAggregator struct {
  this,prev  keyStats
}
func (a *baseAggregator) FramePressCount() int {
  return a.prev.press_count
}
func (a *baseAggregator) FrameReleaseCount() int {
  return a.prev.release_count
}
func (a *baseAggregator) FramePressAmt() float64 {
  return a.prev.press_amt
}
func (a *baseAggregator) FramePressSum() float64 {
  return a.prev.press_sum
}
func (a *baseAggregator) FramePressAvg() float64 {
  return a.prev.press_avg
}
func (a *baseAggregator) CurPressCount() int {
  return a.this.press_count
}
func (a *baseAggregator) CurReleaseCount() int {
  return a.this.release_count
}
func (a *baseAggregator) CurPressAmt() float64 {
  return a.this.press_amt
}
func (a *baseAggregator) CurPressSum() float64 {
  return a.this.press_sum
}
func (a *baseAggregator) handleEventType(event_type EventType) {
  switch event_type {
    case Press:
      a.this.press_count++
    case Release:
      a.this.release_count++
  }
}

// the standardAggregator's sum is an integral of the press_amt over time
type standardAggregator struct {
  baseAggregator
  last_press int64
  last_think int64
}
func (sa *standardAggregator) IsDown() bool {
  return sa.this.press_amt != 0
}
func (sa *standardAggregator) SetPressAmt(amt float64, ms int64, event_type EventType) {
  sa.this.press_sum += sa.this.press_amt * float64(ms - sa.last_press)
  sa.this.press_amt = amt
  sa.last_press = ms
  sa.handleEventType(event_type)
}
func (sa *standardAggregator) Think(ms int64) {
  sa.this.press_sum += sa.this.press_amt * float64(ms - sa.last_press)
  sa.this.press_avg = sa.this.press_sum / float64(ms - sa.last_think)
  sa.prev = sa.this
  sa.this = keyStats{
    press_amt : sa.prev.press_amt,
  }
  sa.last_press = ms
  sa.last_think = ms
}

// The axisAggregator's sum is the sum of all press amounts specified by SetPressAmt()
// FramePressAvg() returns the same value as FramePressSum()
type axisAggregator struct {
  baseAggregator
  is_down    bool
}
func (aa *axisAggregator) IsDown() bool {
  return aa.is_down
}
func (aa *axisAggregator) SetPressAmt(amt float64, ms int64, event_type EventType) {
  aa.this.press_sum += amt
  aa.this.press_amt = amt
  if amt != 0 {
    aa.is_down = true
  }
  aa.handleEventType(event_type)
}
func (aa *axisAggregator) Think(ms int64) {
  aa.prev = aa.this
  aa.this = keyStats{}
  if aa.prev.press_amt == 0 {
    aa.is_down = false
  }
  aa.prev.press_avg = aa.prev.press_sum
}


type KeyId int

// natural keys and derived keys all embed a keyState
type keyState struct {
  id     KeyId   // Unique id among all keys ever
  name   string  // Human readable name for the key, 'Right Shift', 'q', 'Space Bar', etc...
  cursor *cursor // cursor associated with this key, or nil if it has no cursor association
  aggregator
}

func (ks *keyState) String() string {
  return fmt.Sprintf("%d: %s", ks.id, ks.name)
}

func (ks *keyState) Id() KeyId {
  return ks.id
}

func (ks *keyState) Cursor() Cursor {
  if ks.cursor == nil {
    return nil
  }
  return ks.cursor
}


// Tells this key that how much it was pressed at a particular time.  Times must be
// monotonically increasing.
// If this press was caused by another event (as is the case with derived keys), then
// cause is the event that made this happen.
func (ks *keyState) SetPressAmt(amt float64, ms int64, cause Event) (event Event) {
  event.Type = NoEvent
  event.Key = ks
  if (ks.CurPressAmt() == 0) != (amt == 0) {
    if amt == 0 {
      event.Type = Release
    } else {
      event.Type = Press
    }
  } else {
    if ks.CurPressAmt() != 0 && ks.CurPressAmt() != amt {
      event.Type = Adjust
    }
  }
  ks.aggregator.SetPressAmt(amt, ms, event.Type)
  return
}