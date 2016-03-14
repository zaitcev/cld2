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

func TestCLDencode_cld_pkt_hdr_1(t *testing.T) {
	var sample = [44]byte {
		0x74, 0x6b, 0x70, 0x31, 0x63, 0x44, 0x4c, 0x43, 0x77, 0x55,
		0x33, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
		0x7a, 0x61, 0x69, 0x74, 0x63, 0x65, 0x76, 0x00, 0x00, 0x00,
		0x00, 0x01, 0x32, 0x7b, 0x23, 0xc6, 0x6b, 0x8b, 0x45, 0x67,
		0x00, 0x00, 0x00, 0x01,
	}
	var x XDR
	var sut cld_pkt_hdr
	strtole8(&sut.magic, CLD_PKT_MAGIC)
	tmp := [8]byte { 0, 0, 0, 0, 0x11, 0x33, 0x55, 0x77 }
	strtole8(&sut.sid, string(tmp[0:8]))
	sut.user = "zaitcev"
	sut_info_2 := new(cld_pkt_msg_info_2)
	sut_infos := new(cld_pkt_msg_infos)
	sut_infos.xid = 0x327b23c66b8b4567
	sut_infos.op = CMO_NEW_SESS
	// XXX Learn Go: Is this a copy? Or why pointer cannot be
	// assigned to... reference? Is that mi a reference or inline?
	sut_info_2.mi = *sut_infos
	sut_info_2.order = CLD_PKT_ORD_FIRST
	sut.mi = sut_info_2
	sut.XDRencode(&x)
	result := x.Fetch()
	if !bytes.Equal(result, sample[:]) {
		// XXX time to factor this out across test procedures
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

func TestCLDencode_cld_pkt_hdr_2(t *testing.T) {
	var sample = [32]byte {
		0x74, 0x6b, 0x70, 0x31, 0x63, 0x44, 0x4c, 0x43, 0x77, 0x55,
		0x33, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
		0x7a, 0x61, 0x69, 0x74, 0x63, 0x65, 0x76, 0x00, 0x00, 0x00,
		0x00, 0x02,
	}
	var x XDR
	var sut cld_pkt_hdr
	strtole8(&sut.magic, CLD_PKT_MAGIC)
	tmp := [8]byte { 0, 0, 0, 0, 0x11, 0x33, 0x55, 0x77 }
	strtole8(&sut.sid, string(tmp[0:8]))
	sut.user = "zaitcev"
	sut_info_1 := new(cld_pkt_msg_info_1)
	sut_info_1.order = CLD_PKT_ORD_LAST
	sut.mi = sut_info_1
	sut.XDRencode(&x)
	result := x.Fetch()
	if !bytes.Equal(result, sample[:]) {
		// XXX time to factor this out across test procedures
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
