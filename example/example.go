package main

import (
  "glop/gos"
  "glop/gui"
  "glop/gin"
  "glop/system"
  "runtime"
  "fmt"
  "os"
  "flag"
)

var (
  sys system.System
  font_path *string = flag.String("font", "../../fonts/skia.ttf", "relative path of a font")
  sprite_path *string=flag.String("sprite", "../../sprites/test_sprite", "relative path of sprite")
)

func init() {
  sys = system.Make(gos.GetSystemInterface())
}


type Foo struct {
  *gui.BoxWidget
}
func (f *Foo) HandleEventGroup(group gin.EventGroup) (bool, bool, *gui.Node) {
  // TODO: 304!!!!!!!!!!!!!!!!!
  if found,event := group.FindEvent(304); found {
    if c := event.Key.Cursor(); c != nil && c.Name() == "Mouse"  && event.Type == gin.Press {
      f.Dims.Dx += 5
    }
  }
  return false, false, nil
}

func (f *Foo) Think(_ int64, has_focus bool, previous gui.Region, _ map[gui.Widget]gui.Dims) (bool, gui.Dims) {
  cx,cy := sys.GetCursorPos()
  cursor := gui.Point{ X : cx, Y : cy }
  if cursor.Inside(previous) {
    f.G = 0
    f.B = 0
  } else {
    f.G = 1
    f.B = 1
  }
  return f.BoxWidget.Think(0, false, previous, nil)
}

func main() {
  runtime.LockOSThread()

  // Running a binary via osx's package mechanism will add a flag that begins with
  // '-psn' so we have to find it and pretend like we were expecting it so that go
  // doesn't asplode because of an unexpected flag.
  for _,arg := range os.Args {
    if len(arg) >= 4 && arg[0:4] == "-psn" {
      flag.Bool(arg[1:], false, "HERE JUST TO APPEASE GO'S FLAG PACKAGE")
    }
  }
  flag.Parse()
  sys.Startup()

  err := gui.LoadFont("standard", *font_path)
  if err != nil {
    panic(err.String())
  }

  factor := 0.875
  wdx := int(factor * float64(1024))
  wdy := int(factor * float64(768))

  sys.CreateWindow(10, 10, wdx, wdy)
  ui := gui.Make(sys.Input(), wdx, wdy)
  anch := ui.Root.InstallWidget(gui.MakeAnchorBox(gui.Dims{wdx - 50, wdy - 50}), nil)
  manch := anch.InstallWidget(gui.MakeAnchorBox(gui.Dims{wdx - 150, wdy - 150}), gui.Anchor{1,1,1,1})
  manch.InstallWidget(&Foo{gui.MakeBoxWidget(100, 100, 1,1,1,1)}, gui.Anchor{1,0,1,0})
  text_widget := gui.MakeSingleLineText("standard", "Funk Monkey 7$", 1,0.9,0.9,1)
  v := 1000.0 / 8.0
  terrain,err := gui.MakeTerrain("../../maps/chess.jpg", int(v), 8, 8, 65)
  if err != nil {
    panic(err.String())
  } else {
    manch.InstallWidget(terrain, gui.Anchor{0,0,0,0})
  }

  table := anch.InstallWidget(&gui.VerticalTable{}, gui.Anchor{0,0, 0,0})

  frame_count_widget := gui.MakeSingleLineText("standard", "Frame", 1,0,1,1)
  table.InstallWidget(frame_count_widget, nil)
  table.InstallWidget(text_widget, nil)
  n := 0
  sys.EnableVSync(true)
//  ticker := time.Tick(3e7)
  for {
    n++
    cx,cy := sys.GetCursorPos()
    text_widget.SetText(fmt.Sprintf("Cursor pos: %d %d\n", cx, cy))
    frame_count_widget.SetText(fmt.Sprintf("               %d", n/10))
    sys.Think()
    sys.SwapBuffers()
    groups := sys.GetInputEvents()
    for _,group := range groups {
      if found,_ := group.FindEvent('q'); found {
        return
      }
    }
    kw := sys.Input().GetKey('w')
    ka := sys.Input().GetKey('a')
    ks := sys.Input().GetKey('s')
    kd := sys.Input().GetKey('d')
    m_factor := 0.0075
    dx := m_factor * (kd.FramePressSum() - ka.FramePressSum())
    dy := m_factor * (kw.FramePressSum() - ks.FramePressSum())
    terrain.Move(dx, dy)
    zoom := sys.Input().GetKey('r').FramePressSum() - sys.Input().GetKey('f').FramePressSum()
    terrain.Zoom(zoom * 0.0025)
  }

  fmt.Printf("")
}
