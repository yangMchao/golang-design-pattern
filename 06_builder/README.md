# 建造者模式 (Builder Pattern)

## 概述

建造者模式是一种创建型设计模式，它允许你逐步构建复杂对象。该模式将对象的构造过程与其表示分离，使得同样的构建过程可以创建不同的表示。建造者模式适用于创建复杂对象的场景，特别是当对象有多个组成部分且创建过程需要多个步骤时。

## 核心概念

建造者模式通过将复杂对象的构建过程封装在独立的建造者对象中，提供了一种更加灵活和可读的方式来创建复杂对象。它将一个复杂对象的构建与其表示分离，使得同样的构建过程可以创建不同的产品。

## 模式结构

### 1. 建造者角色 (Builder)
- `Builder` 接口：定义了创建产品各个部件的抽象方法
- `Builder1` 结构体：具体的建造者实现，用于构建字符串类型的产品
- `Builder2` 结构体：具体的建造者实现，用于构建整数类型的产品

### 2. 指挥者角色 (Director)
- `Director` 结构体：负责管理建造过程，按照特定的顺序调用建造者的方法来构建产品
- `Construct()` 方法：定义了构建过程的算法，按顺序调用各个建造步骤

### 3. 产品角色 (Product)
- 由具体建造者构建的最终产品，在示例中分别为字符串和整数类型的结果

## 代码结构分析

### 建造者接口设计
```go
// Builder 是生成器接口，定义了构建产品各个部件的抽象方法
type Builder interface {
    Part1()  // 构建第一部分
    Part2()  // 构建第二部分
    Part3()  // 构建第三部分
}
```

### 指挥者实现
```go
type Director struct {
    builder Builder  // 持有建造者接口
}

// NewDirector 创建新的指挥者实例
func NewDirector(builder Builder) *Director {
    return &Director{
        builder: builder,
    }
}

// Construct 定义了构建过程的算法
// 按照特定顺序调用建造者的方法来构建完整产品
func (d *Director) Construct() {
    d.builder.Part1()  // 第一步
    d.builder.Part2()  // 第二步
    d.builder.Part3()  // 第三步
}
```

### 具体建造者实现
```go
type Builder1 struct {
    result string  // 存储构建结果
}

// 实现各个构建步骤
func (b *Builder1) Part1() {
    b.result += "1"
}

func (b *Builder1) Part2() {
    b.result += "2"
}

func (b *Builder1) Part3() {
    b.result += "3"
}

// GetResult 获取最终构建的产品
func (b *Builder1) GetResult() string {
    return b.result
}
```

## 使用场景

### 1. 复杂对象构建
当对象有多个组成部分且构建过程复杂时，如构建包含多个配置选项的对象。

### 2. 不同表示的相同构建过程
当需要创建不同表示形式的对象，但构建过程相同时。

### 3. 避免构造函数参数过多
当构造函数参数过多，使用建造者模式可以提供更好的可读性。

### 4. 不可变对象构建
需要创建不可变对象，且对象状态需要在构建时确定。

### 5. 配置对象创建
创建具有大量可选配置的对象，如数据库连接配置、HTTP客户端配置等。

## 优点

1. **封装性好**：将复杂构建过程封装在建造者中，客户端无需知道具体细节
2. **扩展性强**：新增具体建造者无需修改现有代码，符合开闭原则
3. **精确控制**：可以更加精细地控制产品的构建过程
4. **解耦**：将构建过程与最终产品解耦，相同的构建过程可以创建不同的产品
5. **更好的可读性**：通过链式调用或分步骤构建，代码更加清晰

## 缺点

1. **增加复杂性**：需要定义多个类（建造者、指挥者、产品），增加了系统复杂度
2. **适用场景有限**：只适用于创建复杂对象，简单对象使用建造者模式反而会增加复杂性
3. **内部修改困难**：如果产品的内部结构发生变化，可能需要修改所有建造者
4. **额外开销**：相比直接使用构造函数，建造者模式有一定的性能开销

## 与其他模式的关系

### 建造者模式 vs 工厂模式
- **工厂模式**：关注创建单个产品，不关心创建过程
- **建造者模式**：关注复杂产品的构建过程，可以分步骤创建

### 建造者模式 vs 抽象工厂模式
- **抽象工厂模式**：创建产品族，创建的产品相互关联
- **建造者模式**：创建单一复杂产品，关注构建过程

### 建造者模式 vs 模板方法模式
- **模板方法模式**：定义算法骨架，子类实现具体步骤
- **建造者模式**：将构建算法委托给独立的对象（指挥者）

## 实际应用示例

### 1. HTTP客户端构建
```go
// HTTP客户端建造者
type HTTPClientBuilder struct {
    client *http.Client
}

func (b *HTTPClientBuilder) SetTimeout(timeout time.Duration) *HTTPClientBuilder {
    b.client.Timeout = timeout
    return b
}

func (b *HTTPClientBuilder) SetTransport(transport *http.Transport) *HTTPClientBuilder {
    b.client.Transport = transport
    return b
}

func (b *HTTPClientBuilder) Build() *http.Client {
    return b.client
}
```

### 2. 数据库连接配置
```go
type DatabaseConfigBuilder struct {
    config DatabaseConfig
}

func (b *DatabaseConfigBuilder) SetHost(host string) *DatabaseConfigBuilder {
    b.config.Host = host
    return b
}

func (b *DatabaseConfigBuilder) SetPort(port int) *DatabaseConfigBuilder {
    b.config.Port = port
    return b
}

func (b *DatabaseConfigBuilder) SetCredentials(username, password string) *DatabaseConfigBuilder {
    b.config.Username = username
    b.config.Password = password
    return b
}

func (b *DatabaseConfigBuilder) Build() DatabaseConfig {
    return b.config
}
```

### 3. 复杂配置对象
```go
type ServerConfigBuilder struct {
    config ServerConfig
}

func (b *ServerConfigBuilder) WithSSL(certFile, keyFile string) *ServerConfigBuilder {
    b.config.SSLEnabled = true
    b.config.SSLCert = certFile
    b.config.SSLKey = keyFile
    return b
}

func (b *ServerConfigBuilder) WithMiddleware(middlewares ...Middleware) *ServerConfigBuilder {
    b.config.Middlewares = append(b.config.Middlewares, middlewares...)
    return b
}

func (b *ServerConfigBuilder) Build() ServerConfig {
    return b.config
}
```

## 最佳实践

### 1. 使用链式调用
通过返回建造者自身实现方法链式调用，提高代码可读性。

### 2. 验证最终产品
在建造者的Build方法中验证最终产品的有效性，确保构建的产品符合要求。

### 3. 考虑不可变对象
对于配置类对象，考虑构建不可变对象以提高线程安全性。

### 4. 提供默认值
为可选参数提供合理的默认值，避免客户端必须设置所有参数。

### 5. 使用接口定义
通过接口定义建造者和产品，提高代码的可测试性和可扩展性。

### 6. 分离关注点
指挥者只负责构建顺序，具体建造者负责具体实现，保持职责单一。

## 测试策略

### 1. 单元测试
- 测试各个建造步骤是否正确执行
- 测试最终产品是否符合预期
- 测试指挥者的构建算法

### 2. 集成测试
- 测试整个建造过程的端到端功能
- 验证不同建造者创建的产品的兼容性

### 3. Mock测试
- 使用mock对象测试指挥者与建造者的交互
- 验证构建顺序是否正确

## 扩展思考

### 1. 流式建造者
支持流式构建，允许在构建过程中动态调整构建步骤。

### 2. 条件建造
根据运行时条件决定构建哪些部分或如何构建。

### 3. 并行建造
对于相互独立的构建步骤，可以考虑并行执行以提高性能。

### 4. 建造者工厂
使用工厂模式来创建具体的建造者实例。

### 5. 建造者与依赖注入
结合依赖注入框架，可以更灵活地配置建造者。

### 6. 建造者模式与函数式编程
使用函数式编程思想实现建造者模式，如使用函数选项模式。

## 总结

建造者模式通过将复杂对象的构建过程与其表示分离，提供了一种优雅的方式来创建复杂对象。它特别适用于需要创建具有多个配置选项或组成部分的对象的场景。在实际开发中，合理运用建造者模式可以显著提高代码的可维护性、可读性和灵活性。虽然建造者模式会增加一些额外的复杂性，但对于构建复杂对象来说，这些付出通常是值得的。