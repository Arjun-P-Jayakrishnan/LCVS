package internal

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	//"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

//Logic for how UI should look like should be written inside
type Layout func() error



type UI struct {
	//theme to be called
	Theme *material.Theme
	//the context for the graphical pipeline
	gtx    layout.Context
	//render function
	renderLayout Layout
}



var (
	titleColor = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
)

var AppUI = UI{
	Theme: material.NewTheme(),
}

/*
	Configure the ui theme
*/
func (ui UI) Configure() error {

	return nil
}

/*
	Running the main app thread
*/
func (ui UI) RunApp(fn Layout) {
	/*
		UI loop is seperated from application window creation
		such that it can be used for testing
	*/
	fmt.Println("App running")
	ui.renderLayout=fn

	go func() {
		w := new(app.Window)
		w.Option(app.Title("LCVS"))
		w.Option(app.Size(unit.Dp(800), unit.Dp(400)))

		err := ui.handleEvents(w)

		if err != nil {
			log.Println(err)
			os.Exit(1)
			panic(err)
		}

		os.Exit(0)
	}()

	app.Main()
}





/*
	Run handles window events and renders the application.
*/
func (ui *UI) handleEvents(w *app.Window) error {

	//op will be used to encode different operations.
	var ops op.Ops

	//listen for events happening on the window.
	for {

		e := w.Event()

		//detect event type
		switch eType := e.(type) {

			/*
				Frame Event: when frame is to be re-rendered
			*/
			case app.FrameEvent:
				//gtx is used to pass around rendering and event information.
				gtx := app.NewContext(&ops, eType)
				//graphics context
				ui.gtx=gtx
				//handle all UI logic.
				ui.renderLayout()
				//render and handle the operations from the UI.
				eType.Frame(gtx.Ops)
			/*
				This is sent when the application is closed
			*/
			case app.DestroyEvent:
				return eType.Err
		}

	}

}


// Title creates a center aligned H1.
func title(th *material.Theme, caption string) material.LabelStyle {
	label := material.H1(th, caption)
	label.Color = titleColor
	label.Alignment = text.Middle
	return label
}

