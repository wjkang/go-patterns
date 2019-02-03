package main

import "fmt"

type AbstractDisplay interface {
	open();
	print();
	close();
	display();
}

type StringDisplay struct {
	str string;
}

func (this *StringDisplay) open(){
	fmt.Print("<<");
}
func (this *StringDisplay) print(){
	fmt.Print(this.str);
}
func (this *StringDisplay) close(){
	fmt.Print(">>")
}
func (this *StringDisplay) display(){
	this.open();
	for i:=0;i<=5;i++{
		this.print();
	}
	this.close();
}

func main(){
	display:=StringDisplay{str:"123"};
	display.display();
}
//定义了一个或多个抽象操作，以便让子类实现。这些抽象操作叫做基本操作，它们是一个顶级逻辑的组成步骤。定义并实现了一个模板方法。这个模板方法一般是一个具体方法，它给出了一个顶级逻辑的骨架，而逻辑的组成步骤在相应的抽象操作中，推迟到子类实现。顶级逻辑也有可能调用一些具体方法
