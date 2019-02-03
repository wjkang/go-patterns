package factorymethod

import (
	"fmt"
)

//抽象工厂（Creator）角色：任何在模式中创建对象的工厂类必须实现这个接口
type Company interface {
	BuildProduct(parameters string) Product
}

//抽象产品（Product）角色：工厂方法模式所创建的对象的超类型，也就是产品对象的共同父类或共同拥有的接口
type Product interface {
	DoUse()
}

//具体产品（Concrete Product）角色：这个角色实现了抽象产品角色所声明的接口
type ProductA struct {
	ProductType       string
	ProductParameters string
}

func (this *ProductA) DoUse() {
	fmt.Println("产品A的实现")
}

type ProductB struct {
	ProductType       string
	ProductParameters string
}

func (this *ProductB) DoUse() {
	fmt.Println("产品B的实现")
}

//具体工厂（Concrete Creator）角色：担任这个角色的是实现了抽象工厂接口的具体类
type CompanyA struct {
}

func (this *CompanyA) BuildProduct(parameters string) Product {
	if parameters == "A" {
		product := ProductA{ProductType: "ProductA", ProductParameters: parameters}
		return &product
	} else if parameters == "B" {
		product := ProductB{ProductType: "ProductB", ProductParameters: parameters}
		return &product
	} else {
		return nil
	}
}

func main() {
	var company Company
	company = &CompanyA{}
	productA := company.BuildProduct("A")
	productA.DoUse()

	productB := company.BuildProduct("B")
	productB.DoUse()
}
