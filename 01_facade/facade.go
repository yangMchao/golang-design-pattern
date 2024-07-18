package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{
		A: NewAModuleAPI(),
		B: NewBModuleAPI(),
	}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

// apiImpl facade implement
type apiImpl struct {
	A AModuleAPI
	B BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.A.TestA()
	bRet := a.B.TestB()
	return fmt.Sprintf("%s \t %s", aRet, bRet)
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
