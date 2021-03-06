// CLD2

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
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
	var _listenport string
	var portfile string

	flag.StringVar(&listenhost, "h", "", "Hostname to bind for listening")
	flag.StringVar(&_listenport, "p", "8081", "Port to bind for listening")
	flag.StringVar(&portfile,   "f", "",
	    "File to write out the listen port number")

	flag.Usage = usage
	flag.Parse()

	if _listenport == "auto" {
		listenport = 0
	} else {
		var err error
		listenport, err = strconv.Atoi(_listenport)
		if err != nil {
			fmt.Fprintf(os.Stderr,
			    "%s: Argument for -p is invalid\n", TAG)
			os.Exit(2)
		}
		if listenport < 1 || listenport > 65535 {
			fmt.Fprintf(os.Stderr,
			    "%s: Port number %d is out of range\n",
			    TAG, listenport)
			os.Exit(2)
		}
	}

	fmt.Printf("host %s port %d\n", listenhost, listenport) // P3
	_main(listenhost, listenport, portfile)
}

func _main(listenhost string, listenport int, portfile string) {
	listen_netloc := net.JoinHostPort(listenhost, strconv.Itoa(listenport))
	ln, err := net.Listen("tcp", listen_netloc)
	if err != nil {
		fmt.Fprintf(os.Stderr,
		    "%s: Listen(%s) error: %s\n",
		    TAG, listen_netloc, err.Error())
		os.Exit(1)
	}
	if len(portfile) != 0 {
		_, portstr, err := net.SplitHostPort(ln.Addr().String())
		if err == nil {
			write_port_to_file(portfile, portstr)
		}
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr,
			    "%s: Accept error: %s\n", TAG, err.Error())
			os.Exit(1)
		}

		fmt.Printf("connection\n") // P3
		conn.Close()
	}
}

/*
 * Errors are ignored for now. When test reads the port file, it will
 * report an error if it's not written.
 */
func write_port_to_file(portfile string, portstr string) {
	fp, err := os.OpenFile(portfile,
	    os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err == nil {
		// File implements interface io.Writer, so...
		fmt.Fprintf(fp, "%s\n", portstr)
		fp.Close()
	}
}

/*
 * This existential horror comes from practices like these:
 *
 * #define CLD_PKT_MAGIC "CLDc1pkt"
 * struct cld_pkt_hdr {
 *     quad_t magic;
 * } pkt;
 * memcpy(&pkt.magic, CLD_PKT_MAGIC, sizeof(pkt.magic));
 *
 * In C, this produces different results on little endian and big endian hosts.
 * In other words, the author employed a platform-independent format, XDR,
 * then defeated it by coercing strings into ints.
 */
func strtole8(dst *int64, src string) {
	var a int64 = 0
	for i := 0; i < 8; i++ {
		a |= (int64(src[i]) << uint(i*8))
	}
	*dst = a
}

func le8tostr(v int64) string {
	var a8 [8]byte
	for i := 0; i < 8; i++ {
		a8[i] = byte(v >> uint(i*8))
	}
	return string(a8[0:8])
}
