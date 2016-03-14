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
