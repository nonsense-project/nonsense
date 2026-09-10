#!/bin/bash
rm -rf /tmp/nonsensed-temp

nonsensed --devnet --appdir=/tmp/nonsensed-temp --profile=6061 --loglevel=debug &
NONSENSED_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $NONSENSED_PID

wait $NONSENSED_PID
NONSENSED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nonsensed exit code: $NONSENSED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NONSENSED_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
