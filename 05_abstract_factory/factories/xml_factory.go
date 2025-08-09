package factories

import (
	"golang-design-pattern/05_abstract_factory/interfaces"
	"golang-design-pattern/05_abstract_factory/products"
)

// XMLDAOFactory XML存储的DAO工厂实现
// 实现了interfaces.DAOFactory接口，用于创建XML文件存储相关的DAO对象族
type XMLDAOFactory struct{}

// CreateOrderMainDAO 创建XML存储的订单主记录DAO
// 返回一个实现了OrderMainDAO接口的XMLMainDAO实例
func (*XMLDAOFactory) CreateOrderMainDAO() interfaces.OrderMainDAO {
	return &products.XMLMainDAO{}
}

// CreateOrderDetailDAO 创建XML存储的订单详情记录DAO
// 返回一个实现了OrderDetailDAO接口的XMLDetailDAO实例
func (*XMLDAOFactory) CreateOrderDetailDAO() interfaces.OrderDetailDAO {
	return &products.XMLDetailDAO{}
}

// Ensure XMLDAOFactory implements DAOFactory interface at compile time
var _ interfaces.DAOFactory = (*XMLDAOFactory)(nil)
