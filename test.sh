# !/bin/bash

assert() {
  expect="$1"
  input="$2"

  go run main.go "$input" > tmp.s || exit
  gcc -o tmp tmp.s
  ./tmp
  actual="$?"

  if [ "$actual" = "$expect" ]; then
    echo "$input => $actual"
  else
    echo "$input => got $actual expected $expect"
    exit 1
  fi
}

assert 0 0
assert 42 42

echo OK