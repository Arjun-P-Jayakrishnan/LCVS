package ui

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/Arjun-P-Jayakrishnan/LCVS/ui/components"
	"github.com/Arjun-P-Jayakrishnan/LCVS/ui/elements"
)

// Context is for the global state of the application
type Context struct {
	//The theme of the entire application
	Theme *material.Theme
	//Graphical context
	gtx *layout.Context
}

// Creates a new Context whenever the new window is created
func NewContext(gtx *layout.Context) *Context {
	return &Context{
		Theme: material.NewTheme(),
		gtx:   gtx,
	}
}

// Allows user to create the application widgets using the graphical context
func Layout(gtx *layout.Context) layout.Dimensions {
	context := NewContext(gtx)


	

	return elements.Background(elements.BackgroundProps{
		Gtx:             *gtx,
		BackgroundColor: color.NRGBA{R:225,G:225,B:225,A:255},
		Child: func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceEnd,
			}.Layout(gtx,
				//Sidepane
				layout.Rigid(func(_gtx layout.Context) layout.Dimensions {
					_gtx.Constraints.Max.X = gtx.Constraints.Max.X / 6

					return components.Sidepane(components.SidepaneProps{
						Gtx:   _gtx,
						Theme: *context.Theme,
						Hero: components.Hero{
				
						},
						Tiles: make([]components.Tile, 1),
					})
				}),
				//Main Area
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					//gtx.Constraints.Min.X = gtx.Constraints.Max.X / 6

					return material.Body1(context.Theme, "Main").Layout(gtx)
				}),
			)
		},
	})

}
