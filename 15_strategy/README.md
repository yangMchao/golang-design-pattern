# 策略模式 (Strategy Pattern)

## 概述

策略模式是一种行为型设计模式，它定义了一系列算法，并将每个算法封装起来，使它们可以互相替换。策略模式让算法的变化独立于使用算法的客户端，符合开闭原则。

该模式通过组合机制实现，将算法的定义与使用分离，使得算法可以独立于客户端变化。在Go语言中，策略模式通过接口实现，具有天然的灵活性和扩展性。

## 核心概念

策略模式的核心思想是：**定义一系列算法，把它们一个个封装起来，并且使它们可以相互替换**。策略模式使得算法可以独立于使用它的客户端而变化。

在Go语言中，策略模式通过接口和组合实现，避免了传统继承带来的层次复杂性，提供了更加灵活的算法替换机制。

## 模式结构

### 1. 策略接口角色 (Strategy Interface)
- `PaymentStrategy` 接口：定义所有支持的支付算法的公共接口
- 所有具体策略都必须实现这个接口

### 2. 具体策略角色 (Concrete Strategies)
- `Cash`：现金支付策略的具体实现
- `Bank`：银行转账支付策略的具体实现
- 可以根据需要添加更多的支付策略（如支付宝、微信支付等）

### 3. 上下文角色 (Context)
- `Payment` 结构体：使用策略的上下文环境
- 持有一个策略接口的引用，可以在运行时切换不同的策略

### 4. 支付上下文数据结构
- `PaymentContext`：封装支付相关的数据（姓名、卡号、金额等）

## 代码结构分析

### 策略接口设计
```go
type PaymentStrategy interface {
    Pay(*PaymentContext)
}
```
策略接口定义了所有支付策略的公共行为，每个具体策略都必须实现这个接口。

### 上下文设计
```go
type Payment struct {
    context  *PaymentContext  // 支付相关的数据
    strategy PaymentStrategy  // 支付策略接口引用
}

type PaymentContext struct {
    Name, CardID string
    Money        int
}
```
上下文负责维护对策略对象的引用，并在需要时将策略相关的数据传递给策略对象。

### 具体策略实现
```go
// 现金支付策略
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
    fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

// 银行转账策略
type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
    fmt.Printf("Pay $%d to %s by bank account %s", 
        ctx.Money, ctx.Name, ctx.CardID)
}
```
每个具体策略都实现了`PaymentStrategy`接口，提供了特定算法的实现。

### 客户端使用
```go
// 使用现金支付策略
payment := NewPayment("Ada", "", 123, &Cash{})
payment.Pay()

// 使用银行转账策略
payment := NewPayment("Bob", "0002", 888, &Bank{})
payment.Pay()
```
客户端通过创建`Payment`实例并传入不同的策略对象来使用不同的支付算法。

## 使用场景

### 1. 多种算法变体
当一个系统需要多种算法变体，并且需要在运行时动态选择时，使用策略模式。

### 2. 避免条件语句
当代码中存在大量的条件语句（如if-else或switch-case）来选择不同行为时，可以使用策略模式消除这些条件语句。

### 3. 算法需要独立于客户端
当算法的实现细节不应该暴露给客户端时，策略模式提供了良好的封装。

### 4. 需要动态切换算法
当系统需要在运行时根据不同条件选择不同算法时，策略模式提供了灵活的机制。

### 5. 避免继承层次过深
当使用继承会导致类层次结构过于复杂时，策略模式通过组合提供了更灵活的替代方案。

## 优点

1. **算法可互换**：可以在运行时动态切换算法
2. **避免条件语句**：消除大量的if-else或switch-case语句
3. **提高可扩展性**：添加新策略无需修改现有代码，符合开闭原则
4. **封装性好**：策略的实现细节对客户端透明
5. **复用性强**：策略可以在不同的上下文中复用
6. **测试方便**：每个策略可以独立测试
7. **Go实现简洁**：利用接口实现，天然支持策略模式

## 缺点

1. **类数量增加**：每个策略都是一个类，可能导致类数量过多
2. **客户端必须了解策略**：客户端需要知道不同策略的区别才能选择合适的策略
3. **策略间通信困难**：如果策略间需要共享数据，实现会变得复杂
4. **性能开销**：频繁创建和切换策略对象可能带来性能开销
5. **调试困难**：策略的动态切换可能使调试变得困难

## 与其他模式的关系

### 策略模式 vs 模板方法模式
- **策略模式**：通过组合改变整个算法，支持运行时切换
- **模板方法模式**：通过继承改变算法的部分步骤，编译时确定

### 策略模式 vs 工厂方法模式
- **策略模式**：关注算法的封装和替换
- **工厂方法模式**：关注对象的创建

### 策略模式 vs 状态模式
- **策略模式**：客户端主动选择策略
- **状态模式**：状态对象控制状态转换

### 策略模式 vs 装饰器模式
- **策略模式**：替换整个算法
- **装饰器模式**：在不改变对象接口的情况下动态添加功能

## 实际应用扩展

### 1. 支付系统扩展
```go
// 支付宝支付策略
type Alipay struct{}
func (*Alipay) Pay(ctx *PaymentContext) {
    fmt.Printf("Pay $%d to %s by Alipay", ctx.Money, ctx.Name)
}

// 微信支付策略
type WechatPay struct{}
func (*WechatPay) Pay(ctx *PaymentContext) {
    fmt.Printf("Pay $%d to %s by WeChat Pay", ctx.Money, ctx.Name)
}
```

### 2. 策略选择器
```go
type StrategySelector struct {
    strategies map[string]PaymentStrategy
}

func (s *StrategySelector) Select(strategyType string) PaymentStrategy {
    return s.strategies[strategyType]
}
```

### 3. 缓存策略
策略模式常用于实现缓存策略（LRU、FIFO、LFU等）。

### 4. 排序算法
可以根据数据特征选择不同的排序算法（快速排序、归并排序、堆排序等）。

## 总结

策略模式通过定义一系列算法，并将每个算法封装起来，使它们可以互相替换，为系统提供了灵活的算法选择机制。在Go语言中，策略模式的实现特别简洁，利用接口机制天然支持策略的替换和扩展。

该模式特别适用于需要动态选择算法、避免条件语句、提高代码可扩展性的场景。合理运用策略模式可以显著提高代码的可维护性、可扩展性和复用性，是现代软件开发中不可或缺的设计模式之一。