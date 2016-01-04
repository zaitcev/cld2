// CLD2

package main

import (
	"flag"
	"fmt"
	// "net"
	"os"
)

const TAG = "cld2"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: cld2 [-h host] [-p port]\n")
}

func main() {
	// if len(os.Args) > 1 {
	// 	fmt.Fprintf(os.Stderr, "Usage: cld2\n");
	// 	os.Exit(1)
	// }

	var listenhost string
	var listenport int

	flag.StringVar(&listenhost, "h", "", "Hostname to bind for listening")
	flag.IntVar(&listenport, "p", 8081, "Port to bind for listening")

	flag.Usage = usage
	flag.Parse()

	fmt.Printf("host %s port %d\n", listenhost, listenport) // P3

	// P3
	var x XDR
	// var t cld_msg_get_resp -- too complex, leave out for now
	var t cld_msg_open
	t.mode = COM_READ|COM_LOCK
	t.events = CE_MASTER_FAILOVER
	t.inode_name = "/something"
	t.XDRencode(&x)
	buf := x.Fetch()
	for _, n := range buf {
		fmt.Printf(" %02x", n)
	}
	fmt.Printf("\n")
}
