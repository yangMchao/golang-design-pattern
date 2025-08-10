# 组合模式 (Composite Pattern)

## 概述

组合模式是一种结构型设计模式，它允许你将对象组合成树形结构来表现"整体-部分"的层次结构。组合模式使得客户端对单个对象和组合对象的使用具有一致性，统一了对象和对象集合的接口。

## 核心概念

组合模式的核心思想是通过抽象出统一的组件接口，让叶子节点（单个对象）和组合节点（对象集合）都实现这个接口。这样客户端就可以用同样的方式处理单个对象和组合对象，无需区分它们的具体类型。

## 模式结构

### 1. 组件角色 (Component)
- `Component` 接口：定义了组合对象和叶子对象的公共接口
  - `Parent()`：获取父节点
  - `SetParent(Component)`：设置父节点
  - `Name()`：获取节点名称
  - `SetName(string)`：设置节点名称
  - `AddChild(Component)`：添加子节点
  - `Print(string)`：打印节点信息

### 2. 叶子角色 (Leaf)
- `Leaf` 结构体：表示叶子节点，没有子节点
- 实现了组件接口的基本行为
- 叶子节点是组合结构的基本单元

### 3. 组合角色 (Composite)
- `Composite` 结构体：表示组合节点，可以包含子节点
- 实现了组件接口的所有行为
- 负责管理子节点的添加和遍历

### 4. 工厂方法
- `NewComponent(kind int, name string) Component`：根据类型创建组件
- `NewLeaf()`：创建叶子节点
- `NewComposite()`：创建组合节点

## 代码结构分析

### 组件接口设计
```go
// Component 定义了组合对象和叶子对象的公共接口
type Component interface {
    Parent() Component
    SetParent(Component)
    Name() string
    SetName(string)
    AddChild(Component)
    Print(string)
}
```

### 基础组件实现
```go
// component 提供了Component接口的基础实现
type component struct {
    parent Component
    name   string
}

func (c *component) Parent() Component {
    return c.parent
}

func (c *component) SetParent(parent Component) {
    c.parent = parent
}

func (c *component) Name() string {
    return c.name
}

func (c *component) SetName(name string) {
    c.name = name
}

// 叶子节点和组合节点可以嵌入这个基础结构
func (c *component) AddChild(Component) {}
func (c *component) Print(string) {}
```

### 叶子节点实现
```go
// Leaf 表示叶子节点，没有子节点
type Leaf struct {
    component  // 嵌入基础组件
}

func NewLeaf() *Leaf {
    return &Leaf{}
}

// Print 实现叶子节点的打印逻辑
func (c *Leaf) Print(pre string) {
    fmt.Printf("%s-%s\n", pre, c.Name())
}
```

### 组合节点实现
```go
// Composite 表示组合节点，可以包含子节点
type Composite struct {
    component        // 嵌入基础组件
    childs []Component  // 子节点集合
}

func NewComposite() *Composite {
    return &Composite{
        childs: make([]Component, 0),
    }
}

// AddChild 添加子节点，并设置父节点关系
func (c *Composite) AddChild(child Component) {
    child.SetParent(c)
    c.childs = append(c.childs, child)
}

// Print 实现组合节点的打印逻辑（递归遍历）
func (c *Composite) Print(pre string) {
    fmt.Printf("%s+%s\n", pre, c.Name())
    pre += " "
    for _, comp := range c.childs {
        comp.Print(pre)
    }
}
```

### 工厂方法
```go
const (
    LeafNode = iota
    CompositeNode
)

// NewComponent 根据类型创建对应的组件
func NewComponent(kind int, name string) Component {
    var c Component
    switch kind {
    case LeafNode:
        c = NewLeaf()
    case CompositeNode:
        c = NewComposite()
    }
    c.SetName(name)
    return c
}
```

## 使用场景

### 1. 文件系统
文件和文件夹的统一管理，文件夹可以包含文件和子文件夹。

### 2. 图形界面
窗口和控件的树形结构，容器控件可以包含其他控件。

### 3. 组织机构
部门和员工的层次结构，部门可以包含子部门和员工。

### 4. 产品分类
商品分类的树形结构，分类可以包含子分类和具体商品。

### 5. 菜单系统
菜单和菜单项的层次结构，子菜单可以包含菜单项和其他子菜单。

## 优点

1. **统一接口**：客户端可以一致地使用单个对象和组合对象
2. **简化客户端代码**：客户端无需区分叶子节点和组合节点
3. **易于扩展**：新增节点类型不影响现有代码
4. **符合开闭原则**：可以在不修改现有代码的情况下添加新的组件类型
5. **清晰的层次结构**：自然地表达了对象的层次关系
6. **递归处理**：可以方便地实现递归操作

## 缺点

1. **设计复杂性**：需要仔细设计组件接口，确保适用于所有子类
2. **类型检查困难**：可能需要运行时类型检查来区分叶子节点和组合节点
3. **性能开销**：递归操作可能带来性能开销
4. **限制灵活性**：某些操作可能不适合所有组件类型
5. **内存占用**：需要额外的内存来维护层次结构

## 与其他模式的关系

### 组合模式 vs 装饰器模式
- **组合模式**：关注对象的组合和层次结构
- **装饰器模式**：关注对象功能的动态扩展
- 可以结合使用：组合模式构建结构，装饰器模式扩展功能

### 组合模式 vs 迭代器模式
- **组合模式**：关注如何构建和表示树形结构
- **迭代器模式**：关注如何遍历集合元素
- 两者可以结合使用：组合模式构建结构，迭代器模式提供遍历

### 组合模式 vs 访问者模式
- **组合模式**：提供树形结构的构建和基本操作
- **访问者模式**：定义对树形结构中元素的新操作
- 访问者模式可以为组合结构添加新的操作而不修改组件类

### 组合模式 vs 责任链模式
- **组合模式**：构建树形层次结构
- **责任链模式**：处理请求沿着链传递
- 组合结构的父节点引用可以形成责任链

## 实际应用示例

### 1. 文件系统设计
```go
// 文件系统组件接口
type FileSystem interface {
    Name() string
    Size() int64
    IsDir() bool
    Children() []FileSystem
    Add(FileSystem)
    Remove(FileSystem)
}

// 文件（叶子节点）
type File struct {
    name string
    size int64
}

// 文件夹（组合节点）
type Directory struct {
    name     string
    children []FileSystem
}
```

### 2. GUI控件系统
```go
// 控件接口
type Widget interface {
    Draw()
    SetBounds(x, y, width, height int)
    AddChild(Widget)
    RemoveChild(Widget)
    GetChildren() []Widget
}

// 按钮控件（叶子节点）
type Button struct {
    label string
}

// 面板控件（组合节点）
type Panel struct {
    children []Widget
}
```

### 3. 组织机构管理
```go
// 组织成员接口
type OrgMember interface {
    GetName() string
    GetSalary() float64
    AddSubordinate(OrgMember)
    RemoveSubordinate(OrgMember)
    GetSubordinates() []OrgMember
    GetSuperior() OrgMember
}

// 员工（叶子节点）
type Employee struct {
    name   string
    salary float64
}

// 部门（组合节点）
type Department struct {
    name         string
    subordinates []OrgMember
    superior     OrgMember
}
```

### 4. 产品分类系统
```go
// 产品分类接口
type ProductCategory interface {
    GetName() string
    GetPrice() float64
    AddProduct(ProductCategory)
    RemoveProduct(ProductCategory)
    GetProducts() []ProductCategory
    ApplyDiscount(float64)
}

// 具体产品（叶子节点）
type ConcreteProduct struct {
    name  string
    price float64
}

// 产品分类（组合节点）
type Category struct {
    name     string
    products []ProductCategory
}
```

### 5. 菜单系统
```go
// 菜单项接口
type MenuItem interface {
    Display()
    GetPrice() float64
    AddItem(MenuItem)
    RemoveItem(MenuItem)
    GetItems() []MenuItem
}

// 菜品（叶子节点）
type Dish struct {
    name  string
    price float64
}

// 菜单分类（组合节点）
type MenuCategory struct {
    name  string
    items []MenuItem
}
```

## 最佳实践

### 1. 接口设计原则
- 保持接口简洁，只包含必要的方法
- 确保接口适用于所有组件类型
- 考虑使用空接口方法避免运行时错误

### 2. 安全性考虑
- 在叶子节点中实现AddChild等方法时返回错误或空操作
- 提供类型检查方法帮助客户端判断节点类型

```go
func (c *Leaf) AddChild(child Component) error {
    return fmt.Errorf("leaf nodes cannot have children")
}

func (c *Leaf) IsLeaf() bool {
    return true
}
```

### 3. 性能优化
- 使用缓存来存储计算结果（如总大小、总价格）
- 考虑使用对象池来减少内存分配
- 对于频繁操作，考虑使用更高效的数据结构

### 4. 内存管理
- 注意避免循环引用导致的内存泄漏
- 提供清理方法来释放资源
- 考虑使用弱引用来避免内存泄漏

### 5. 并发安全
- 如果组合结构可能被多个goroutine访问，确保线程安全
- 考虑使用读写锁来提高并发性能
- 提供原子操作来更新组合结构

```go
type SafeComposite struct {
    mu       sync.RWMutex
    children []Component
}

func (c *SafeComposite) AddChild(child Component) {
    c.mu.Lock()
    defer c.mu.Unlock()
    child.SetParent(c)
    c.children = append(c.children, child)
}
```

## 测试策略

### 1. 单元测试
- 测试叶子节点的基本功能
- 测试组合节点的添加、删除、遍历功能
- 测试边界条件：空组合、单个子节点、循环引用

### 2. 集成测试
- 测试整个树形结构的构建和遍历
- 验证递归操作的正确性
- 测试内存使用和性能表现

### 3. 异常测试
- 测试非法操作（在叶子节点上添加子节点）
- 测试空指针和边界条件
- 测试并发访问的安全性

### 4. 性能测试
- 测试大树的构建和遍历性能
- 测试内存使用效率
- 测试并发场景下的性能表现

### 5. 示例测试
```go
func TestCompositePattern(t *testing.T) {
    // 测试叶子节点
    leaf := NewComponent(LeafNode, "leaf1")
    assert.Equal(t, "leaf1", leaf.Name())
    
    // 测试组合节点
    composite := NewComponent(CompositeNode, "root")
    composite.AddChild(leaf)
    assert.Equal(t, 1, len(composite.(*Composite).childs))
    
    // 测试树形结构
    root := NewComponent(CompositeNode, "root")
    child1 := NewComponent(CompositeNode, "child1")
    child2 := NewComponent(LeafNode, "child2")
    
    root.AddChild(child1)
    root.AddChild(child2)
    
    assert.Equal(t, 2, len(root.(*Composite).childs))
    assert.Equal(t, root, child1.Parent())
}
```

## 扩展思考

### 1. 双向遍历
支持从子节点到父节点的反向遍历：
```go
type BidirectionalComponent interface {
    Component
    GetPath() string
    GetDepth() int
    GetRoot() Component
}
```

### 2. 迭代器支持
为组合结构提供迭代器模式支持：
```go
type ComponentIterator interface {
    HasNext() bool
    Next() Component
    GetCurrent() Component
}

// 支持多种遍历方式：深度优先、广度优先
func (c *Composite) DepthFirstIterator() ComponentIterator
func (c *Composite) BreadthFirstIterator() ComponentIterator
```

### 3. 访问者模式集成
为组合结构添加访问者模式支持：
```go
type ComponentVisitor interface {
    VisitLeaf(leaf *Leaf)
    VisitComposite(composite *Composite)
}

type Component interface {
    Accept(visitor ComponentVisitor)
}
```

### 4. 事件机制
为组合结构添加事件通知机制：
```go
type ComponentEvent struct {
    Type   string
    Source Component
    Data   interface{}
}

type Component interface {
    AddEventListener(func(ComponentEvent))
    RemoveEventListener(func(ComponentEvent))
    FireEvent(ComponentEvent)
}
```

### 5. 序列化支持
为组合结构添加序列化和反序列化支持：
```go
type SerializableComponent interface {
    Component
    MarshalJSON() ([]byte, error)
    UnmarshalJSON(data []byte) error
}
```

## 总结

组合模式是构建树形层次结构的强大工具，它通过统一单个对象和组合对象的接口，简化了客户端代码的复杂性。在Go语言中，通过接口和结构体的组合可以优雅地实现组合模式，支持各种类型的层次结构构建。

在实际应用中，组合模式特别适合需要表达"整体-部分"关系的场景，如文件系统、GUI控件、组织机构等。通过合理的设计和扩展，组合模式可以为复杂的层次结构提供简洁而强大的解决方案，同时保持代码的可维护性和可扩展性。关键在于正确设计组件接口，确保既能满足当前需求，又能支持未来的扩展。