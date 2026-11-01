#!/bin/bash
set -euo pipefail

BASE_DIR="./atcoder/ABC"
competition_number="${1:-}"

if [ -z "$competition_number" ]; then
  echo 'argument is not enough' >&2
  echo "USAGE: $0 [competition_number]" >&2
  exit 1
fi

difficulties=(a b)

for difficulty in "${difficulties[@]}"; do
  dir="$BASE_DIR/$competition_number/$difficulty"
  mkdir -p "$dir"
  if [ -e "$dir/main.go" ]; then
    echo "skip: $dir/main.go already exists" >&2
    continue
  fi
  cat << 'EOF' > "$dir/main.go"
package main

func main() {

}
EOF
done

echo "$BASE_DIR/$competition_number"
