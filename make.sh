#!/bin/sh
# Run this to build the project (cld2). You don't really need a build
# script with Go, but we like to have it for now. Also provides an
# inventory of source files.

# PACKAGE=main
go build -x cld2.go cld_msg_rpc.go xdr.go
