package factories

import (
	"golang-design-pattern/05_abstract_factory/interfaces"
	"golang-design-pattern/05_abstract_factory/products"
)

// RDBDAOFactory 关系型数据库的DAO工厂实现
// 实现了interfaces.DAOFactory接口，用于创建关系型数据库相关的DAO对象族
type RDBDAOFactory struct{}

// CreateOrderMainDAO 创建关系型数据库的订单主记录DAO
// 返回一个实现了OrderMainDAO接口的RDBMainDAO实例
func (*RDBDAOFactory) CreateOrderMainDAO() interfaces.OrderMainDAO {
	return &products.RDBMainDAO{}
}

// CreateOrderDetailDAO 创建关系型数据库的订单详情记录DAO
// 返回一个实现了OrderDetailDAO接口的RDBDetailDAO实例
func (*RDBDAOFactory) CreateOrderDetailDAO() interfaces.OrderDetailDAO {
	return &products.RDBDetailDAO{}
}

// Ensure RDBDAOFactory implements DAOFactory interface at compile time
var _ interfaces.DAOFactory = (*RDBDAOFactory)(nil)
