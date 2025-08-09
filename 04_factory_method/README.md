# 工厂方法模式 (Factory Method Pattern)

工厂方法模式是一种创建型设计模式，它提供了一种创建对象的接口，但由子类决定要实例化的类是哪一个。工厂方法让类的实例化推迟到子类中进行。

## 意图

- 定义一个用于创建对象的接口，让子类决定实例化哪个类
- 将对象的创建与使用分离，降低耦合度
- 提供一种扩展机制，可以在不修改现有代码的情况下引入新的产品类

## 适用场景

- 当一个类不知道它所必须创建的对象的类时
- 当一个类希望由它的子类来指定它所创建的对象时
- 当类将创建对象的职责委托给多个帮助子类中的某一个，并且希望将哪一个帮助子类是代理者这一信息局部化时

## 结构

### 参与者

- **Product（产品接口）**：定义工厂方法所创建的对象的接口
- **ConcreteProduct（具体产品）**：实现Product接口的具体类
- **Creator（创建者）**：声明工厂方法，该方法返回一个Product类型的对象
- **ConcreteCreator（具体创建者）**：重写工厂方法以返回一个ConcreteProduct实例

### 在Go中的实现特点

本示例展示了Go语言中工厂方法模式的简洁实现：

1. **接口定义**：
   - `Operator`：产品接口，定义了所有运算操作的基本行为
   - `OperatorFactory`：创建者接口，定义了创建Operator的方法

2. **组合而非继承**：
   - 使用结构体嵌套（匿名组合）和接口实现
   - 避免了复杂的继承层次，符合Go的设计理念

3. **延迟实例化**：
   - 具体的Operator实例由对应的工厂类创建
   - 客户端只依赖于抽象接口

## 代码分析

### 产品接口（Operator）
```go
type Operator interface {
    SetA(int)
    SetB(int)
    Result() int
}
```

### 创建者接口（OperatorFactory）
```go
type OperatorFactory interface {
    Create() Operator
}
```

### 基础实现（OperatorBase）
```go
type OperatorBase struct {
    a, b int
}

func (o *OperatorBase) SetA(a int) { o.a = a }
func (o *OperatorBase) SetB(b int) { o.b = b }
```

### 具体产品实现

#### 加法运算器（PlusOperator）
```go
type PlusOperator struct {
    *OperatorBase
}

func (o PlusOperator) Result() int {
    return o.a + o.b
}
```

#### 减法运算器（MinusOperator）
```go
type MinusOperator struct {
    *OperatorBase
}

func (o MinusOperator) Result() int {
    return o.a - o.b
}
```

### 具体工厂实现

#### 加法工厂（PlusOperatorFactory）
```go
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
    return &PlusOperator{
        OperatorBase: &OperatorBase{},
    }
}
```

#### 减法工厂（MinusOperatorFactory）
```go
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
    return &MinusOperator{
        OperatorBase: &OperatorBase{},
    }
}
```

### 使用示例
```go
// 创建加法工厂
plusFactory := &PlusOperatorFactory{}
plusOp := plusFactory.Create()
plusOp.SetA(10)
plusOp.SetB(5)
result := plusOp.Result()  // 返回 15

// 创建减法工厂
minusFactory := &MinusOperatorFactory{}
minusOp := minusFactory.Create()
minusOp.SetA(10)
minusOp.SetB(5)
result := minusOp.Result()  // 返回 5
```

## 优点

- **解耦**：将对象的创建与使用分离，降低耦合度
- **扩展性**：添加新产品类时无需修改现有代码，符合开闭原则
- **隐藏实现**：客户端只需要知道所需产品的接口，无需了解具体实现
- **灵活性**：可以通过配置文件等方式动态选择具体工厂

## 缺点

- **类数量增加**：每增加一个产品类就需要增加一个对应的工厂类，可能导致类爆炸
- **复杂度增加**：引入了额外的抽象层，增加了系统复杂度
- **理解难度**：需要理解工厂和产品之间的关系，增加了学习成本

## 与其他模式的关系

- **简单工厂模式**：工厂方法是简单工厂的进一步抽象和推广，解决了简单工厂违背开闭原则的问题
- **抽象工厂模式**：抽象工厂是工厂方法的升级版，用于创建产品族
- **模板方法模式**：工厂方法常被模板方法调用，用于创建算法中需要使用的对象

## Go语言特色

在Go语言中，工厂方法模式具有以下特点：

1. **接口驱动**：利用Go的接口机制，实现更灵活的工厂方法
2. **组合优于继承**：通过结构体嵌套实现代码复用，避免了复杂的继承层次
3. **零值安全**：工厂方法可以安全地返回nil值，处理错误情况
4. **并发安全**：可以设计并发安全的工厂实现
5. **简洁性**：Go的语法简洁，使得工厂方法的实现更加直观

## 实际应用场景

1. **数据库连接**：根据不同的数据库类型返回相应的连接对象
2. **日志记录器**：根据配置创建不同的日志记录器实例
3. **UI组件**：根据不同的平台创建相应的UI组件
4. **支付处理**：根据不同的支付方式创建相应的支付处理器
