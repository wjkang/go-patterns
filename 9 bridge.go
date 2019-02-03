package main

import "fmt"

//类的功能层次结构：父类具有基本功能，在子类中增加新的功能
//类的实现层次结构：父类通过声明抽象方法来定义接口，子类通过实现具体方法来实现接口

//类的实现层次结构顶层
type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}
type StringDisplayImpl struct {
	str string
	width int
}
func (this *StringDisplayImpl) init(str string){
	this.str=str
	this.width=len(str)
}
func (this *StringDisplayImpl) printLine(){
	fmt.Print("+")
	for i:=0;i<this.width;i++{
		fmt.Print("-")
	}
	fmt.Print("+\n")
}

func (this *StringDisplayImpl) rawOpen(){
   this.printLine()
}
func (this *StringDisplayImpl) rawPrint(){
	fmt.Println("|"+this.str+"|")
}
func (this *StringDisplayImpl) rawClose(){
	this.printLine()
}
//类的功能层结构顶层
type Display struct {
	 impl DisplayImpl
}
func (this *Display) open(){
	this.impl.rawOpen()
}
func (this *Display) print(){
	this.impl.rawPrint()
}
func (this *Display) close(){
	this.impl.rawClose()
}
func (this *Display) display(){
	this.open()
	this.print()
	this.close()
}

type CountDisplay struct {
	Display
}
func (this *CountDisplay) multiDisplay(times int){
	this.open()
	for i:=0;i<times;i++{
		this.print()
	}
	this.close()
}

func main(){
	impl1:=StringDisplayImpl{}
	impl1.init("Hello Go!")
	display1:=Display{&impl1}
	display1.display()

	mountDisplay1:=CountDisplay{display1}
	mountDisplay1.multiDisplay(5)
}

//impl就是功能层次结构和实现层次结构之间的桥梁

