package x

import "net"

type Dialer func(network, address string) (net.Conn, error)

var DefaultDial Dialer = net.Dial
