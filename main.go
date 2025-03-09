package main

import (
	//"log"

	"github.com/Arjun-P-Jayakrishnan/LCVS/internal"
	//"github.com/Arjun-P-Jayakrishnan/LCVS/p2p"
	"github.com/Arjun-P-Jayakrishnan/LCVS/ui"
)


func main() {


	// tcpOpts:=p2p.TCPTransportOpts{
	// 	ListenAddr: ":3000",
	// 	HandshakeFunc: p2p.NOPHandshakeFunc,
	// 	Decoder: p2p.DefaultDecoder{},

	// }
	// tr:=p2p.NewTCPTransport(tcpOpts)
	// if err:=tr.ListenAndAccept();err !=nil{
	// 	log.Fatal(err)
	// }
	
	// select{}
 internal.AppUI.RunApp(ui.Layout)
  
}
