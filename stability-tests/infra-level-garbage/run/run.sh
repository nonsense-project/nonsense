#!/bin/bash
rm -rf /tmp/nonsensed-temp

nonsensed --devnet --appdir=/tmp/nonsensed-temp --profile=6061 &
NONSENSED_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:42611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $NONSENSED_PID

wait $NONSENSED_PID
NONSENSED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nonsensed exit code: $NONSENSED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NONSENSED_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
