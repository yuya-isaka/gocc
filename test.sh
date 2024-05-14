#!/bin/bash

# 定義: assert関数
# 用途: Goプログラムをコンパイルし、指定された入力で実行後、結果を期待値と比較する。
# 引数:
#   $1 (expect) - 期待される結果の終了ステータス
#   $2 (input)  - Goプログラムに渡す入力
assert() {
  local expect="$1"
  local input="$2"

  # Goプログラムを実行し、出力をアセンブリファイルにリダイレクト
  go run main.go "$input" > tmp.s || exit 1

  # アセンブリファイルをGCCでコンパイル
  gcc -o tmp tmp.s || exit 1

  # コンパイルされたプログラムを実行
  ./tmp
  local actual="$?"

  # 結果の検証と出力
  if [ "$actual" -eq "$expect" ]; then
    echo "$input => $actual"
  else
    echo "$input => got $actual expected $expect"
    exit 1
  fi
}

# テストケース
assert 0 0
assert 42 42
assert 21 '5+20-4'

# すべてのテストが成功した場合
echo "OK"
