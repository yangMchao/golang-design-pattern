# 享元模式 (Flyweight Pattern)

## 概述

享元模式是一种结构型设计模式，它通过共享技术有效地支持大量细粒度对象的复用。该模式通过将对象的内部状态（intrinsic state）和外部状态（extrinsic state）分离，使得多个对象可以共享相同的内部状态，从而大幅减少内存使用和提高性能。

## 核心概念

享元模式的核心思想是共享细粒度对象，将对象的状态分为两类：
- **内部状态（Intrinsic State）**：存储在享元对象内部，不会随环境的改变而改变的共享部分
- **外部状态（Extrinsic State）**：随环境改变而改变的不可共享部分，由客户端负责存储和传递

## 模式结构

### 1. 享元工厂角色 (Flyweight Factory)
- `ImageFlyweightFactory` 结构体：负责创建和管理享元对象
- 确保享元对象的共享和复用，避免重复创建相同对象
- 提供获取享元对象的统一接口

### 2. 享元角色 (Flyweight)
- `ImageFlyweight` 结构体：包含可以共享的内部状态
- 存储图像数据等不变的信息
- 可以被多个对象共享使用

### 3. 客户端角色 (Client)
- `ImageViewer` 结构体：使用享元对象的客户端
- 维护外部状态（如显示位置、大小等可变属性）
- 通过享元工厂获取享元对象

## 代码结构分析

### 享元工厂设计
```go
type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}
```

### 享元获取机制
```go
func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}
	return image
}
```

### 享元对象结构
```go
type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}
```

### 客户端使用
```go
type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}
```

## 使用场景

### 1. 大量相似对象
当系统中存在大量相似或相同的对象，造成内存开销过大时，可以使用享元模式共享这些对象。

### 2. 字符串常量池
Java 中的字符串常量池、Python 中的小整数缓存等都是享元模式的应用。

### 3. 数据库连接池
数据库连接池通过共享连接对象，避免频繁创建和销毁连接。

### 4. 图形系统
在图形系统中，大量相同或相似的图形对象可以通过享元模式共享，减少内存占用。

### 5. 游戏开发
游戏中大量相同的纹理、模型等资源可以通过享元模式共享使用。

### 6. 文档编辑器
文档编辑器中，相同的字符、格式等可以通过享元模式共享。

## 优点

1. **减少内存使用**：通过共享对象，大幅减少内存中对象的数量
2. **提高性能**：减少对象创建和销毁的开销
3. **集中管理**：享元工厂统一管理对象，便于维护和扩展
4. **细粒度控制**：可以精确控制共享对象的粒度
5. **支持大量对象**：能够有效地支持大量细粒度对象

## 缺点

1. **增加复杂性**：需要分离内部状态和外部状态，增加了系统的复杂性
2. **运行时开销**：查找和共享对象需要额外的运行时开销
3. **状态管理困难**：需要仔细管理内部状态和外部状态
4. **线程安全**：在多线程环境下需要考虑线程安全问题
5. **设计难度**：需要良好的设计才能有效识别和分离状态

## 与其他模式的关系

### 享元模式 vs 单例模式
- **享元模式**：共享多个对象的内部状态，允许存在多个实例
- **单例模式**：确保一个类只有一个实例，提供全局访问点

### 享元模式 vs 原型模式
- **享元模式**：通过共享现有对象来减少内存使用
- **原型模式**：通过克隆现有对象来创建新对象

### 享元模式 vs 对象池模式
- **享元模式**：强调对象状态的共享
- **对象池模式**：强调对象的重用，通常不共享状态

### 享元模式 vs 组合模式
- **享元模式**：优化对象的存储和共享
- **组合模式**：构建对象的部分-整体层次结构

## 实际应用示例

### 1. 文本编辑器字符管理
```go
// 字符享元
type CharacterFlyweight struct {
	char     rune
	font     string
	size     int
	color    string
}

// 字符工厂
type CharacterFactory struct {
	characters map[string]*CharacterFlyweight
}

func (f *CharacterFactory) GetCharacter(char rune, font string, size int, color string) *CharacterFlyweight {
	key := fmt.Sprintf("%c-%s-%d-%s", char, font, size, color)
	if c, exists := f.characters[key]; exists {
		return c
	}
	
	newChar := &CharacterFlyweight{
		char:  char,
		font:  font,
		size:  size,
		color: color,
	}
	f.characters[key] = newChar
	return newChar
}
```

### 2. 游戏纹理管理
```go
// 纹理享元
type TextureFlyweight struct {
	name   string
	width  int
	height int
	data   []byte
}

// 纹理工厂
type TextureFactory struct {
	textures map[string]*TextureFlyweight
}

func (f *TextureFactory) GetTexture(name string) *TextureFlyweight {
	if texture, exists := f.textures[name]; exists {
		return texture
	}
	
	// 加载纹理数据
	data := loadTextureData(name)
	texture := &TextureFlyweight{
		name:   name,
		width:  256,
		height: 256,
		data:   data,
	}
	f.textures[name] = texture
	return texture
}
```

### 3. 数据库连接池
```go
// 数据库连接享元
type DBConnectionFlyweight struct {
	connectionString string
	conn             *sql.DB
}

// 连接工厂
type ConnectionPool struct {
	connections map[string]*DBConnectionFlyweight
	mu          sync.RWMutex
}

func (p *ConnectionPool) GetConnection(connStr string) (*sql.DB, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if conn, exists := p.connections[connStr]; exists {
		return conn.conn, nil
	}
	
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	
	p.connections[connStr] = &DBConnectionFlyweight{
		connectionString: connStr,
		conn:             db,
	}
	return db, nil
}
```

### 4. 图形对象管理
```go
// 图形享元
type ShapeFlyweight struct {
	typeName string
	vertices []float32
	indices  []uint32
}

// 图形工厂
type ShapeFactory struct {
	shapes map[string]*ShapeFlyweight
}

func (f *ShapeFactory) GetShape(typeName string) *ShapeFlyweight {
	if shape, exists := f.shapes[typeName]; exists {
		return shape
	}
	
	// 根据类型创建图形数据
	vertices, indices := createShapeData(typeName)
	shape := &ShapeFlyweight{
		typeName: typeName,
		vertices: vertices,
		indices:  indices,
	}
	f.shapes[typeName] = shape
	return shape
}
```

## 总结

享元模式通过对象共享技术有效地支持大量细粒度对象的复用，在需要创建大量相似对象的场景中具有重要价值。该模式通过分离内部状态和外部状态，既减少了内存占用，又保持了对象的灵活性。在系统设计中，合理应用享元模式可以显著提高系统性能和资源利用率，特别适用于游戏开发、图形系统、文本处理等领域。

