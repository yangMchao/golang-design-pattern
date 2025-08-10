# 迭代器模式 (Iterator Pattern)

## 概述

迭代器模式是一种行为型设计模式，它提供了一种方法来顺序访问聚合对象中的各个元素，而又不暴露其内部的表示。迭代器模式将遍历行为从聚合对象中分离出来，封装到一个独立的迭代器对象中。

## 核心概念

迭代器模式的核心思想是提供一种统一的方式来遍历不同类型的集合（数组、链表、树等），使客户端代码可以相同的方式处理不同的数据结构，而无需了解底层实现细节。

## 模式结构

### 1. 迭代器角色 (Iterator)
- `Iterator` 接口：定义了遍历集合所需的抽象方法
  - `First()`：将迭代器重置到集合的第一个元素
  - `IsDone()`：检查是否已经遍历完所有元素
  - `Next()`：获取集合中的下一个元素

### 2. 具体迭代器角色 (Concrete Iterator)
- `NumbersIterator` 结构体：实现了Iterator接口的具体迭代器
- 负责维护当前遍历的位置和状态
- 实现了具体的遍历算法

### 3. 聚合角色 (Aggregate)
- `Aggregate` 接口：定义了创建迭代器的抽象方法 `Iterator()`
- 表示一个可以被遍历的集合

### 4. 具体聚合角色 (Concrete Aggregate)
- `Numbers` 结构体：实现了Aggregate接口的具体集合
- 包含实际的业务数据（本例中为整数范围）
- 负责创建对应的具体迭代器实例

## 代码结构分析

### 迭代器接口设计
```go
// Iterator 定义了遍历集合的抽象接口
type Iterator interface {
    First()           // 重置到第一个元素
    IsDone() bool     // 检查是否遍历完成
    Next() interface{} // 获取下一个元素
}
```

### 聚合接口设计
```go
// Aggregate 定义了可遍历集合的抽象接口
type Aggregate interface {
    Iterator() Iterator  // 创建对应的迭代器
}
```

### 具体迭代器实现
```go
type NumbersIterator struct {
    numbers *Numbers  // 关联的集合对象
    next    int       // 当前遍历位置
}

// First 重置迭代器到集合开始位置
func (i *NumbersIterator) First() {
    i.next = i.numbers.start
}

// IsDone 检查是否已经遍历完所有元素
func (i *NumbersIterator) IsDone() bool {
    return i.next > i.numbers.end
}

// Next 获取下一个元素并推进遍历位置
func (i *NumbersIterator) Next() interface{} {
    if !i.IsDone() {
        next := i.next
        i.next++
        return next
    }
    return nil
}
```

### 具体聚合实现
```go
type Numbers struct {
    start, end int  // 整数范围 [start, end]
}

// Iterator 创建并返回对应的具体迭代器
func (n *Numbers) Iterator() Iterator {
    return &NumbersIterator{
        numbers: n,
        next:    n.start,
    }
}
```

### 遍历工具函数
```go
// IteratorPrint 使用迭代器遍历集合并打印元素
func IteratorPrint(i Iterator) {
    for i.First(); !i.IsDone(); {
        c := i.Next()
        fmt.Printf("%#v\n", c)
    }
}
```

## 使用场景

### 1. 统一遍历接口
当需要为不同类型的集合提供统一的遍历接口时，迭代器模式可以隐藏不同集合的内部结构差异。

### 2. 隐藏集合实现细节
当不希望暴露集合的内部数据结构时，迭代器模式可以提供一个抽象层。

### 3. 支持多种遍历方式
可以为同一个集合提供多种不同的遍历算法（如前序、中序、后序遍历）。

### 4. 惰性加载
迭代器可以支持惰性加载，只在需要时才获取下一个元素。

### 5. 数据流处理
在处理大规模数据或数据流时，迭代器模式可以避免一次性加载所有数据。

## 优点

1. **单一职责**：将遍历算法从集合类中分离出来，使集合类专注于数据存储
2. **开闭原则**：可以新增迭代器类型而不修改原有集合类
3. **统一接口**：为不同类型集合提供一致的遍历接口
4. **隐藏实现**：客户端无需了解集合的内部结构
5. **支持并发**：可以同时使用多个迭代器遍历同一个集合
6. **惰性计算**：只在需要时才计算下一个元素

## 缺点

1. **额外开销**：对于简单集合，使用迭代器可能增加额外的类和对象开销
2. **性能损失**：相比直接索引访问，迭代器可能带来轻微的性能损失
3. **使用复杂性**：对于简单场景，使用迭代器可能比直接遍历更复杂
4. **错误处理**：迭代器错误处理可能不够直观

## 与其他模式的关系

### 迭代器模式 vs 组合模式
- **迭代器模式**：关注如何遍历集合元素
- **组合模式**：关注如何构建树形结构的集合
- 两者可以结合使用：组合模式构建结构，迭代器模式提供遍历

### 迭代器模式 vs 访问者模式
- **迭代器模式**：提供遍历机制，不定义对元素的操作
- **访问者模式**：定义对元素的操作，通常需要遍历机制

### 迭代器模式 vs 生成器模式
- **迭代器模式**：基于现有集合进行遍历
- **生成器模式**：按需生成新的对象序列

## 实际应用示例

### 1. 文件系统遍历
```go
// 文件系统迭代器
type FileIterator interface {
    First()
    IsDone() bool
    Next() *FileInfo
}

// 目录树遍历
type DirectoryIterator struct {
    root     string
    current  string
    files    []string
    index    int
}
```

### 2. 数据库结果集
```go
// 数据库查询结果迭代器
type ResultSetIterator struct {
    rows     *sql.Rows
    current  Row
}

func (r *ResultSetIterator) Next() interface{} {
    if r.rows.Next() {
        var row Row
        r.rows.Scan(&row)
        return row
    }
    return nil
}
```

### 3. 网络数据流
```go
// 网络数据包迭代器
type NetworkPacketIterator struct {
    conn     net.Conn
    buffer   []byte
}

func (n *NetworkPacketIterator) Next() interface{} {
    // 从网络连接读取下一个数据包
    packet, _ := readPacket(n.conn)
    return packet
}
```

### 4. 树形结构遍历
```go
// 二叉树迭代器 - 支持中序遍历
type BinaryTreeIterator struct {
    root    *TreeNode
    stack   []*TreeNode
    current *TreeNode
}

func (b *BinaryTreeIterator) Next() interface{} {
    // 实现中序遍历算法
    for len(b.stack) > 0 || b.current != nil {
        if b.current != nil {
            b.stack = append(b.stack, b.current)
            b.current = b.current.Left
        } else {
            node := b.stack[len(b.stack)-1]
            b.stack = b.stack[:len(b.stack)-1]
            b.current = node.Right
            return node.Value
        }
    }
    return nil
}
```

## 最佳实践

### 1. 接口设计原则
- 保持迭代器接口简单，只包含必要的方法
- 考虑支持泛型以提高类型安全性

### 2. 错误处理
- 在Next()方法中返回错误信息
- 提供错误状态检查方法

```go
type Iterator interface {
    First() error
    IsDone() bool
    Next() (interface{}, error)
    Error() error
}
```

### 3. 并发安全
- 如果集合可能被多个goroutine访问，确保迭代器的并发安全
- 考虑使用读写锁或其他同步机制

### 4. 资源管理
- 确保迭代器正确释放资源（如文件句柄、数据库连接）
- 考虑实现`Close()`方法

### 5. 性能优化
- 对于频繁使用的迭代器，考虑对象池化
- 对于大集合，考虑支持分块遍历

## 测试策略

### 1. 单元测试
- 测试迭代器的基本功能：First, Next, IsDone
- 测试边界条件：空集合、单元素集合、大集合
- 测试错误处理：越界访问、并发访问

### 2. 集成测试
- 测试迭代器与具体集合的集成
- 验证遍历结果的正确性
- 测试多个迭代器同时访问同一集合

### 3. 性能测试
- 测试大集合的遍历性能
- 测试内存使用情况
- 测试并发场景下的性能表现

### 4. 示例测试
```go
func TestIteratorPattern(t *testing.T) {
    // 测试空集合
    empty := NewNumbers(1, 0)
    iter := empty.Iterator()
    assert.True(t, iter.IsDone())

    // 测试正常遍历
    numbers := NewNumbers(1, 5)
    iter = numbers.Iterator()
    
    expected := []int{1, 2, 3, 4, 5}
    actual := []int{}
    
    for iter.First(); !iter.IsDone(); {
        val := iter.Next().(int)
        actual = append(actual, val)
    }
    
    assert.Equal(t, expected, actual)
}
```

## 扩展思考

### 1. 双向迭代器
支持向前和向后遍历的双向迭代器：
```go
type BidirectionalIterator interface {
    Iterator
    HasPrevious() bool
    Previous() interface{}
}
```

### 2. 可重置迭代器
支持保存和恢复遍历状态的迭代器：
```go
type ResettableIterator interface {
    Iterator
    SaveState() int
    RestoreState(state int)
}
```

### 3. 过滤迭代器
在遍历过程中应用过滤条件的迭代器：
```go
type FilteringIterator struct {
    iterator    Iterator
    predicate   func(interface{}) bool
}

func (f *FilteringIterator) Next() interface{} {
    for !f.iterator.IsDone() {
        val := f.iterator.Next()
        if f.predicate(val) {
            return val
        }
    }
    return nil
}
```

### 4. 转换迭代器
在遍历过程中对元素进行转换：
```go
type MappingIterator struct {
    iterator Iterator
    mapper   func(interface{}) interface{}
}

func (m *MappingIterator) Next() interface{} {
    val := m.iterator.Next()
    return m.mapper(val)
}
```

### 5. 并行迭代器
支持并行处理集合元素的迭代器：
```go
type ParallelIterator struct {
    iterator Iterator
    workers  int
}

func (p *ParallelIterator) ForEach(fn func(interface{})) {
    // 使用goroutine池并行处理元素
}
```

## 总结

迭代器模式是处理集合遍历的强大工具，它通过将遍历算法从集合类中分离出来，提高了代码的复用性和灵活性。在Go语言中，迭代器模式可以通过接口优雅地实现，支持各种类型的集合结构。

在实际应用中，迭代器模式特别适合需要隐藏集合内部实现、支持多种遍历方式或处理大规模数据的场景。通过合理的设计和扩展，迭代器模式可以为复杂的数据处理提供简洁而强大的解决方案。