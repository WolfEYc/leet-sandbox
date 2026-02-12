#!/bin/bash

if [[ -z "$beans" ]]; then
    echo "empty or unset"
fi
beans="trouble"
if [[ -z "$beans" ]]; then
    echo "empty or unset"
fi

if [[ -v beans ]]; then
  echo "beans: ${beans}"
fi



