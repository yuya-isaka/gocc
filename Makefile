# Makefile for managing test execution and cleanup

# ターゲット: test
# 用途: テストスクリプトを実行する
test:
	@./test.sh

# ターゲット: clean
# 用途: 不要な一時ファイルを削除する
clean:
	@rm -f tmp*

# .PHONYを使って、実際のファイルではなくターゲット名であることを明示
.PHONY: test clean
