package debug

import (
	"io"
	"net"
	"testing"
)

func TestPrinterConn_Read(t *testing.T) {
	type fields struct {
		Conn   net.Conn
		Prefix string
		Writer io.Writer
	}
	type args struct {
		b []byte
	}

	server, client := net.Pipe()
	go func() {
		x := []byte("hello!")
		for {
			_, err := server.Write(x)
			if err != nil {
				break
			}
		}
		// Do some stuff
		server.Close()
	}()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Simple",
			fields: fields{
				Conn:   client,
				Prefix: "test",
			},
			args: args{
				b: make([]byte, 100),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := PrinterConn{
				Conn:   tt.fields.Conn,
				Prefix: tt.fields.Prefix,
				Writer: tt.fields.Writer,
			}
			got, err := pc.Read(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrinterConn.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrinterConn.Read() = %v, want %v", got, tt.want)
			}
		})
	}
	// Do some stuff
	client.Close()
}
func TestPrinterConn_Write(t *testing.T) {
	type fields struct {
		Conn   net.Conn
		Prefix string
		Writer io.Writer
	}
	type args struct {
		b []byte
	}

	server, client := net.Pipe()
	go func() {
		x := []byte("hello!")
		for {
			_, err := server.Write(x)
			if err != nil {
				break
			}
		}
		// Do some stuff
		server.Close()
	}()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Simple",
			fields: fields{
				Conn:   client,
				Prefix: "test",
			},
			args: args{
				b: make([]byte, 100),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := PrinterConn{
				Conn:   tt.fields.Conn,
				Prefix: tt.fields.Prefix,
				Writer: tt.fields.Writer,
			}
			got, err := pc.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrinterConn.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrinterConn.Write() = %v, want %v", got, tt.want)
			}
		})
	}
	// Do some stuff
	client.Close()
}
