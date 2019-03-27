package scs

import (
	"errors"
	"syscall"
)

// NewFD return a new file descriptor for a given network.
func NewFD(network string) (int, error) {
	var domain, typ, proto int

	switch network {
	case "tcp", "tcp4":
		domain = syscall.AF_INET
		typ = syscall.SOCK_STREAM
		proto = syscall.IPPROTO_TCP
	case "tcp6":
		domain = syscall.AF_INET6
		typ = syscall.SOCK_STREAM
		proto = syscall.IPPROTO_TCP

	case "udp", "udp4":
		domain = syscall.AF_INET
		typ = syscall.SOCK_DGRAM
		proto = syscall.IPPROTO_UDP
	case "udp6":
		domain = syscall.AF_INET6
		typ = syscall.SOCK_DGRAM
		proto = syscall.IPPROTO_UDP

	default:
		return 0, errors.New("unknown network")
	}

	fd, err := syscall.Socket(domain, typ, proto)
	if err != nil {
		return 0, err
	}
	return fd, nil
}
