package p2p

import "errors"

//ErrInvalidHandshake is returned if  the handshake between the local 
// and the remote node could not be established 
var ErrInvalidHandshake = errors.New("invalid handshake")

// Handshake func is any
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }