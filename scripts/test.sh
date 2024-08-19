#!/bin/bash
make build
./basic 2>&1 | sed 's/time="[^"]*" //' >test/actual.txt
diff test/actual.txt test/expected.txt
