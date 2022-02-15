package net

// TCPConn is an implementation of the Conn interface for TCP network
// connections.
type TCPConn struct {
	conn
}

func (c *TCPConn) CloseWrite() error {
	return &OpError{"close", "", nil, nil, ErrNotImplemented}
}

// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 scoped addressing zone
}

func ResolveTCPAddr(network, address string) (*TCPAddr, error) {
	panic("unimplemented tcpsock.ResoveTCPAddr()")
}

func (a *TCPAddr) Network() string { return "tcp" }
func (a *TCPAddr) String() string {
	panic("unimplemented: TCPAddr.String()")
}