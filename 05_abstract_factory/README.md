# 抽象工厂模式 (Abstract Factory Pattern)

抽象工厂模式是一种创建型设计模式，它提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。

## 意图

- 提供一个创建一系列相关或相互依赖对象的接口
- 封装产品对象的创建过程，使得客户端与具体产品实现解耦
- 确保同一产品族中的对象能够协同工作
- 提供一种切换产品族的机制，而不需要修改客户端代码

## 适用场景

- 系统需要独立于其产品的创建、组合和表示时
- 系统需要由多个产品系列中的一个来配置时
- 相关产品对象的系列需要一起使用时
- 需要提供一个产品类库，而只想显示它们的接口而不是实现时

## 结构

### 参与者

- **AbstractFactory（抽象工厂）**：声明创建抽象产品对象的操作接口
- **ConcreteFactory（具体工厂）**：实现创建具体产品对象的操作
- **AbstractProduct（抽象产品）**：为一类产品对象声明接口
- **ConcreteProduct（具体产品）**：定义一个将被相应的具体工厂创建的产品对象
- **Client（客户端）**：仅使用由AbstractFactory和AbstractProduct类声明的接口

### 在Go中的实现特点

本示例展示了Go语言中抽象工厂模式的简洁实现：

1. **接口驱动**：
   - 使用Go的接口机制定义抽象工厂和抽象产品
   - 接口的隐式实现使得代码更加灵活

2. **组合实现**：
   - 通过结构体组合实现具体工厂和产品
   - 避免了复杂的继承层次

3. **产品族概念**：
   - 每个具体工厂创建一整套相关的产品对象
   - 保证同一产品族内的对象兼容性

## 代码分析

### 抽象产品接口

#### 订单主记录接口
```go
type OrderMainDAO interface {
    SaveOrderMain()
}
```

#### 订单详情接口
```go
type OrderDetailDAO interface {
    SaveOrderDetail()
}
```

### 抽象工厂接口
```go
type DAOFactory interface {
    CreateOrderMainDAO() OrderMainDAO
    CreateOrderDetailDAO() OrderDetailDAO
}
```

### 关系型数据库产品族

#### 具体产品实现
```go
// RDB主订单存储实现
type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
    fmt.Print("rdb main save\n")
}

// RDB订单详情存储实现
type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() {
    fmt.Print("rdb detail save\n")
}
```

#### 具体工厂实现
```go
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
    return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
    return &RDBDetailDAO{}
}
```

### XML存储产品族

#### 具体产品实现
```go
// XML主订单存储实现
type XMLMainDAO struct{}

func (*XMLMainDAO) SaveOrderMain() {
    fmt.Print("xml main save\n")
}

// XML订单详情存储实现
type XMLDetailDAO struct{}

func (*XMLDetailDAO) SaveOrderDetail() {
    fmt.Print("xml detail save")
}
```

#### 具体工厂实现
```go
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
    return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
    return &XMLDetailDAO{}
}
```

### 使用示例

#### 客户端代码
```go
// 定义客户端使用函数
func getMainAndDetail(factory DAOFactory) {
    factory.CreateOrderMainDAO().SaveOrderMain()
    factory.CreateOrderDetailDAO().SaveOrderDetail()
}

// 使用RDB产品族
func ExampleRdbFactory() {
    var factory DAOFactory
    factory = &RDBDAOFactory{}
    getMainAndDetail(factory)
    // Output:
    // rdb main save
    // rdb detail save
}

// 使用XML产品族
func ExampleXmlFactory() {
    var factory DAOFactory
    factory = &XMLDAOFactory{}
    getMainAndDetail(factory)
    // Output:
    // xml main save
    // xml detail save
}
```

## 产品族概念

在本示例中，我们定义了两个产品族：

1. **RDB产品族**：使用关系型数据库进行数据存储
   - `RDBMainDAO`：关系型数据库的主订单存储
   - `RDBDetailDAO`：关系型数据库的订单详情存储

2. **XML产品族**：使用XML文件进行数据存储
   - `XMLMainDAO`：XML文件的主订单存储
   - `XMLDetailDAO`：XML文件的订单详情存储

每个产品族内部的对象是相互关联的，它们共享相同的存储策略（RDB或XML），确保数据的一致性和兼容性。

## 动态切换机制

抽象工厂模式的核心优势之一是能够动态切换产品族，而不需要修改客户端代码：

```go
// 根据配置选择工厂
func createFactory(storageType string) DAOFactory {
    switch storageType {
    case "rdb":
        return &RDBDAOFactory{}
    case "xml":
        return &XMLDAOFactory{}
    default:
        return &RDBDAOFactory{} // 默认使用RDB
    }
}

// 使用示例
factory := createFactory("xml") // 只需改变参数即可切换存储方式
getMainAndDetail(factory)
```

## 优点

- **产品族一致性**：确保同一产品族中的对象能够协同工作
- **解耦**：客户端与具体产品实现解耦，只依赖于抽象接口
- **易于交换产品族**：只需改变具体工厂即可使用不同的产品配置
- **支持新产品族**：符合开闭原则，添加新的产品族不需要修改现有代码
- **封装性**：封装了产品对象的创建过程，隐藏实现细节

## 缺点

- **难以支持新种类的产品**：如果需要添加新的产品类型，需要修改抽象工厂接口和所有具体工厂类
- **系统复杂度增加**：引入了多个抽象层，增加了系统的复杂度
- **类爆炸**：每个产品族都需要对应的具体工厂和具体产品类
- **运行时开销**：通过抽象层创建对象可能带来轻微的性能开销

## 与工厂方法模式的区别

| 特性 | 工厂方法模式 | 抽象工厂模式 |
|------|--------------|--------------|
| **产品数量** | 创建单一产品 | 创建产品族（多个相关产品） |
| **复杂度** | 相对简单 | 更复杂 |
| **扩展性** | 易于添加新产品类型 | 易于添加新产品族 |
| **使用场景** | 单一产品创建 | 相关产品族的创建 |
| **Go实现** | 单个工厂接口 | 多个相关工厂方法 |

## 与其他模式的关系

- **工厂方法模式**：抽象工厂通常使用工厂方法来实现，每个工厂方法负责创建一个具体产品
- **单例模式**：具体工厂通常实现为单例，确保全局只有一个工厂实例
- **原型模式**：抽象工厂可以通过原型模式来创建产品对象，提高创建效率
- **建造者模式**：当产品族中的对象构造过程复杂时，可以结合建造者模式

## Go语言特色

在Go语言中，抽象工厂模式具有以下特点：

1. **接口隐式实现**：不需要显式声明实现关系，降低耦合度
2. **零值安全**：Go的零值特性使得工厂方法可以安全地返回nil
3. **并发安全**：可以设计并发安全的工厂实现
4. **简洁语法**：Go的语法简洁，使得抽象工厂的实现更加直观
5. **组合优于继承**：通过结构体组合实现具体产品，避免了复杂的继承层次

## 实际应用场景

1. **跨平台UI库**：为不同操作系统创建一致的UI组件族
2. **数据库访问层**：为不同数据库类型创建相应的DAO对象族
3. **文件处理系统**：为不同文件格式（XML、JSON、CSV等）创建相应的读写器族
4. **网络通信**：为不同协议（HTTP、TCP、UDP）创建相应的通信组件族
5. **游戏引擎**：为不同平台（PC、移动设备、控制台）创建相应的渲染组件族
