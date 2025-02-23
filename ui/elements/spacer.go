package elements

import (
	"gioui.org/layout"
)

/**/
type SpacerProps struct{
	//graphical context
	Gtx layout.Context;
	//inset the spacing amount
	Inset layout.Inset;
	//the widgets to be rendered inside
	Child layout.Widget;
}

/*
	Create  a spacer
*/
func Spacer(props SpacerProps) layout.Dimensions{
	return props.Inset.Layout(props.Gtx,props.Child)
}