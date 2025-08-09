package products

import (
	"fmt"
	"golang-design-pattern/05_abstract_factory/interfaces"
)

// RDBMainDAO 关系型数据库的OrderMainDAO实现
// 实现了interfaces.OrderMainDAO接口，用于关系型数据库中订单主记录的存储
type RDBMainDAO struct{}

// SaveOrderMain 将订单主记录保存到关系型数据库
// 实际应用中这里会包含具体的数据库操作逻辑
func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// Ensure RDBMainDAO implements OrderMainDAO interface at compile time
var _ interfaces.OrderMainDAO = (*RDBMainDAO)(nil)

// RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
// 实现了interfaces.OrderDetailDAO接口，用于关系型数据库中订单详情记录的存储
type RDBDetailDAO struct{}

// SaveOrderDetail 将订单详情记录保存到关系型数据库
// 实际应用中这里会包含具体的数据库操作逻辑
func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// Ensure RDBDetailDAO implements OrderDetailDAO interface at compile time
var _ interfaces.OrderDetailDAO = (*RDBDetailDAO)(nil)
