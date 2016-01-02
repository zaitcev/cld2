//
// In a complete implementation of XDR, this file would be generated
// from cld_msg_rpc.x by rpcgen. For now we re-coded cld_msg_rpc.h
// by hand.
//
package main

const CLD_PKT_MAGIC = "CLDc1pkt"
const CLD_SID_SZ = 8
const CLD_INODE_NAME_MAX = 256
const CLD_MAX_USERNAME = 32
const CLD_MAX_PKT_MSG_SZ =   1024
const CLD_MAX_PAYLOAD_SZ = 131072
const CLD_MAX_MSG_SZ     = 196608
const CLD_MAX_SECRET_KEY = 128

/** available RPC operations */
type cld_msg_op int
const (
	/* client -> server */
	CMO_NOP = 0
	CMO_NEW_SESS = 1
	CMO_OPEN = 2
	CMO_GET_META = 3
	CMO_GET = 4
	CMO_PUT = 6
	CMO_CLOSE = 7
	CMO_DEL = 8
	CMO_LOCK = 9
	CMO_UNLOCK = 10
	CMO_TRYLOCK = 11
	CMO_ACK = 12
	CMO_END_SESS = 13

	/* server -> client */
	CMO_PING = 14
	CMO_NOT_MASTER = 15
	CMO_EVENT = 16
	CMO_ACK_FRAG = 17

	CMO_AFTER_LAST = 18	// variable value in the source
)

/** CLD error codes */
type cle_err_codes int
const (
	CLE_OK = 0			/**< success / no error */
	CLE_SESS_EXISTS = 1		/**< session exists */
	CLE_SESS_INVAL = 2		/**< session doesn't exist */
	CLE_DB_ERR = 3			/**< db error */
	CLE_BAD_PKT = 4			/**< invalid/corrupted packet */
	CLE_INODE_INVAL = 5		/**< inode doesn't exist */
	CLE_NAME_INVAL = 6		/**< inode name invalid */
	CLE_OOM = 7			/**< server out of memory */
	CLE_FH_INVAL = 8		/**< file handle invalid */
	CLE_DATA_INVAL = 9		/**< invalid data pkt */
	CLE_LOCK_INVAL = 10		/**< invalid lock */
	CLE_LOCK_CONFLICT = 11		/**< conflicting lock held */
	CLE_LOCK_PENDING = 12		/**< lock waiting to be acq. */
	CLE_MODE_INVAL = 13		/**< op incompat. w/ file mode */
	CLE_INODE_EXISTS = 14		/**< inode exists */
	CLE_DIR_NOTEMPTY = 15		/**< dir not empty */
	CLE_INTERNAL_ERR = 16		/**< nonspecific internal err */
	CLE_TIMEOUT = 17		/**< session timed out */
	CLE_SIG_INVAL = 18		/**< HMAC sig bad / auth failed */
)

/** availble OPEN mode flags */
type cld_open_modes int
const (
	COM_READ = 0x01			/**< read */
	COM_WRITE = 0x02		/**< write */
	COM_LOCK = 0x04			/**< lock */
	COM_ACL = 0x08			/**< ACL update */
	COM_CREATE = 0x10		/**< create file, if not exist */
	COM_EXCL = 0x20			/**< fail create if file exists */
	COM_DIRECTORY = 0x40		/**< operate on a directory */
)

/** potential events client may receive */
type cld_events int
const (
	CE_UPDATED = 0x01		/**< contents updated */
	CE_DELETED = 0x02		/**< inode deleted */
	CE_LOCKED = 0x04		/**< lock acquired */
	CE_MASTER_FAILOVER = 0x08	/**< master failover */
	CE_SESS_FAILED = 0x10
)

type cld_inode_flags int
const (
	CIFL_DIR = 0x01			/**< inode is a directory */
)

/** LOCK flags */
type cld_lock_flags int
const (
	CLF_SHARED = 0x01		/**< a shared (read) lock */
)

/** Describes whether a packet begins, continues, or ends a message. */
type cld_pkt_order_t int
const (
	CLD_PKT_ORD_MID = 0x0
	CLD_PKT_ORD_FIRST = 0x1
	CLD_PKT_ORD_LAST = 0x2
	CLD_PKT_ORD_FIRST_LAST = 0x3
)
const CLD_PKT_IS_FIRST = 0x1
const CLD_PKT_IS_LAST = 0x2

/** Information that appears only in the first packet */
type cld_pkt_msg_infos struct {
	xid int64				/**< opaque message id */
	op cld_msg_op				/**< message operation */
}

/** Information about the message contained in this packet */
// union cld_pkt_msg_info switch (enum cld_pkt_order_t order) {
// 	case CLD_PKT_ORD_MID:
// 	case CLD_PKT_ORD_LAST:
// 		void;
// 	case CLD_PKT_ORD_FIRST:
// 	case CLD_PKT_ORD_FIRST_LAST:
// 		struct cld_pkt_msg_infos mi;
// };
// Must use a XDRcoder interface XXX
type cld_pkt_msg_info struct {
	order cld_pkt_order_t
	// union {
	// 	struct cld_pkt_msg_infos mi;
	// } cld_pkt_msg_info_u;
}

type cld_pkt_hdr struct {
	magic int64
	sid int64	// Should be unsigned, but the original .x uses hyper
	// string user<CLD_MAX_USERNAME>;	/**< authenticated user */
	user string
	mi cld_pkt_msg_info
}

/** generic response for PUT, CLOSE, DEL, LOCK, UNLOCK */
type cld_msg_generic_resp struct {
	code cle_err_codes
	xid_in int64				/**< C->S xid */
}

/** ACK-FRAG message */
type cld_msg_ack_frag struct {
	seqid int64
}

/** OPEN message */
type cld_msg_open struct {
	mode int				/**< open mode, COM_xxx */
	events cld_events
	inode_name string
}

/** OPEN message response */
type cld_msg_open_resp struct {
	msg cld_msg_generic_resp
	fh int64
}

/** GET message */
type cld_msg_get struct {
	fh int64
}

/** GET message response */
type cld_msg_get_resp struct {
	msg cld_msg_generic_resp
	inum uint64				/**< unique inode number */
	vers uint64				/**< inode version */
	time_create int64			/**< creation time */
	time_modify int64			/**< last modification time */
	flags cld_inode_flags
	// string			inode_name<CLD_INODE_NAME_MAX>;
	inode_name string
	// opaque			data<CLD_MAX_PAYLOAD_SZ>;
	data []byte
}

/** PUT message */
type cld_msg_put struct {
	fh int64
	// opaque			data<CLD_MAX_PAYLOAD_SZ>;
	data []byte
}

/** CLOSE message */
type cld_msg_close struct {
	fh int64
}

/** DEL message */
type cld_msg_del struct {
	//   string inode_name<CLD_INODE_NAME_MAX>;
	inode_name string
}

/** UNLOCK message */
type cld_msg_unlock struct {
	fh int64
}

/** LOCK message */
type cld_msg_lock struct {
	fh int64
	flags cld_lock_flags
}

/** Server-to-client EVENT message */
type cld_msg_event struct {
	fh int64
	events cld_events
}
