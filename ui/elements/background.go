package elements

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image/color"
)

type ColorWidgetProps struct {
	//The graphical context
	Gtx layout.Context
	//The color
	BackgroundColor color.NRGBA
	//the child widget
	Child layout.Dimensions
}

// Colors the background
func ColorWidget(props ColorWidgetProps) layout.Dimensions {
	defer clip.Rect{Max: props.Gtx.Constraints.Max}.Push(props.Gtx.Ops).Pop()
	paint.ColorOp{Color: props.BackgroundColor}.Add(props.Gtx.Ops)
	paint.PaintOp{}.Add(props.Gtx.Ops)

	return props.Child
}

type BackgroundProps struct {
	Gtx layout.Context
	BackgroundColor color.NRGBA
	Child layout.Widget
}

/*
	This follows the architecture suggested by gioui itself
	https://gioui.org/doc/architecture/layout
*/
func Background(props BackgroundProps) layout.Dimensions {
	return layout.Background{}.Layout(
		props.Gtx,
		//background coloring
		func(gtx layout.Context) layout.Dimensions {
			defer clip.Rect{Max:gtx.Constraints.Max}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops,props.BackgroundColor)
			return layout.Dimensions{
				Size:gtx.Constraints.Min,
			}
		},
		//layout all the widgets in a stack
		props.Child,
	)
}
