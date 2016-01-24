// Unit test for encoding and decoding in CLD2 specific XDR
package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestCLDencode_cld_msg_open(t *testing.T) {
	var sample = [24]byte {
		0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00,
		0x00, 0x0a, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69,
		0x6e, 0x67, 0x00, 0x00,
	}
	var x XDR
	var sut cld_msg_open
	sut.mode = COM_READ|COM_LOCK
	sut.events = CE_MASTER_FAILOVER
	sut.inode_name = "/something"
	sut.XDRencode(&x)
	result := x.Fetch()
	if !bytes.Equal(result, sample[:]) {
		fmt.Fprintf(os.Stderr, "sample:")
		for _, n := range sample {
			fmt.Fprintf(os.Stderr, " %02x", n)
		}
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "result:")
		for _, n := range result {
			fmt.Fprintf(os.Stderr, " %02x", n)
		}
		fmt.Fprintf(os.Stderr, "\n")
		t.Fail()
	}

}

func TestCLDencode_cld_msg_open_resp(t *testing.T) {
	var sample = [20]byte {
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb1, 0x6b,
		0x00, 0xb5, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0xb3,
	}
	var x XDR
	var sut cld_msg_open_resp
	sut.msg.xid_in = 0xB16B00B5
	sut.msg.code = CLE_OK
	sut.fh = 5555
	sut.XDRencode(&x)
	result := x.Fetch()
	if !bytes.Equal(result, sample[:]) {
		fmt.Fprintf(os.Stderr, "sample:")
		for _, n := range sample {
			fmt.Fprintf(os.Stderr, " %02x", n)
		}
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "result:")
		for _, n := range result {
			fmt.Fprintf(os.Stderr, " %02x", n)
		}
		fmt.Fprintf(os.Stderr, "\n")
		t.Fail()
	}
}
