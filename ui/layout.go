package ui

import (
	//"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Widget func(gtx layout.Context) layout.Dimensions

/*
* renders the side pane used for file navigation or ticket
 */
func RenderNavigationPane(globalContext Context, width int) layout.Dimensions {
	gtx := *globalContext.gtx

	size := context.gtx.Constraints.Max
	size.X = size.X / width

	gtx.Constraints.Min = size
	//
	border := widget.Border{
		Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
		CornerRadius: unit.Dp(0),
		Width:        unit.Dp(1),
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		}.Layout(gtx,

			layout.Rigid(func(_gtx layout.Context) layout.Dimensions {

				return material.H3(globalContext.Theme, "Sidepane").Layout(_gtx)

				//return material.H1(globalContext.Theme,"Sidepane").Layout(*globalContext.gtx)
			}),
			layout.Flexed(1, func(_gtx layout.Context) layout.Dimensions {
				//RenderImage("./assets/gamer.png", _gtx.Ops)

				return material.H1(globalContext.Theme, "Sidepane").Layout(_gtx)
			}),
		)

	})

}

/*
Code pane is where all code related viewing is to be done
*/
func RenderCodePane(gtx layout.Context) {

}

/*
Description Pane is where you can add your description of code or
provide a markdown editor to showcase your breakdown of problem
*/
func RenderDescriptionPane(gtx layout.Context) {

}
