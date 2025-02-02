package main

import (
	"github.com/Arjun-P-Jayakrishnan/LCVS/internal"
	"github.com/Arjun-P-Jayakrishnan/LCVS/ui"
)


func main() {

  internal.AppUI.RunApp(func(gtx internal.Context) error {
 
    ui.RenderNavigationPane(*gtx)
    
    return nil
  })
  
}
