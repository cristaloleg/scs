package scs

import (
	"bytes"
	"fmt"
	"syscall"
)

// Connection describes a new socket connection.
type Connection struct {
	// fs is a file descriptor
	fd int
	sa syscall.Sockaddr
}

// NewConnection creates a new connection with a given file descriptor and socker address.
func NewConnection(fd int, sa syscall.Sockaddr) Connection {
	return Connection{
		fd: fd,
		sa: sa,
	}
}

func (conn Connection) String() string {
	return fmt.Sprintf("%v %v", conn.sa, conn.fd)
}

// Close will close a connection.
func (conn Connection) Close() error {
	return syscall.Close(conn.fd)
}

func (conn Connection) Read() (bytes.Buffer, error) {
	var result bytes.Buffer

	for {
		bufSize := 64
		buf := make([]byte, bufSize)

		bytesRead, readErr := syscall.Read(conn.fd, buf)
		if readErr != nil {
			return result, readErr
		}
		if bytesRead <= 0 {
			return result, nil
		}

		result.Write(buf)

		if bytesRead < bufSize {
			bytesRead = 0
			return result, nil
		}
	}
}
