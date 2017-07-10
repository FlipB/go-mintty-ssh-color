#!/bin/bash
# Script gets the SSH-process' PID, waits until it exits,
#  and resets the color of the terminal.
# Before this of course it sets the desired color (from Arg1).

test -z $1 && echo "Need to supply arg1 to specify desired color" && exit 1

SSH_BG_COLOR="$1"
ORG_BG_COLOR="002B36"

#SSH_BG_COLOR=202020
#ORG_BG_COLOR=000000



## Using cygwin we have to figure out the PID as seem by Windows.
# 1. Get current process (PID)
# 2. Get parent process (PPID) (this will be the SSH process if launched by it)
# 3. Get windows PID (WINPID) of PPID

PPID_G=$(ps | grep "^ " | awk '{print $1" "$4" "}' | egrep "^$PPID ")
test "$(echo $PPID_G | wc -l)" != "1" && exit 1

WINPID=$(echo $PPID_G | awk '{print $2}')
test -z "$WINPID" && exit 2

mintty-cygwin-ssh-bgcolor-daemon --watch-pid $WINPID --before $SSH_BG_COLOR --after $ORG_BG_COLOR > /dev/null &
