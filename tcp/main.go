package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// open a net.listener
	l, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		log.Fatalln(err)
	}

	// call Accept in a loop, and start new goroutines for each Conn
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		// never block in Accept loop
		//go copyToStderr(conn)
		go logSNI(conn)
	}
}

func logSNI(conn net.Conn) {
	defer conn.Close()

	// set a timeout and read the message
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, conn, 1+2+2); err != nil {
		log.Println(err)
		return
	}

	length := binary.BigEndian.Uint16(buf.Bytes()[3:5])
	if _, err := io.CopyN(&buf, conn, int64(length)); err != nil {
		log.Println(err)
		return
	}

	// parse TLS
	ch, ok := ParseClientHello(buf.Bytes())
	if ok {
		log.Printf("Got a connection with SNI %q", ch.SNI)
	}
	c := prefixConn{
		Conn:   conn,
		Reader: io.MultiReader(&buf, conn),
	}

	conn.SetReadDeadline(time.Time{})

	cert, err := tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{
			cert,
		},
	}

	// wrap the net.Conn with tls.Server
	tlsConn := tls.Server(c, config)

	// service the wrapped plaintext connection
	proxy(tlsConn)

}

type prefixConn struct {
	net.Conn
	io.Reader
}

func (c prefixConn) Read(p []byte) (int, error) {
	return c.Reader.Read(p)
}

func proxy(conn net.Conn) {
	defer conn.Close()

	// connect to remote
	remote, err := net.Dial("tcp", "gophercon.com:443")
	if err != nil {
		log.Println(err)

		return
	}
	defer remote.Close()

	// copy to both direction
	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}

func copyToStderr(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("finished with err = %v", err)
			return
		}

		os.Stderr.Write(buf[:n])
	}
}
