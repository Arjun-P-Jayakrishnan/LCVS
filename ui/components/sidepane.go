package components

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/Arjun-P-Jayakrishnan/LCVS/ui/elements"
)

type Hero struct {
	heroGtx          layout.Context
	heightPercentage float32
}

type Tile struct {
	Gtx layout.Context
	Theme material.Theme
	Label string
}

type SidepaneProps struct {
	Gtx   layout.Context
	Theme material.Theme
	Hero  Hero
	Tiles []Tile
}

func (props Hero) Init() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Max.Y = int(float32((gtx.Constraints.Max.Y-gtx.Constraints.Min.Y))*props.heightPercentage) + gtx.Constraints.Min.Y

		dr := image.Rectangle{
			Min: gtx.Constraints.Min,
			Max: gtx.Constraints.Max,
		}

		paint.LinearGradientOp{
			Stop1:  layout.FPt(dr.Min),
			Stop2:  layout.FPt(dr.Max),
			Color1: color.NRGBA{R: 0, G: 0xff, B: 0x10, A: 0xFF},
			Color2: color.NRGBA{R: 0x10, G: 0x10, B: 0xff, A: 0xFF},
		}.Add(gtx.Ops)

		defer clip.Rect(dr).Push(gtx.Ops).Pop()

		paint.PaintOp{}.Add(gtx.Ops)

		return layout.Dimensions{
			Size: gtx.Constraints.Max,
		}
	}
}

func (tile Tile) Init() layout.Dimensions {
	
	return elements.Button(elements.ButtonProps{
		Theme: tile.Theme,
		Gtx:   tile.Gtx,
		Label: tile.Label,
	})
}

func Sidepane(props SidepaneProps) layout.Dimensions {


	tiles := make([]layout.FlexChild, len(props.Tiles)+1)

	for i := 0; i < len(props.Tiles)+1; i++ {
		tiles[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max=props.Gtx.Constraints.Max
			tile := Tile{
				Gtx: gtx,
				Theme: props.Theme,
				Label: "Tile" + string(i),
			}
			return tile.Init()
		})
	}

	return elements.Background(elements.BackgroundProps{
		Gtx:             props.Gtx,
		BackgroundColor: color.NRGBA{R: 250, G: 250, B: 250, A: 0xFF},
		Child: func(gtx layout.Context) layout.Dimensions {
			hero := Hero{
				heroGtx: gtx,
				heightPercentage: 0.3,
			}
			tiles[0]= layout.Rigid(hero.Init())

			return layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceEnd,
			}.Layout(gtx, tiles...)
		},
	},
	)
}
