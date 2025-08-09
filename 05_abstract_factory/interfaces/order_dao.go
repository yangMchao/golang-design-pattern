package interfaces

// OrderMainDAO 为订单主记录
// 负责订单主记录的持久化操作
// 实现此接口的具体类型需要实现SaveOrderMain方法
// 可以基于不同的存储技术实现（如RDB、XML、JSON等）
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情记录
// 负责订单详情记录的持久化操作
// 实现此接口的具体类型需要实现SaveOrderDetail方法
// 通常与OrderMainDAO配套使用，组成完整的产品族
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象模式工厂接口
// 定义创建订单相关DAO对象的抽象工厂
// 具体工厂实现此接口来创建特定存储技术的DAO对象族
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}
