#!/bin/sh

# For now you could just do "go run *.go" but we may have *_test.go
# in this directory in the future.

# PACKAGE=main
go build -x cld2.go cld_msg_rpc.go xdr.go
