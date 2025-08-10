# 模板方法模式 (Template Method Pattern)

## 概述

模板方法模式是一种行为型设计模式，它在一个方法中定义了一个算法的骨架，将某些步骤推迟到子类中实现。模板方法使得子类可以在不改变算法结构的情况下，重新定义算法中的某些步骤。

该模式通过继承机制实现，将通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现，使得实现符合开闭原则。

## 核心概念

模板方法模式的核心思想是：**定义一个操作中的算法骨架，而将一些步骤延迟到子类中**。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

在Go语言中，由于没有传统的继承机制，需要通过匿名组合来模拟实现继承。这里需要注意：因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有子类的引用。

## 模式结构

### 1. 抽象模板角色 (Abstract Template)
- `template` 结构体：定义模板方法的骨架，实现了算法的通用步骤
- `implement` 接口：定义需要子类实现的方法，作为模板和实现类之间的契约

### 2. 具体实现角色 (Concrete Implementations)
- `HTTPDownloader`：HTTP下载的具体实现，重写了`download`和`save`方法
- `FTPDownloader`：FTP下载的具体实现，重写了`download`方法，使用默认的`save`方法

### 3. 客户端角色 (Client)
- 通过`Downloader`接口使用模板方法，无需关心具体实现

## 代码结构分析

### 抽象模板设计
```go
type template struct {
    implement    // 嵌入实现接口，用于调用子类实现的方法
    uri string   // 存储下载URI
}

type implement interface {
    download()   // 需要子类实现的下载方法
    save()       // 需要子类实现的保存方法
}
```

### 模板方法实现
```go
func (t *template) Download(uri string) {
    t.uri = uri
    fmt.Print("prepare downloading\n")  // 通用步骤：准备下载
    t.implement.download()               // 子类实现的下载逻辑
    t.implement.save()                   // 子类实现的保存逻辑
    fmt.Print("finish downloading\n")    // 通用步骤：完成下载
}

// 默认实现，子类可以选择重写
func (t *template) save() {
    fmt.Print("default save\n")
}
```

### 具体实现类设计
```go
type HTTPDownloader struct {
    *template  // 匿名组合模板类，模拟继承
}

func (d *HTTPDownloader) download() {
    fmt.Printf("download %s via http\n", d.uri)  // HTTP特有的下载实现
}

func (*HTTPDownloader) save() {
    fmt.Printf("http save\n")  // HTTP特有的保存实现
}
```

## 使用场景

### 1. 算法骨架固定，步骤可变
当算法的整体结构固定不变，但某些步骤的具体实现可能变化时，使用模板方法模式。

### 2. 避免代码重复
多个类有相同的方法，并且逻辑相同，只有部分步骤的实现不同。

### 3. 控制子类扩展
父类通过模板方法控制子类的扩展，只允许子类在特定点进行扩展。

### 4. 框架设计
在框架设计中，框架提供算法的骨架，具体实现由应用程序完成。

## 优点

1. **代码复用**：将公共代码放在父类中，避免重复
2. **扩展性好**：子类可以通过重写特定步骤来扩展功能
3. **符合开闭原则**：在不修改父类的情况下，通过子类扩展功能
4. **控制反转**：父类控制子类的扩展点，实现控制反转
5. **便于维护**：算法的骨架在父类中，修改时只需要改一处

## 缺点

1. **类层次复杂**：需要继承结构，增加了类的层次
2. **子类受限**：子类必须遵守父类定义的算法骨架
3. **Go实现复杂**：在Go中需要通过接口和组合模拟继承，实现相对复杂
4. **灵活性降低**：某些步骤可能不需要，但子类必须实现

## 与其他模式的关系

### 模板方法模式 vs 策略模式
- **模板方法模式**：通过继承改变算法的部分步骤
- **策略模式**：通过组合改变整个算法

### 模板方法模式 vs 工厂方法模式
- **模板方法模式**：关注算法的骨架
- **工厂方法模式**：关注对象的创建

### 模板方法模式 vs 装饰器模式
- **模板方法模式**：通过继承扩展功能
- **装饰器模式**：通过组合扩展功能

## 实际应用示例

### 1. 数据库操作模板
```go
type DatabaseTemplate struct {
    implement
}

type implement interface {
    connect() error
    query(sql string) ([]Row, error)
    close() error
}

func (t *DatabaseTemplate) Execute(sql string) ([]Row, error) {
    if err := t.implement.connect(); err != nil {
        return nil, err
    }
    defer t.implement.close()
    
    return t.implement.query(sql)
}
```

### 2. 网络请求模板
```go
type HTTPClientTemplate struct {
    implement
}

type implement interface {
    prepareRequest() (*http.Request, error)
    handleResponse(resp *http.Response) error
}

func (t *HTTPClientTemplate) Do(url string) error {
    req, err := t.implement.prepareRequest()
    if err != nil {
        return err
    }
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    return t.implement.handleResponse(resp)
}
```

### 3. 文件处理模板
```go
type FileProcessorTemplate struct {
    implement
    filename string
}

type implement interface {
    open() error
    read() ([]byte, error)
    process(data []byte) ([]byte, error)
    write(data []byte) error
    close() error
}

func (t *FileProcessorTemplate) Process() error {
    if err := t.implement.open(); err != nil {
        return err
    }
    defer t.implement.close()
    
    data, err := t.implement.read()
    if err != nil {
        return err
    }
    
    processed, err := t.implement.process(data)
    if err != nil {
        return err
    }
    
    return t.implement.write(processed)
}
```

## 最佳实践

### 1. 合理划分抽象层次
将真正通用的步骤放在父类中，可能变化的步骤留给子类实现。

### 2. 使用接口定义钩子方法
通过接口定义需要子类实现的方法，提高灵活性。

### 3. 提供默认实现
对于某些步骤，父类可以提供默认实现，子类可以选择性地重写。

### 4. 避免过度设计
不要为了使用模板方法而强行创建继承层次，简单的组合可能更合适。

### 5. 考虑组合替代
在Go中，有时使用策略模式或组合模式可能比模板方法更合适。

### 6. 命名规范
- 模板方法通常命名为`TemplateMethod`或`Execute`
- 钩子方法使用具体的功能命名，如`download`、`save`等

## 测试策略

### 1. 单元测试
- 测试模板方法的算法骨架是否正确
- 测试各个钩子方法的正确性
- 测试默认实现的正确性

### 2. 集成测试
- 测试整个模板方法流程的正确性
- 验证不同实现类的行为差异

### 3. Mock测试
- 使用mock对象测试模板方法的调用顺序
- 验证钩子方法是否被正确调用

```go
func TestTemplateMethod(t *testing.T) {
    mockImpl := &mockImplement{}
    template := newTemplate(mockImpl)
    
    template.Download("test")
    
    // 验证调用顺序
    assert.Equal(t, []string{"prepare", "download", "save", "finish"}, mockImpl.calls)
}
```

## 扩展思考

### 1. 动态模板
可以根据运行时条件动态选择不同的实现步骤。

### 2. 参数化模板
通过参数控制模板方法的行为，而不是完全依赖子类实现。

### 3. 模板方法链
多个模板方法可以组合成更复杂的处理流程。

### 4. 并发模板
考虑在模板方法中处理并发操作，如并行下载、并行处理等。

### 5. 错误处理模板
在模板方法中统一处理错误，提供重试机制或回滚机制。

## 总结

模板方法模式通过定义算法骨架并将具体实现延迟到子类，实现了代码复用和扩展的完美结合。虽然在Go语言中需要通过接口和组合来模拟继承，但这种实现方式更加灵活，避免了传统继承带来的层次复杂性。

在实际开发中，模板方法模式特别适用于框架设计、业务流程处理、数据转换等场景。合理运用该模式可以显著提高代码的可维护性、可扩展性和复用性。