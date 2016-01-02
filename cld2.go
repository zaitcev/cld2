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
	fmt.Printf("OK=%d\n", CLE_OK)
	var t cld_msg_get_resp
	Print(t)
	Print(listenport) // P3
}
