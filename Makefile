# SQLite Unit 测试工具

.PHONY: help test test-multi test-all coverage clean

# 默认目标
help:
	@echo "SQLite Unit 测试工具"
	@echo ""
	@echo "可用命令:"
	@echo "  make test        - 运行单实例测试 (默认)"
	@echo "  make test-multi  - 运行多实例测试"
	@echo "  make test-all    - 运行所有测试 (单实例 + 多实例)"
	@echo "  make coverage    - 生成测试覆盖率报告"
	@echo "  make clean       - 清理测试生成的文件"

# 单实例测试
test:
	go test -v ./...

# 多实例测试
test-multi:
	TAO_TEST_MULTI_INSTANCE=true go test -v ./...

# 运行所有测试
test-all:
	@echo "=== 运行单实例测试 ==="
	go test -v ./...
	@echo ""
	@echo "=== 运行多实例测试 ==="
	TAO_TEST_MULTI_INSTANCE=true go test -v ./...

# 生成覆盖率报告
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 清理
clean:
	rm -f coverage.out coverage.html
	rm -f *.db
	rm -rf data/
