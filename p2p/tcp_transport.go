package p2p

import (
	//"bytes"
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn
	//if we dial and retrieve a conn => outbound == true
	//if we accept and retrieve a conn => outbound  == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr string
	HandshakeFunc HandshakeFunc
	Decoder Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listner       net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}


func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listner, err = net.Listen("tcp", t.ListenAddr)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listner.Accept()
		if err != nil {
			fmt.Printf("TCP accpet error: %s \n", err)
		}
		fmt.Printf("new incomming connection %+v\n",conn)

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {

	peer :=NewTCPPeer(conn,true)

	if err:=t.HandshakeFunc(peer);err!=nil{
		conn.Close()
		return 
	}

	//Read Loop
	msg :=&Message{}
	for {


		if err:=t.Decoder.Decode(conn,msg);err!=nil{
			fmt.Printf("TCP error: %s\n",err)
			continue
		}
		msg.From=conn.RemoteAddr()

		fmt.Printf("message: %+v\n",msg)
	}

	
}
