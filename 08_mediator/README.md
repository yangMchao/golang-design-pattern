# 中介者模式

中介者模式用一个中介对象来封装一系列的对象交互。中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。

在这个 `mediator` 包中实现了一个计算机系统中各个组件通过中介者进行交互的例子：

- `CDDriver`：光驱，负责读取光盘数据
- `CPU`：中央处理器，负责处理数据
- `VideoCard`：显卡，负责显示视频
- `SoundCard`：声卡，负责播放音频

## 组件协作流程

1. `CDDriver` 读取数据后，通过中介者通知 `CPU`
2. `CPU` 处理数据后，通过中介者通知 `VideoCard` 和 `SoundCard`
3. `VideoCard` 显示视频数据
4. `SoundCard` 播放音频数据

## 核心实现

### 单例中介者
```go
var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}
```

### 状态变更处理
```go
func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}
```

## 使用示例

```go
mediator := GetMediatorInstance()
mediator.CD = &CDDriver{}
mediator.CPU = &CPU{}
mediator.Video = &VideoCard{}
mediator.Sound = &SoundCard{}

// 触发协作流程
mediator.CD.ReadData()
```

## 设计优势

- **降低耦合度**：各组件不需要直接相互引用
- **集中管理**：所有交互逻辑集中在中介者中
- **易于扩展**：新增组件只需修改中介者的处理逻辑
- **职责清晰**：每个组件只关注自己的职责，通过中介者协调
