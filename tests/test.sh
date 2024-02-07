#!/bin/bash

set -euo pipefail

NAME="testing-$(date +%F)"

e2b build --name "${NAME}"
e2b list | grep "${NAME}"

RESULT=$(node test.js)
if [ "$RESULT" != "Hello World" ]; then
    echo "Test failed: $RESULT"
    exit 1
fi

e2b delete -y

if [[ $(e2b list) =~ ${NAME} ]]; then
   echo "The template '${NAME}' wasn't deleted."
fi

