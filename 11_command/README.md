# 命令模式 (Command Pattern)

## 概述

命令模式是一种行为型设计模式，它将请求封装成对象，从而使你可以用不同的请求对客户进行参数化。这种模式将请求的发送者和接收者解耦，使得发送者不需要知道接收者的任何细节。

命令模式的核心思想是将操作（命令）封装为对象，这样就可以将命令存储、传递、排队、记录日志或者撤销操作。

## 核心概念

命令模式将请求的发送者（Invoker）和接收者（Receiver）解耦。发送者只需要知道如何发送命令，而不需要知道命令是如何执行的，也不需要知道接收者是谁。命令对象封装了接收者和一个或多个动作。

## 模式结构

### 1. 命令接口 (Command)
- `Command` 接口：定义了执行操作的接口
- `Execute()` 方法：执行命令的方法

### 2. 具体命令角色 (Concrete Commands)
- `StartCommand`：启动系统的具体命令实现
- `RebootCommand`：重启系统的具体命令实现
- 持有接收者的引用，并实现具体的执行逻辑

### 3. 接收者角色 (Receiver)
- `MotherBoard`：接收者，真正执行命令的对象
- 实现了具体的业务逻辑：`Start()` 和 `Reboot()` 方法

### 4. 调用者角色 (Invoker)
- `Box`：调用者，持有命令对象并在某个时间点调用命令
- `PressButton1()` 和 `PressButton2()`：调用命令的方法

### 5. 客户端角色 (Client)
- 创建具体的命令对象并设置其接收者
- 将命令对象传递给调用者

## 代码结构分析

### 1. 命令接口设计
```go
// Command 接口定义了所有命令的通用接口
type Command interface {
    Execute()
}
```

### 2. 具体命令实现
```go
// StartCommand 实现了启动系统的命令
type StartCommand struct {
    mb *MotherBoard  // 持有接收者的引用
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
    return &StartCommand{
        mb: mb,
    }
}

func (c *StartCommand) Execute() {
    c.mb.Start()  // 委托给接收者执行
}

// RebootCommand 实现了重启系统的命令
type RebootCommand struct {
    mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
    return &RebootCommand{
        mb: mb,
    }
}

func (c *RebootCommand) Execute() {
    c.mb.Reboot()
}
```

### 3. 接收者实现
```go
// MotherBoard 是命令的接收者，真正执行操作的对象
type MotherBoard struct{}

func (*MotherBoard) Start() {
    fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
    fmt.Print("system rebooting\n")
}
```

### 4. 调用者实现
```go
// Box 是调用者，持有命令并在需要时执行
// 相当于遥控器，按钮可以配置不同的命令
type Box struct {
    button1 Command  // 按钮1对应的命令
    button2 Command  // 按钮2对应的命令
}

func NewBox(button1, button2 Command) *Box {
    return &Box{
        button1: button1,
        button2: button2,
    }
}

func (b *Box) PressButton1() {
    b.button1.Execute()  // 执行按钮1对应的命令
}

func (b *Box) PressButton2() {
    b.button2.Execute()  // 执行按钮2对应的命令
}
```

## 使用场景

### 1. GUI按钮和菜单项
每个按钮或菜单项都可以对应一个命令对象，实现点击时的操作。

### 2. 宏命令
可以将多个命令组合成一个宏命令，一次性执行多个操作。

### 3. 事务操作
命令模式可以用来实现事务，可以记录操作历史以便撤销。

### 4. 队列请求和日志请求
可以将命令对象存储在队列中，实现延迟执行或持久化。

### 5. 远程控制
通过网络发送命令对象，实现远程操作。

## 示例详解

本示例展示了如何通过命令模式实现机箱按钮的灵活配置：

1. **第一个机箱** `box1` 配置：
   - 按钮1：启动系统
   - 按钮2：重启系统

2. **第二个机箱** `box2` 配置：
   - 按钮1：重启系统
   - 按钮2：启动系统

这种配置方式体现了命令模式的灵活性，相同的硬件按钮可以绑定不同的操作，而不需要修改硬件本身。

## 优点

1. **解耦**：将发送者和接收者完全解耦，发送者不知道接收者的具体实现
2. **扩展性**：新增命令很容易，不需要修改现有代码
3. **组合命令**：可以组合多个命令形成复合命令
4. **支持撤销**：可以存储命令历史，支持撤销操作
5. **支持日志**：可以记录命令的执行历史
6. **宏命令**：可以将多个命令组合成一个宏命令

## 缺点

1. **类数量增加**：每个命令都需要一个具体类，可能导致类数量过多
2. **复杂性**：对于简单的操作，使用命令模式可能会增加不必要的复杂性
3. **性能开销**：封装命令对象会带来额外的性能开销

## 与其他模式的关系

### 命令模式 vs 策略模式
- **命令模式**：关注于将请求封装为对象，支持撤销、日志等功能
- **策略模式**：关注于算法的替换，通常不包含接收者信息

### 命令模式 vs 观察者模式
- **命令模式**：将请求封装为对象，通常是一对一的
- **观察者模式**：定义对象间的一对多依赖关系

### 命令模式 vs 备忘录模式
- **命令模式**：可以用于实现撤销功能
- **备忘录模式**：专门用于保存和恢复对象状态

## 实际应用示例

### 1. 文本编辑器操作
```go
// 复制命令
type CopyCommand struct {
    editor *TextEditor
    backup string
}

func (c *CopyCommand) Execute() {
    c.backup = c.editor.GetSelection()
    c.editor.Copy()
}

func (c *CopyCommand) Undo() {
    c.editor.SetSelection(c.backup)
}
```

### 2. 数据库事务
```go
// 数据库命令
type DatabaseCommand struct {
    db     *Database
    query  string
    backup []byte
}

func (c *DatabaseCommand) Execute() {
    c.backup = c.db.Backup()
    c.db.Execute(c.query)
}

func (c *DatabaseCommand) Undo() {
    c.db.Restore(c.backup)
}
```

### 3. 文件操作
```go
// 文件创建命令
type FileCreateCommand struct {
    path    string
    content []byte
}

func (c *FileCreateCommand) Execute() {
    os.WriteFile(c.path, c.content, 0644)
}

func (c *FileCreateCommand) Undo() {
    os.Remove(c.path)
}
```

### 4. 网络请求
```go
// HTTP请求命令
type HTTPCommand struct {
    client  *http.Client
    request *http.Request
    response *http.Response
}

func (c *HTTPCommand) Execute() error {
    resp, err := c.client.Do(c.request)
    c.response = resp
    return err
}
```

## 最佳实践

### 1. 添加撤销功能
```go
// 支持撤销的命令接口
type UndoableCommand interface {
    Command
    Undo()
}
```

### 2. 使用函数式编程简化
```go
// 使用函数类型简化命令创建
type CommandFunc func()

func (f CommandFunc) Execute() {
    f()
}

// 使用示例
startCmd := CommandFunc(func() {
    motherboard.Start()
})
```

### 3. 命令队列
```go
// 命令队列
type CommandQueue struct {
    commands []Command
}

func (q *CommandQueue) Add(cmd Command) {
    q.commands = append(q.commands, cmd)
}

func (q *CommandQueue) ExecuteAll() {
    for _, cmd := range q.commands {
        cmd.Execute()
    }
}
```

### 4. 宏命令
```go
// 宏命令组合多个命令
type MacroCommand struct {
    commands []Command
}

func (m *MacroCommand) Execute() {
    for _, cmd := range m.commands {
        cmd.Execute()
    }
}

func (m *MacroCommand) Add(cmd Command) {
    m.commands = append(m.commands, cmd)
}
```

## 测试策略

### 1. 单元测试
- 测试每个具体命令是否正确调用接收者
- 测试调用者的行为是否符合预期

### 2. 集成测试
- 测试整个命令链路的执行
- 验证命令与接收者的集成

### 3. Mock测试
- 使用mock对象模拟接收者
- 验证命令是否正确委托给接收者

```go
// 测试示例
func TestStartCommand(t *testing.T) {
    mockMB := &MockMotherBoard{}
    cmd := NewStartCommand(mockMB)
    
    cmd.Execute()
    
    if !mockMB.StartCalled {
        t.Error("Start should be called")
    }
}
```

## 扩展思考

### 1. 异步命令
实现支持异步执行的命令，可以在后台执行耗时操作。

### 2. 优先级命令
为命令添加优先级，实现优先级队列。

### 3. 条件命令
实现根据条件决定是否执行的命令。

### 4. 命令链
实现命令链，一个命令的执行可以触发下一个命令。

### 5. 命令缓存
实现命令缓存，避免重复创建相同的命令对象。

## 总结

命令模式通过将请求封装为对象，实现了请求发送者和接收者的解耦。这种模式特别适合需要支持撤销、日志记录、宏命令等功能的场景。在实际开发中，命令模式常用于GUI应用、事务系统、远程控制等场景，能够显著提高系统的灵活性和可扩展性。

通过合理运用命令模式，可以构建出更加灵活、可维护的系统架构，特别是在需要处理复杂用户交互和系统操作的场景中。这种设计模式不仅提供了代码的解耦，还为系统的功能扩展提供了强大的支持。

