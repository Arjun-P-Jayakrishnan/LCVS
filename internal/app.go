package internal

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

// LayoutFunc refers to the layout logic to be added later on
type LayoutFunc func(gtx layout.Context) layout.Dimensions

// App refers to state holding the lifecycle and rendering context
type App struct {
	window   *app.Window
	layoutFn LayoutFunc
	ops      op.Ops
}

// Create a NewApp Instance
func NewApp(layoutFn LayoutFunc) *App {

	win := new(app.Window)
	win.Option(app.Title("LCVS"))
	win.Option(app.Size(unit.Dp(1200), unit.Dp(700)))

	return &App{
		window:   win,
		layoutFn: layoutFn,
	}
}

// Run handles the system level event loop and handles frame rendering and app shutdown
func (a *App) Run() {

	go func() {
		if err := a.loop(); err != nil {
			log.Printf("Fatal Error : %v \n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Main()
}

// Internal Loop function to handle system level events
func (a *App) loop() error {

	//listen for events happening on the window.
	for {
		e := a.window.Event()

		switch evt := e.(type) {

		//	Frame Event: when frame is to be re-rendered
		case app.FrameEvent:
			//gtx is used to pass around rendering and event information.
			gtx := app.NewContext(&a.ops, evt)
			//call layout function as defined from external
			a.layoutFn(gtx)
			//render and handle the operations from the UI.
			evt.Frame(gtx.Ops)

		//	This is sent when the application is closed
		case app.DestroyEvent:
			return evt.Err
		}

	}
}
