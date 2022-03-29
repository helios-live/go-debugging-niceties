package debugging

import (
	"io"
	"log"

	"github.com/davecgh/go-spew/spew"
)

// PrintReader Prints whatever it reads
type PrintReader struct {
	io.Reader
	Prefix string
}

func (p PrintReader) Read(b []byte) (n int, err error) {
	n, err = p.Reader.Read(b)
	log.Println(p.Prefix, "PrintReader", n, err)
	if n > 0 {
		spew.Dump(b[0:n])
	}
	return n, err
}
