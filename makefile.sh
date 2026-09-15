#!/bin/bash

cd atcoder/ABC
competition_number=$1

if [ -z "${competition_number}"]; then
  echo 'arguement is not ecnough'
  echo 'USAGE: ./makfile.sh [competition_number]'
  exit 0
fi


difficulties=("a" "b" "c")

for difficulty in ${difficulties[@]}
do
  mkdir -p ${competition_number}/${difficulty}
  touch ${competition_number}/${difficulty}/main.go
done

cd ${competition_number}
