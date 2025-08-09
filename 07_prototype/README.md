# 原型模式 (Prototype Pattern)

## 概述

原型模式是一种创建型设计模式，它使对象能够复制自身，并且暴露到接口中。这种模式允许客户端在不知道接口实际对象类型的情况下生成新的对象，通过克隆现有实例来创建新的实例，而不是通过new关键字实例化。

原型模式特别适用于创建成本较高的复杂对象，通过复制预配置的原型实例来提高性能并简化对象创建过程。

## 核心概念

原型模式通过克隆现有对象来创建新对象，避免了重复初始化带来的开销。它包含一个原型管理器，用于存储和管理各种原型实例，客户端可以通过管理器获取预先配置好的原型副本。

## 模式结构

### 1. 原型接口 (Prototype Interface)
- `Cloneable` 接口：定义了克隆自身的方法
- 所有可以克隆的对象都需要实现这个接口

### 2. 原型管理器 (Prototype Manager)
- `PrototypeManager` 结构体：负责管理所有原型实例
- 提供注册和获取原型的功能
- 客户端通过管理器获取原型副本

### 3. 具体原型 (Concrete Prototype)
- `Type1`、`Type2` 等具体实现类型
- 实现了 `Cloneable` 接口的克隆方法
- 包含需要预配置的属性

## 代码结构分析

### 原型接口设计
```go
// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
    Clone() Cloneable
}
```

### 原型管理器实现
```go
type PrototypeManager struct {
    prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
    return &PrototypeManager{
        prototypes: make(map[string]Cloneable),
    }
}

func (p *PrototypeManager) Get(name string) Cloneable {
    return p.prototypes[name].Clone()
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
    p.prototypes[name] = prototype
}
```

### 具体原型实现
```go
type Type1 struct {
    name string
}

func (t *Type1) Clone() Cloneable {
    tc := *t
    return &tc
}
```

### 客户端使用示例
```go
// 创建管理器
manager := NewPrototypeManager()

// 注册原型
t1 := &Type1{name: "type1"}
manager.Set("t1", t1)

// 获取原型副本
clone := manager.Get("t1")
```

## 使用场景

### 1. 对象创建成本高
当对象创建涉及大量初始化操作或需要复杂计算时，使用原型模式可以避免重复创建开销。

### 2. 需要预配置的对象
当系统需要大量相似但略有不同的对象时，可以预先创建并配置原型，然后根据需要克隆和修改。

### 3. 运行时动态对象创建
在运行时不知道对象具体类型的情况下，需要通过接口创建对象实例。

### 4. 避免重复初始化
当对象初始化过程复杂且耗时，但创建后的对象状态相对稳定时。

### 5. 保护对象状态
需要创建对象的副本而不影响原始对象的状态。

## 优点

1. **性能提升**：避免重复的初始化开销，特别是复杂对象的创建
2. **简化对象创建**：隐藏对象创建的具体细节，客户端只需要调用Clone方法
3. **动态运行时配置**：可以在运行时动态地添加和删除原型
4. **减少子类构造**：不需要为每个产品类创建工厂类
5. **状态保护**：可以基于现有对象创建新对象，而不影响原对象

## 缺点

1. **深拷贝复杂性**：对于包含引用类型成员的对象，实现深拷贝可能很复杂
2. **循环引用问题**：对象图中存在循环引用时，克隆可能导致问题
3. **初始化灵活性降低**：每个克隆对象都继承原对象的状态，可能不适合所有场景
4. **接口限制**：需要实现Cloneable接口，对现有类可能不够灵活

## 与其他模式的关系

### 原型模式 vs 工厂方法模式
- **原型模式**：通过克隆现有实例创建新对象
- **工厂方法模式**：通过new关键字创建新实例

### 原型模式 vs 建造者模式
- **原型模式**：适合创建相似对象的多个变体
- **建造者模式**：适合创建复杂对象的多种表示

### 原型模式与单例模式
- 原型管理器通常使用单例模式实现，确保全局只有一个管理器实例

## 实际应用示例

### 1. 文档模板系统
```go
type DocumentTemplate struct {
    Title    string
    Content  string
    Style    string
    Header   *Header
    Footer   *Footer
}

func (d *DocumentTemplate) Clone() Cloneable {
    newDoc := *d
    newDoc.Header = d.Header.Clone()
    newDoc.Footer = d.Footer.Clone()
    return &newDoc
}
```

### 2. 游戏对象系统
```go
type GameCharacter struct {
    Name     string
    Health   int
    Position Vector3
    Skills   []Skill
}

func (g *GameCharacter) Clone() Cloneable {
    newChar := *g
    newChar.Skills = make([]Skill, len(g.Skills))
    copy(newChar.Skills, g.Skills)
    return &newChar
}
```

### 3. 配置对象系统
```go
type AppConfig struct {
    DatabaseURL string
    CacheSize   int
    Features    map[string]bool
    Timeout     time.Duration
}

func (a *AppConfig) Clone() Cloneable {
    newConfig := *a
    newConfig.Features = make(map[string]bool)
    for k, v := range a.Features {
        newConfig.Features[k] = v
    }
    return &newConfig
}
```

## 最佳实践

### 1. 实现深拷贝
对于包含引用类型成员的对象，确保实现正确的深拷贝逻辑：
```go
func (t *ComplexType) Clone() Cloneable {
    newObj := *t
    newObj.refField = t.refField.Clone() // 递归克隆引用字段
    return &newObj
}
```

### 2. 使用原型缓存
将常用的原型实例缓存起来，避免重复创建：
```go
var prototypeCache = make(map[string]*Prototype)

func GetPrototype(key string) *Prototype {
    if proto, exists := prototypeCache[key]; exists {
        return proto.Clone().(*Prototype)
    }
    return nil
}
```

### 3. 原型版本控制
为原型添加版本信息，便于管理和升级：
```go
type VersionedPrototype struct {
    Version string
    Data    interface{}
}

func (v *VersionedPrototype) Clone() Cloneable {
    newObj := *v
    // 深度克隆数据
    return &newObj
}
```

### 4. 原型组合
支持原型的组合和嵌套：
```go
type CompositePrototype struct {
    Parts []Cloneable
}

func (c *CompositePrototype) Clone() Cloneable {
    newComposite := &CompositePrototype{
        Parts: make([]Cloneable, len(c.Parts)),
    }
    for i, part := range c.Parts {
        newComposite.Parts[i] = part.Clone()
    }
    return newComposite
}
```

## 测试策略

### 1. 克隆正确性测试
验证克隆对象与原始对象的等价性：
```go
func TestCloneCorrectness(t *testing.T) {
    original := &Type1{name: "test"}
    clone := original.Clone().(*Type1)
    
    if original == clone {
        t.Error("Clone should be a different instance")
    }
    if original.name != clone.name {
        t.Error("Clone should have same data")
    }
}
```

### 2. 深拷贝测试
验证引用类型的深拷贝：
```go
func TestDeepClone(t *testing.T) {
    original := &ComplexType{
        refField: &SubType{value: 42},
    }
    clone := original.Clone().(*ComplexType)
    
    if original.refField == clone.refField {
        t.Error("Deep clone should create new reference")
    }
}
```

### 3. 原型管理器测试
验证管理器的注册和获取功能：
```go
func TestPrototypeManager(t *testing.T) {
    manager := NewPrototypeManager()
    proto := &Type1{name: "prototype"}
    
    manager.Set("key", proto)
    clone := manager.Get("key").(*Type1)
    
    if clone.name != "prototype" {
        t.Error("Manager should return correct prototype")
    }
}
```

## 扩展思考

### 1. 原型池模式
结合对象池模式，创建原型池来管理对象生命周期：
```go
type PrototypePool struct {
    prototypes map[string]*sync.Pool
}
```

### 2. 序列化克隆
使用序列化/反序列化实现深拷贝：
```go
func (p *Prototype) DeepClone() *Prototype {
    data, _ := json.Marshal(p)
    var clone Prototype
    json.Unmarshal(data, &clone)
    return &clone
}
```

### 3. 原型工厂
将原型模式与工厂模式结合：
```go
type PrototypeFactory struct {
    manager *PrototypeManager
}

func (f *PrototypeFactory) Create(name string) interface{} {
    return f.manager.Get(name)
}
```

### 4. 原型配置
通过配置文件定义原型：
```go
type PrototypeConfig struct {
    Name    string                 `json:"name"`
    Type    string                 `json:"type"`
    Config  map[string]interface{} `json:"config"`
}
```

## 总结

原型模式通过对象克隆机制提供了一种高效的对象创建方式，特别适合创建成本高或需要预配置的场景。它简化了复杂对象的创建过程，提高了系统性能，并为运行时对象创建提供了灵活性。在实际应用中，需要注意深拷贝的实现细节，合理设计原型管理器，并结合具体业务场景选择合适的实现策略。
