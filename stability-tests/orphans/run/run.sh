#!/bin/bash
rm -rf /tmp/nonsensed-temp

nonsensed --simnet --appdir=/tmp/nonsensed-temp --profile=6061 &
NONSENSED_PID=$!

sleep 1

orphans --simnet -alocalhost:42511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $NONSENSED_PID

wait $NONSENSED_PID
NONSENSED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nonsensed exit code: $NONSENSED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NONSENSED_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
