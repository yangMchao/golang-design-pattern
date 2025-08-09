package products

import (
	"fmt"
	"golang-design-pattern/05_abstract_factory/interfaces"
)

// XMLMainDAO XML存储的OrderMainDAO实现
// 实现了interfaces.OrderMainDAO接口，用于XML文件中订单主记录的存储
type XMLMainDAO struct{}

// SaveOrderMain 将订单主记录保存到XML文件
// 实际应用中这里会包含具体的XML文件操作逻辑
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// Ensure XMLMainDAO implements OrderMainDAO interface at compile time
var _ interfaces.OrderMainDAO = (*XMLMainDAO)(nil)

// XMLDetailDAO XML存储的OrderDetailDAO实现
// 实现了interfaces.OrderDetailDAO接口，用于XML文件中订单详情记录的存储
type XMLDetailDAO struct{}

// SaveOrderDetail 将订单详情记录保存到XML文件
// 实际应用中这里会包含具体的XML文件操作逻辑
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

// Ensure XMLDetailDAO implements OrderDetailDAO interface at compile time
var _ interfaces.OrderDetailDAO = (*XMLDetailDAO)(nil)
