package ui

import (
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

var progress float32
var progressIncrementer chan float32
var isBoiling bool

func RunApp() {

	go createWindow()

	createChannelForBar()
	app.Main()
}

func createWindow() {

	window := new(app.Window)
	window.Option(app.Title(""))
	window.Option(app.Size(unit.Dp(400), unit.Dp(600)))

	err := draw(window)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)

}

func draw(window *app.Window) error {

	//ops are operations from ui
	var ops op.Ops

	// button is a clickable type
	var startButton widget.Clickable

	// defines the style and theme
	theme := material.NewTheme()

  getFromChannel(window)

	// Listen for events in window
	for {

		//grab the event
		evt := window.Event()

		//decide the type
		switch evntTyp := evt.(type) {

		//this is sent when application should exit
		case app.DestroyEvent:
			return evntTyp.Err

		// this is sent when application should re-render
		case app.FrameEvent:

			gtx := app.NewContext(&ops, evntTyp)

			if startButton.Clicked(gtx) {
				isBoiling = !isBoiling
			}

			frame := createLayout(gtx, theme, &startButton)

			evntTyp.Frame(frame)

		}
	}

}

func createLayout(gtx C, theme *material.Theme, startButton *widget.Clickable) *op.Ops {

	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,

		layout.Rigid(
			func(gtx C) D {
				bar := material.ProgressBar(theme, progress)
				return bar.Layout(gtx)
			},
		),
		layout.Rigid(

			func(gtx C) D {

				//define margins around the button
				margin := layout.Inset{
					Top:    unit.Dp(35),
					Bottom: unit.Dp(35),
					Left:   unit.Dp(35),
					Right:  unit.Dp(35),
				}

				//layout
				return margin.Layout(gtx,
					//
					func(gtx C) D {

						var text string

						if !isBoiling {
							text = "Start"
						} else {
							text = "Stop"
						}
						//
						btn := material.Button(theme, startButton, text)
						return btn.Layout(gtx)
					},
				)

			},
		),

		layout.Rigid(
			layout.Spacer{Height: unit.Dp(25)}.Layout,
		),
	)

	return gtx.Ops
}

func createChannelForBar() {
	progressIncrementer = make(chan float32)

	go func() {
		for {
			time.Sleep(time.Second / 25)
			progressIncrementer <- 0.004
		}
	}()
}

func getFromChannel(window *app.Window) {
	go func() {
		for p := range progressIncrementer {
        
      if isBoiling && progress <1 {
          progress+=p
          window.Invalidate()
      }
		}

	}()

}
