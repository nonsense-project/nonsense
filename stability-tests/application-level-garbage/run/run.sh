#!/bin/bash
rm -rf /tmp/nonsensed-temp

nonsensed --devnet --appdir=/tmp/nonsensed-temp --profile=6061 --loglevel=debug &
NONSENSED_PID=$!
NONSENSED_KILLED=0
function killNonsensedIfNotKilled() {
    if [ $NONSENSED_KILLED -eq 0 ]; then
      kill $NONSENSED_PID
    fi
}
trap "killNonsensedIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:42611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $NONSENSED_PID

wait $NONSENSED_PID
NONSENSED_KILLED=1
NONSENSED_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nonsensed exit code: $NONSENSED_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NONSENSED_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
