package ui

import (
	"gioui.org/layout"
)

/**/
type SpacerProps struct{
	//graphical context
	gtx layout.Context;
	//inset the spacing amount
	inset layout.Inset;
	//the widgets to be rendered inside
	child layout.Widget;
}

/*
	Create  a spacer
*/
func Spacer(props SpacerProps) layout.Dimensions{
	return props.inset.Layout(props.gtx,props.child)
}