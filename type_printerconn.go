package debug // import go.ideatocode.tech/debug

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// PrinterConn prints all bytes that go through it
type PrinterConn struct {
	net.Conn
	Prefix string
	Writer io.Writer
}

func (pc PrinterConn) Read(b []byte) (int, error) {
	w := pc.Writer
	if w == nil {
		w = os.Stderr
	}
	n, err := pc.Conn.Read(b)
	old := log.Writer()
	log.Print(
		"\n============================================================================\n",
		"Read: ", pc.Prefix, ", Addr: ", pc.RemoteAddr().String(), ", Len: ", n, ", Err: ", err, "\n",
		spew.Sdump(
			b[0:n],
		),
		"============================================================================\n",
	)
	log.SetOutput(old)
	return n, err
}

func (pc PrinterConn) Write(b []byte) (int, error) {
	w := pc.Writer
	if w == nil {
		w = os.Stderr
	}
	n, err := pc.Conn.Write(b)
	old := log.Writer()
	log.Print(
		"\n============================================================================\n",
		"Write: ", pc.Prefix, ", Addr: ", pc.RemoteAddr().String(), ", Len: ", n, ", Err: ", err, "\n",
		spew.Sdump(
			b[0:],
		),
		"============================================================================\n",
	)
	log.SetOutput(old)
	return n, err
}
