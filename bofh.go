package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	var (
		lst net.Listener
		con net.Conn
		err error

		addr = flag.String("addr", "localhost", "listen address")
		port = flag.String("port", "666", "listen port")
	)

	flag.Usage = usage
	flag.Parse()

	// Initialize the server.
	lst, err = net.Listen("tcp", fmt.Sprintf("%s:%s", *addr, *port))
	if err != nil {
		log.Fatalf("listen: %s", err)
	}
	defer lst.Close()

	log.Printf(
		"[%s:%s] Started `The Bastard Operator From Hell' excuse server.",
		*addr, *port,
	)

	// Accept requests.
	for {
		con, err = lst.Accept()
		if err != nil {
			log.Fatalf("accept: %s", err)
		}
		go getExcuse(con)
	}
}

// getExcuse gets a random excuse from a list of excuses and writes it out
// to the connection.
func getExcuse(c net.Conn) {
	c.Write([]byte(Excuses[randIdx(len(Excuses)-1)] + "\n"))
	c.Close()
	io.Copy(ioutil.Discard, c)
}

// randIdx returns a random number within the specified range.
func randIdx(n int) uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(n))
}
