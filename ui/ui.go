package ui

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func RunApp() {

  go createWindow()
  app.Main()
}

func createWindow() {

  window := new(app.Window)
  window.Option(app.Title(""))
  window.Option(app.Size(unit.Dp(400),unit.Dp(600)))


  err:= run(window)

  if err!=nil {
  log.Fatal(err)
  }

  os.Exit(0)

}


func run(window *app.Window) error {

 

  //ops are operations from ui
  var ops op.Ops

  // button is a clickable type
  var startButton widget.Clickable

  // defines the style and theme
  theme:= material.NewTheme()


  // Listen for events in window
  for {
      

    //grab the event
      evt := window.Event()
      
      //decide the type
      switch typ:= evt.(type) {
      
        //this is sent when application should exit
      case app.DestroyEvent:
        return typ.Err

      // this is sent when application should re-render
      case app.FrameEvent:

        gtx := app.NewContext(&ops,typ)
      

        layout.Flex{
          Axis:layout.Vertical,
          Spacing: layout.SpaceStart,
        }.Layout(gtx,

          layout.Rigid(
              func (gtx layout.Context) layout.Dimensions{  
                  btn := material.Button(theme,&startButton,"start")
                  return btn.Layout(gtx)
              },
          ),


        layout.Rigid(
          layout.Spacer{Height:unit.Dp(25)}.Layout,
        ),

        )

        typ.Frame(gtx.Ops)


      }
  }

}
