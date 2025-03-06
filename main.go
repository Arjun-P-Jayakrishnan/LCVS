package main

import (
	"log"

	//"github.com/Arjun-P-Jayakrishnan/LCVS/internal"
	"github.com/Arjun-P-Jayakrishnan/LCVS/p2p"
	//"github.com/Arjun-P-Jayakrishnan/LCVS/ui"
)


func main() {
	tr:=p2p.NewTCPTransport(":3000")
	if err:=tr.ListenAndAccept();err !=nil{
		log.Fatal(err)
	}
	
	select{}
 // internal.AppUI.RunApp(ui.Layout)
  
}
