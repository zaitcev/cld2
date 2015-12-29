//
// In a complete implementation of XDR, this file would be generated
// from cld_msg_rpc.x by rpcgen. For now we re-coded cld_msg_rpc.h
// by hand.
//
package main

type Cle_err_codes int
const (
	CLE_OK = 0
	CLE_SESS_EXISTS = 1
	CLE_SESS_INVAL = 2
	CLE_DB_ERR = 3
	CLE_BAD_PKT = 4
	CLE_INODE_INVAL = 5
	CLE_NAME_INVAL = 6
	CLE_OOM = 7
	CLE_FH_INVAL = 8
	CLE_DATA_INVAL = 9
	CLE_LOCK_INVAL = 10
	CLE_LOCK_CONFLICT = 11
	CLE_LOCK_PENDING = 12
	CLE_MODE_INVAL = 13
	CLE_INODE_EXISTS = 14
	CLE_DIR_NOTEMPTY = 15
	CLE_INTERNAL_ERR = 16
	CLE_TIMEOUT = 17
	CLE_SIG_INVAL = 18
)

type Cld_msg_generic_resp struct {
	code Cle_err_codes
	xid_in uint64
}

type Cld_msg_get_resp struct {
	msg Cld_msg_generic_resp
	inum uint64
	vers uint64
	time_create uint64
	time_modify uint64
	flags int
	inode_name string
	data []byte
}
