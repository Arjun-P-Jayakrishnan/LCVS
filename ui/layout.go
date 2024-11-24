package ui

import (
	"gioui.org/layout"
)

/*
* renders the side pane used for file navigation or ticket
 */
func RenderNavigationPane(gtx layout.Context) {
	RenderImage("./assets/gamer.png",gtx.Ops)
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
