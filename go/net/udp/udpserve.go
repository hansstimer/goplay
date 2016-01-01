package main

import (
	"net"
	"os"
)

// DialUDP connects to the remote address raddr on the network net,
// which must be "udp", "udp4", or "udp6".  If laddr is not nil, it is used
// as the local address for the connection.
func DialUDP(netType string, laddr, raddr *net.UDPAddr) (c *net.UDPConn, err os.Error) {
    switch netType {
    case "udp", "udp4", "udp6":
    default:
        return nil, net.UnknownNetworkError(netType)
    }
    if raddr == nil {
        return nil, &net.OpError{"dial", "udp", nil, errMissingAddress}
    }
    fd, e := internetSocket(net, laddr.toAddr(), raddr.toAddr(), syscall.SOCK_DGRAM, 0, "dial", sockaddrToUDP)
    if e != nil {
        return nil, e
    }
    return newUDPConn(fd), nil
}

func main() {
	// body
}

