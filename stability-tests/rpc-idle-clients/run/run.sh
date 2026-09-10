#!/bin/bash
rm -rf /tmp/nonsensed-temp

NUM_CLIENTS=128
nonsensed --devnet --appdir=/tmp/nonsensed-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
NONSENSED_PID=$!
NONSENSED_KILLED=0
function killNonsensedIfNotKilled() {
  if [ $NONSENSED_KILLED -eq 0 ]; then
    kill $NONSENSED_PID
  fi
}
trap "killNonsensedIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $NONSENSED_PID

wait $NONSENSED_PID
NONSENSED_EXIT_CODE=$?
NONSENSED_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Nonsensed exit code: $NONSENSED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NONSENSED_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
