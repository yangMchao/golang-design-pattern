# 装饰模式 (Decorator Pattern)

## 概述

装饰模式是一种结构型设计模式，它允许在运行时动态地向对象添加新的功能，而不改变其结构。该模式通过创建一个装饰类来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。

装饰模式相比继承提供了更加灵活的扩展方式，可以在不创建大量子类的情况下扩展对象的功能，避免了类爆炸问题。

## 核心概念

装饰模式的核心思想是通过组合而非继承来扩展对象功能：
- **透明性**：装饰后的对象与原始对象具有相同的接口，客户端可以透明地使用
- **动态扩展**：可以在运行时动态地添加或删除装饰
- **嵌套装饰**：支持多个装饰器的嵌套使用，实现功能的组合
- **单一职责**：每个装饰器只负责添加特定的功能

## 模式结构

### 1. 组件接口角色 (Component Interface)
- `Component` 接口：定义了被装饰对象和装饰器的公共接口
- 确保装饰器和被装饰对象具有相同的方法签名

### 2. 具体组件角色 (Concrete Component)
- `ConcreteComponent` 结构体：实现了组件接口的基础对象
- 提供最基本的功能实现

### 3. 装饰器角色 (Decorator)
- `MulDecorator` 结构体：乘法装饰器，为组件添加乘法功能
- `AddDecorator` 结构体：加法装饰器，为组件添加加法功能
- 都通过匿名组合嵌入了 `Component` 接口

### 4. 客户端角色 (Client)
- 负责创建组件对象和装饰器对象
- 通过组合不同的装饰器实现复杂功能

## 代码结构分析

### 组件接口设计
```go
type Component interface {
	Calc() int
}
```

### 具体组件实现
```go
type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}
```

### 装饰器实现 - 乘法装饰器
```go
type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}
```

### 装饰器实现 - 加法装饰器
```go
type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
```

### 客户端使用示例
```go
func ExampleDecorator() {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c, 10)  // 添加10
	c = WrapMulDecorator(c, 8)   // 乘以8
	res := c.Calc()
	
	fmt.Printf("res %d\n", res)  // 输出: res 80
}
```

## 使用场景

### 1. 动态功能扩展
当需要在运行时动态地为对象添加功能，而不是通过静态继承时，可以使用装饰模式。

### 2. 功能组合
需要支持多种功能组合，且这些功能可以任意搭配使用时，装饰模式提供了灵活的组合方式。

### 3. 避免类爆炸
当存在大量可能的组合方式时，使用继承会导致子类数量爆炸性增长，装饰模式可以有效避免这个问题。

### 4. 不改变原有代码
需要扩展第三方类库或系统类的功能，而又无法修改源码时，可以使用装饰模式进行扩展。

### 5. 撤销功能
需要支持功能的动态撤销，装饰器可以在运行时移除特定的装饰。

### 6. 日志和监控
为现有方法添加日志记录、性能监控、权限检查等横切关注点。

## 优点

1. **灵活性高**：可以在运行时动态添加或删除功能
2. **遵循开闭原则**：无需修改现有代码即可扩展功能
3. **避免类爆炸**：通过组合而非继承实现功能扩展
4. **单一职责**：每个装饰器只关注特定的功能扩展
5. **可组合性强**：多个装饰器可以嵌套使用，实现复杂功能
6. **透明性**：装饰后的对象与原始对象接口一致

## 缺点

1. **增加复杂性**：会引入许多小对象，增加系统复杂度
2. **调试困难**：多层装饰可能导致调试和错误定位困难
3. **性能开销**：装饰器链可能带来轻微的性能开销
4. **设计难度**：需要精心设计装饰器的接口和实现
5. **过度使用**：可能导致过度设计，简单场景使用继承更合适

## 与其他模式的关系

### 装饰模式 vs 继承
- **装饰模式**：运行时扩展，更加灵活，支持功能组合
- **继承**：编译时扩展，静态绑定，不支持运行时修改

### 装饰模式 vs 代理模式
- **装饰模式**：关注功能的增强和扩展
- **代理模式**：关注对对象的访问控制和管理

### 装饰模式 vs 适配器模式
- **装饰模式**：保持接口不变，增强功能
- **适配器模式**：改变接口，使其与其他接口兼容

### 装饰模式 vs 组合模式
- **装饰模式**：为单个对象添加功能
- **组合模式**：构建对象的层次结构

### 装饰模式 vs 策略模式
- **装饰模式**：改变对象的外观（功能）
- **策略模式**：改变对象的内部实现（算法）

## 实际应用示例

### 1. 文本格式化装饰器
```go
// 基础文本组件
type TextComponent interface {
	Render() string
}

type PlainText struct {
	text string
}

func (p *PlainText) Render() string {
	return p.text
}

// 粗体装饰器
type BoldDecorator struct {
	TextComponent
}

func (b *BoldDecorator) Render() string {
	return "<b>" + b.TextComponent.Render() + "</b>"
}

// 斜体装饰器
type ItalicDecorator struct {
	TextComponent
}

func (i *ItalicDecorator) Render() string {
	return "<i>" + i.TextComponent.Render() + "</i>"
}

// 使用示例
text := &PlainText{text: "Hello World"}
bold := &BoldDecorator{TextComponent: text}
boldItalic := &ItalicDecorator{TextComponent: bold}
fmt.Println(boldItalic.Render())  // 输出: <i><b>Hello World</b></i>
```

### 2. 日志装饰器
```go
// 基础服务接口
type Service interface {
	Execute() error
}

// 日志装饰器
type LoggingDecorator struct {
	Service
	logger *log.Logger
}

func (l *LoggingDecorator) Execute() error {
	l.logger.Printf("Starting service execution...")
	err := l.Service.Execute()
	if err != nil {
		l.logger.Printf("Service execution failed: %v", err)
	} else {
		l.logger.Printf("Service execution completed successfully")
	}
	return err
}
```

### 3. 缓存装饰器
```go
// 数据访问接口
type DataAccess interface {
	GetData(key string) (string, error)
}

// 缓存装饰器
type CacheDecorator struct {
	DataAccess
	cache map[string]string
}

func (c *CacheDecorator) GetData(key string) (string, error) {
	if value, exists := c.cache[key]; exists {
		return value, nil
	}
	
	value, err := c.DataAccess.GetData(key)
	if err == nil {
		c.cache[key] = value
	}
	return value, err
}
```

### 4. 权限检查装饰器
```go
// 基础操作接口
type Operation interface {
	Perform(user string) error
}

// 权限检查装饰器
type AuthorizationDecorator struct {
	Operation
	permissions map[string]bool
}

func (a *AuthorizationDecorator) Perform(user string) error {
	if !a.permissions[user] {
		return fmt.Errorf("user %s has no permission", user)
	}
	return a.Operation.Perform(user)
}
```

### 5. 性能监控装饰器
```go
// 方法调用接口
type Calculator interface {
	Calculate(input int) int
}

// 性能监控装饰器
type PerformanceDecorator struct {
	Calculator
}

func (p *PerformanceDecorator) Calculate(input int) int {
	start := time.Now()
	result := p.Calculator.Calculate(input)
	duration := time.Since(start)
	fmt.Printf("Calculation took %v\n", duration)
	return result
}
```

### 6. 输入验证装饰器
```go
// 处理器接口
type Handler interface {
	Handle(data string) error
}

// 输入验证装饰器
type ValidationDecorator struct {
	Handler
}

func (v *ValidationDecorator) Handle(data string) error {
	if data == "" {
		return fmt.Errorf("empty data not allowed")
	}
	return v.Handler.Handle(data)
}
```

## 总结

装饰模式通过组合而非继承的方式，为对象提供了一种灵活的功能扩展机制。该模式在保持接口不变的前提下，允许在运行时动态地为对象添加新的功能，避免了继承带来的类爆炸问题。

在Go语言中，利用接口和匿名组合可以优雅地实现装饰模式，使得代码结构清晰且易于维护。装饰模式特别适用于需要动态扩展功能、支持功能组合、或者需要为第三方类库添加功能的场景。

在实际应用中，装饰模式常与工厂模式、单例模式等其他设计模式结合使用，构建出更加灵活和可扩展的系统架构。
