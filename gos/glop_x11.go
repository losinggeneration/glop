// +build !sdl

package gos

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"time"
	"C"
)

var (
	display xgb.Conn
	screen *xproto.ScreenInfo
)

type GlopKeyEvent struct {
	index int;
	device int;
	press_amt float64;
	timestamp time.Time;
	cursor_x, cursor_y int;
	num_lock, caps_lock int;
}

func (event *GlopKeyEvent) clearKeyEvent() {
	event.index = 0;
	event.device = 0;
	event.press_amt = 0;
	event.timestamp = time.Time{}
	event.cursor_x = 0;
	event.cursor_y = 0;
	event.num_lock = 0;
	event.caps_lock = 0;
}

//export GlopInit
func GlopInit() {
	display, err := xgb.NewConnDisplay("");
	if err != nil {
		panic(err)
	}

	setupInfo := xproto.Setup(display)
	screen = setupInfo.DefaultScreen(display);
	// Create input context?
}

//export GlopCreateWindow
func GlopCreateWindow(title string, x, y, width, height int) interface{} {
	return nil
}

//export GlopThink
func GlopThink() {
}

//export GlopSwapBuffers
func GlopSwapBuffers() {
}

//export GlopGetMousePosition
func GlopGetMousePosition(x, y *int) {
}

//export GlopGetWindowDims
func GlopGetWindowDims(x, y, dx, dy *int) {
}

//export GlopGetInputEvents
func GlopGetInputEvents(events *interface{}, num_events interface{}, horizon interface{}) {
}

//export GlopEnableVSync
func GlopEnableVSync(enable bool) {
}


