#!/bin/bash

echo "mode: atomic" > coverage.out

PACKAGES=`go list ./...`

EXIT_CODE=0

for PKG in $PACKAGES; do
  echo $PKG
  echo $ROOT
  go test -race -v -coverprofile=coverprofile.out -covermode=atomic $PKG; __EXIT_CODE__=$?

  if [ "$__EXIT_CODE__" -ne "0" ]; then
    EXIT_CODE=$__EXIT_CODE__
  fi

  if [ -f coverprofile.out ]; then
    tail -n +2 coverprofile.out >> coverage.out; rm coverprofile.out
  fi
done

exit $EXIT_CODE
