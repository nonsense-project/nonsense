#!/bin/bash

APPDIR=/tmp/nonsensed-temp
NONSENSED_RPC_PORT=29587

rm -rf "${APPDIR}"

nonsensed --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${NONSENSED_RPC_PORT}" --profile=6061 &
NONSENSED_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${NONSENSED_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $NONSENSED_PID

wait $NONSENSED_PID
NONSENSED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nonsensed exit code: $NONSENSED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NONSENSED_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
