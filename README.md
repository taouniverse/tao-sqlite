# github.com/taouniverse/tao-sqlite

[![Go Report Card](https://goreportcard.com/badge/github.com/taouniverse/tao-sqlite)](https://goreportcard.com/report/github.com/taouniverse/tao-sqlite)
[![GoDoc](https://pkg.go.dev/badge/github.com/taouniverse/tao-sqlite?status.svg)](https://pkg.go.dev/github.com/taouniverse/tao-sqlite?tab=doc)

Tao Universe 组件单元（Unit），基于泛型工厂模式封装 **SQLite** 嵌入式数据库。

## 安装

```bash
go get github.com/taouniverse/tao-sqlite
```

## 使用

### 导入

```go
import _ "github.com/taouniverse/tao-sqlite"
```

### 配置

```yaml
# 单实例配置
sqlite:
  db: app.db

# 多实例配置
sqlite:
  default_instance: app
  app:
    db: data/app.db
  log:
    db: data/log.db
```

### 配置字段说明

| 字段 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `db` | string | `app.db` | SQLite 数据库文件路径 |
| `mode` | string | `rwc` | 打开模式 (ro/rw/rwc/memory) |
| `cache` | string | `shared` | 缓存模式 (shared/private) |
| `busy_timeout` | duration | `5s` | 忙等待超时 |
| `journal_mode` | string | `WAL` | 日志模式 (DELETE/TRUNCATE/PERSIST/MEMORY/WAL/OFF) |
| `max_idle` | int | `10` | 空闲连接池大小 |
| `max_open` | int | `100` | 最大打开连接数 |
| `max_lifetime` | duration | `1h` | 连接最大生命周期 |

## 工厂模式 API

| API | 说明 |
|-----|------|
| `sqlite.M` | 配置实例 `*Config` |
| `sqlite.Factory` | `*tao.BaseFactory[*gorm.DB]` 工厂实例 |
| `sqlite.DB()` | 获取默认数据库连接 `(*gorm.DB, error)` |
| `sqlite.GetDB(name)` | 获取指定名称的连接 `(*gorm.DB, error)` |

## 使用示例

### 获取连接并执行操作

```go
package main

import (
    "log"
    
    "github.com/taouniverse/tao-sqlite"
)

func main() {
    // 获取默认实例
    db, err := sqlite.DB()
    if err != nil {
        log.Fatal(err)
    }
    
    // 获取底层 sql.DB 进行 Ping 测试
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal(err)
    }
    
    err = sqlDB.Ping()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("SQLite 连接成功")
}
```

### GORM 操作

```go
db, _ := sqlite.DB()

// 自动迁移
db.AutoMigrate(&User{})

// 创建记录
db.Create(&User{Name: "tao", Age: 18})

// 查询记录
var user User
db.First(&user, "name = ?", "tao")

// 更新记录
db.Model(&user).Update("age", 20)

// 删除记录
db.Delete(&user)
```

### 多实例使用

```go
// 获取应用数据库
appDB, _ := sqlite.GetDB("app")

// 获取日志数据库
logDB, _ := sqlite.GetDB("log")

// 应用数据操作
appDB.Create(&User{Name: "tao"})

// 日志记录操作
logDB.Create(&Log{Message: "user created"})
```

## 单元测试

### 快速测试

```bash
# 仅运行配置相关测试
go test -v -run "TestConfig" ./...
```

### 完整集成测试

```bash
# 运行单实例测试
make test

# 运行多实例测试
make test-multi

# 运行所有测试
make test-all

# 生成覆盖率报告
make coverage

# 清理测试生成的数据库文件
make clean
```

### 手动测试

```bash
# 运行单实例测试
go test -v ./...

# 运行多实例测试
TAO_TEST_MULTI_INSTANCE=true go test -v ./...
```

## 特性

- **纯 Go 实现**：使用 modernc.org/sqlite，无 CGO 依赖，跨平台编译友好
- **嵌入式**：适合轻量级应用、开发环境和测试场景
- **多实例**：支持同时操作多个 SQLite 数据库文件

## 开发指南

| 文件 | 说明 |
|------|------|
| `config.go` | InstanceConfig 字段 + ValidSelf 默认值 |
| `sqlite.go` | NewSQLite 构造器 + 工厂注册 |
