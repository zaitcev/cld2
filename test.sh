#!/bin/sh
# Run this to obtain functional tests.

#
# Set up the environment
#

srvbin=./cld2
if [ \! -x "$srvbin" ]; then
  echo "test: cld2 not found in ." >&2
  exit 1
fi

clibin=./cldcli
if [ \! -x "$clibin" ]; then
  clibin=/usr/bin/cldcli
  if [ \! -x "$clibin" ]; then
    echo "test: cldcli not found in ., /usr/bin" >&2
    exit 1
  fi
fi

portfile=./cld.port

#
# Start the server under test
#
rm -f "$portfile"
$srvbin -h localhost -p auto -f "$portfile"

#trap
