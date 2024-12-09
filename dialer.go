package x

import "net"

type Dialer net.Dialer

var DefaultDialer = net.Dial
