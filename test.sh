#!/bin/sh
# Run this to obtain functional tests.
# Unlike some environments, we do not run "make" here to build what we're
# trying to test (not even "go build"). So, you run "make.sh" by hand.

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
  # This fallback is really not a good idea on Fedora, where a dead-end
  # version 0.8 is installed. The 0.8 attempts to communicate over UDP,
  # and we don't support that. We need something like 0.7.2 that works
  # over TCP. If we use the system cldcli, we should verify the version. XXX
  clibin=/usr/bin/cldcli
  if [ \! -x "$clibin" ]; then
    echo "test: cldcli not found in ., /usr/bin" >&2
    exit 1
  fi
fi

portfile=./test-cld.port
cld_host=localhost
#cld_port=  see below

#
# Start the server under test
#
rm -f "$portfile"
$srvbin -h "$cld_host" -p auto -f "$portfile" &
srv_pid=$!

i=0
while [ \! -f "$portfile" ]; do
  if [ $i -ge 10 ]; then
    echo "test: no port file $portfile" >&2
    kill $srv_pid
    exit 1
  fi
  sleep 1
  i=$(expr $i + 1)
done
cld_port=$(echo -n $(cat "$portfile"))

cleanup () {
  kill $srv_pid
  rm -f "$portfile"
}
trap cleanup SIGINT

# no commands opens a session anyway
echo quit | $clibin -h "$cld_host:$cld_port"
if [ "$?" != 0 ]; then
  echo "FAIL quit (dummy)" >&2
  cleanup
  exit 1
fi

echo "OK"
cleanup
exit 0
