package abstractfactory

import (
	"golang-design-pattern/05_abstract_factory/factories"
	"golang-design-pattern/05_abstract_factory/interfaces"
	"testing"
)

// getMainAndDetail 客户端函数：使用抽象工厂创建并使用DAO对象
// 该函数只依赖于抽象工厂和抽象产品接口，不依赖于具体实现
func getMainAndDetail(factory interfaces.DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

// ExampleRDBDAOFactory 测试RDB工厂的使用示例
// 演示如何使用RDBDAOFactory创建RDB产品族的DAO对象
func ExampleRDBDAOFactory() {
	var factory interfaces.DAOFactory
	factory = &factories.RDBDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

// ExampleXMLDAOFactory 测试XML工厂的使用示例
// 演示如何使用XMLDAOFactory创建XML产品族的DAO对象
func ExampleXMLDAOFactory() {
	var factory interfaces.DAOFactory
	factory = &factories.XMLDAOFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}

// TestDynamicFactory 测试动态工厂选择
// 演示如何根据配置动态选择不同的工厂实现
func TestDynamicFactory(t *testing.T) {
	// 测试用例结构
	tests := []struct {
		name     string
		factory  interfaces.DAOFactory
		expected []string
	}{
		{
			name:    "RDB Factory",
			factory: &factories.RDBDAOFactory{},
			expected: []string{
				"rdb main save",
				"rdb detail save",
			},
		},
		{
			name:    "XML Factory",
			factory: &factories.XMLDAOFactory{},
			expected: []string{
				"xml main save",
				"xml detail save",
			},
		},
	}

	// 执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 这里可以添加更复杂的测试逻辑
			// 目前主要是验证工厂能够创建正确的对象
			mainDAO := tt.factory.CreateOrderMainDAO()
			detailDAO := tt.factory.CreateOrderDetailDAO()

			if mainDAO == nil || detailDAO == nil {
				t.Errorf("工厂创建的对象不能为nil")
			}
			mainDAO.SaveOrderMain()
			detailDAO.SaveOrderDetail()
		})
	}
}
