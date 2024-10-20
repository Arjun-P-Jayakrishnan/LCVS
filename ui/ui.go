package ui

import (
	"image"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type EventFunction struct {
	eventType string
	render    func(ops *op.Ops)
	input     func(ops *op.Ops)
	inputSize func(ops *op.Ops, windowSize image.Point)
	metric    func(ops *op.Ops, metric unit.Metric)
	layout    func(gtx layout.Context) layout.Dimensions
}

// holds all application state
type UI struct {
	Theme *material.Theme // used to hold fonts thoughout the application
}

func NewUI() *UI {

	ui := &UI{}

	//load theme and fonts
	ui.Theme = material.NewTheme(gofont.Collection())

	return ui
}

func (ui *UI) Run(w *app.Window) error {

	e := w.Event()
  var ops op.Ops

	//listen for events happening on the window
	for {

		//detect the type of event
		switch eType := e.(type) {

		// sent when re rendering must happen
		case app.FrameEvent:

			//gtx is used to pass around system info
			gtx := app.NewContext(&ops, eType)

			//handle all ui logic
			ui.Layout(gtx)

			//render and handler operations from ui
			eType.Frame(gtx.Ops)

			//sent when application is closed
		case app.DestroyEvent:
			if eType.Err != nil {
				log.Println(eType.Err)
				os.Exit(1)
			}

			os.Exit(0)

		}

	}

}

//Layout handles rendering and input
func (ui *UI) Layout(gtx layout.Context) error{
    return nil;
}

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

			if isBoiling && progress < 1 {
				progress += p
				window.Invalidate()
			}
		}

	}()

}

//controls the function

func EventRouter(gtx *layout.Context, ef *EventFunction) {

	switch eType := ef.eventType; eType {

	case "Render":
		ef.render(gtx.Ops)

	case "Input":
		ef.input(gtx.Ops)

	case "InputSize":
		ef.inputSize(gtx.Ops, gtx.Constraints.Max)
	case "Metric":
		ef.metric(gtx.Ops, gtx.Metric)
	case "Layout":
		ef.layout(*gtx)
	}
}

// Handler for various types of Events
func EventHandler(ef *EventFunction) {

	// a coroutine to spin of in the main thread
	go func() {
		w := new(app.Window)

		//ops is used to encode different operations
		var ops op.Ops

		e := w.Event()

		//listen for events happening on the window
		for {

			//detect the type of event
			switch eType := e.(type) {

			// sent when re rendering must happen
			case app.FrameEvent:

				//gtx is used to pass around system info
				gtx := app.NewContext(&ops, eType)

				//render content
				EventRouter(&gtx, ef)

				//render and handler operations from ui
				eType.Frame(gtx.Ops)

				//sent when application is closed
			case app.DestroyEvent:
				if eType.Err != nil {
					log.Println(eType.Err)
					os.Exit(1)
				}

				os.Exit(0)

			}

		}
	}()

	app.Main()
}

// Render is a utility to start rendering a gio app
func Render(fn func(ops *op.Ops)) {
	EventHandler(&EventFunction{render: fn})
}

// Input is a utility to start rendering and input gio app
func Input(fn func(ops *op.Ops)) {
	EventHandler(&EventFunction{input: fn})
}

// Input Size is a utility to start rendering and input gio app
func InputSize(fn func(ops *op.Ops, windowSize image.Point)) {
	EventHandler(&EventFunction{inputSize: fn})
}

// Metric is a utility to start rendering input and metric gio app
func Metric(fn func(ops *op.Ops, metric unit.Metric)) {
	EventHandler(&EventFunction{metric: fn})
}

// Layout is utility to start rendering and layout widgets
func Layout(fn func(gtx layout.Context) layout.Dimensions) {
	EventHandler(&EventFunction{layout: fn})
}
