# 备忘录模式 (Memento Pattern)

## 概述

备忘录模式是一种行为型设计模式，它允许在不暴露对象实现细节的情况下捕获和恢复对象的内部状态。该模式提供了一种状态恢复机制，使得对象可以回滚到之前的状态，而不会破坏封装性。

## 核心概念

备忘录模式通过创建备忘录对象来存储原发器对象的内部状态，当需要时可以将原发器恢复到之前保存的状态。备忘录对象对客户端是透明的，只有原发器才能访问备忘录的内容，从而保证了封装性。

## 模式结构

### 1. 原发器角色 (Originator)
- `Game` 结构体：游戏状态的原发器，包含游戏的核心状态（HP和MP）
- 负责创建备忘录和从备忘录恢复状态
- 包含业务方法和状态管理方法

### 2. 备忘录角色 (Memento)
- `gameMemento` 结构体：存储原发器的内部状态
- 备忘录是不可变的，一旦创建就不能修改
- 对客户端透明，只有原发器可以访问其内容

### 3. 负责人角色 (Caretaker)
- 负责保存和管理备忘录对象
- 不修改或检查备忘录的内容
- 在本例中由客户端代码充当此角色

## 代码结构分析

### 原发器设计
```go
type Game struct {
    hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
    g.mp += mpDelta
    g.hp += hpDelta
}
```

### 备忘录创建
```go
func (g *Game) Save() Memento {
    return &gameMemento{
        hp: g.hp,
        mp: g.mp,
    }
}
```

### 状态恢复
```go
func (g *Game) Load(m Memento) {
    gm := m.(*gameMemento)
    g.mp = gm.mp
    g.hp = gm.hp
}
```

### 备忘录结构
```go
type gameMemento struct {
    hp, mp int
}
```

## 使用场景

### 1. 撤销/重做功能
在文本编辑器、图形编辑器、游戏等应用中实现撤销和重做功能。

### 2. 事务回滚
在数据库操作或业务事务中，当操作失败时回滚到之前的状态。

### 3. 游戏存档
游戏中保存和加载游戏进度，允许玩家回到之前的游戏状态。

### 4. 状态快照
定期保存系统状态，用于故障恢复或历史记录查询。

### 5. 配置管理
保存和恢复应用程序的配置状态。

## 优点

1. **保持封装**：不会破坏原发器对象的封装性
2. **简化原发器**：原发器不需要自己管理状态历史
3. **状态恢复**：提供了一种简单的状态恢复机制
4. **内存友好**：可以限制保存的状态数量
5. **易于扩展**：可以轻松添加新的状态类型

## 缺点

1. **内存消耗**：保存大量备忘录可能会消耗大量内存
2. **性能开销**：创建备忘录和恢复状态都有性能开销
3. **实现复杂**：需要额外的代码来管理备忘录
4. **状态一致性**：需要确保备忘录中的状态与当前状态兼容

## 与其他模式的关系

### 备忘录模式 vs 命令模式
- **备忘录模式**：保存和恢复对象状态
- **命令模式**：封装操作本身，支持撤销操作

### 备忘录模式 vs 原型模式
- **备忘录模式**：保存状态的快照用于恢复
- **原型模式**：创建对象的完整副本

### 备忘录模式 vs 状态模式
- **备忘录模式**：存储和恢复状态
- **状态模式**：定义对象在不同状态下的行为

## 实际应用示例

### 1. 游戏进度保存
```go
// 游戏存档系统
type GameSaveSystem struct {
    saves map[string]*Game
}

func (gss *GameSaveSystem) SaveGame(name string, game *Game) {
    gss.saves[name] = game
}

func (gss *GameSaveSystem) LoadGame(name string) *Game {
    return gss.saves[name]
}
```

### 2. 文档编辑器撤销
```go
// 文档状态备忘录
type DocumentMemento struct {
    content string
    cursor  int
}

type Document struct {
    content string
    cursor  int
}

func (d *Document) Save() *DocumentMemento {
    return &DocumentMemento{
        content: d.content,
        cursor:  d.cursor,
    }
}

func (d *Document) Load(m *DocumentMemento) {
    d.content = m.content
    d.cursor = m.cursor
}
```

### 3. 数据库事务回滚
```go
// 数据库状态快照
type DatabaseMemento struct {
    tables map[string][]Row
}

type Database struct {
    tables map[string][]Row
}

func (db *Database) BeginTransaction() *DatabaseMemento {
    return &DatabaseMemento{tables: db.tables}
}

func (db *Database) Rollback(m *DatabaseMemento) {
    db.tables = m.tables
}
```

### 4. 配置管理器
```go
// 应用配置备忘录
type ConfigMemento struct {
    settings map[string]interface{}
}

type ConfigManager struct {
    settings map[string]interface{}
}

func (cm *ConfigManager) SaveConfig() *ConfigMemento {
    return &ConfigMemento{settings: cm.settings}
}

func (cm *ConfigManager) RestoreConfig(m *ConfigMemento) {
    cm.settings = m.settings
}
```

## 总结

备忘录模式提供了一种优雅的状态保存和恢复机制，在不破坏封装性的前提下，使得对象可以回滚到之前的状态。该模式在游戏开发、文档编辑、数据库事务等场景中都有广泛应用。通过合理的设计和实现，可以有效解决状态管理问题，提高系统的可靠性和用户体验。

你是资深go 语言开发专家，对18_flyweight下进行分析，请必须参考17_memento的README.md 格式，补充到的EADME.md内容