package gui

import (
	"fmt"
	"time"
)

type FrameRateWidget struct {
	TextLine
	frame_times []int64
}

func MakeFrameRateWidget() *FrameRateWidget {
	var w FrameRateWidget
	w.TextLine = *MakeTextLine("standard", "0", 100, 1, 1, 1, 1)
	w.EmbeddedWidget = &BasicWidget{CoreWidget: &w}
	return &w
}

func (w *FrameRateWidget) DoThink(t int64, _ bool) {
	now := time.Now().UnixNano()
	w.frame_times = append(w.frame_times, now)
	prev := now - 1e9
	index := 0
	for w.frame_times[index] < prev {
		index++
	}
	rate := len(w.frame_times) - index + 1
	w.SetText(fmt.Sprintf("%d", rate))
	w.TextLine.DoThink(t, false)
}
