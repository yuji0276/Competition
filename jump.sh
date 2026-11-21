#!/bin/bash
BASE_DIR="./atcoder/ABC"

 if [ -z "${1:-}" ]; then
   echo "USAGE: source jump.sh <competition_number>" >&2
   return 1
 fi

 if [ ! -d "$BASE_DIR/$1" ]; then
   echo "not found: $BASE_DIR/$1 (先に makefile.sh $1 を実行してください)" >&2
   return 1
 fi

 cd "$BASE_DIR/$1"
