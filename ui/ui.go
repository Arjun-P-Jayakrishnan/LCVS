package ui

import (
	//"image"
	//"image/color"
	"gioui.org/layout"
	"github.com/Arjun-P-Jayakrishnan/LCVS/ui/elements"
	//"gioui.org/unit"
	"gioui.org/widget/material"
	//"fmt"
	//"gioui.org/widget" // widget contains state for different widgets
)

type Context struct {
	Theme *material.Theme
	gtx   layout.Context
}

var context = Context{
	Theme: material.NewTheme(),
}

/*
Logic for the rendering is written here
*/
func Layout(gtx *layout.Context) layout.Dimensions {
	context.gtx = *gtx

	return elements.Spacer(
		elements.SpacerProps{
			Gtx:   context.gtx,
			Inset: layout.UniformInset(0),
			Child: func(spacerContext layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Alignment: layout.Middle,
					Spacing:   layout.SpaceEnd,
				}.Layout(spacerContext,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Max.X = gtx.Constraints.Max.X / 6

						return Sidepane(SidepaneProps{
							gtx:   gtx,
							theme: *context.Theme,
							child: nil,
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						//gtx.Constraints.Min.X = gtx.Constraints.Max.X / 6

						return material.Body1(context.Theme,"").Layout(gtx)
					}),
					// layout.Flexed(1,func(gtx layout.Context) layout.Dimensions {
					// 	return RenderNavigationPane(context)
					// }),
					// layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					// 	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					// 		return inset.Layout(gtx, material.Editor(context.Theme, &editor, "").Layout)
					// 	})
					// }),
					// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// 	line, col := editor.CaretPos()
					// 	s := fmt.Sprintf("line:%d col:%d", line, col)
					// 	return Center(material.Body1(context.Theme, s)).Layout(gtx)
					// }),
				)

			}})

}
