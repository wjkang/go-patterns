package main

//抽象工厂
type AbstractCompany interface {
	buildComputer(parameters string) Computer
	buildTelephone(parameters string) Telephone
}

//抽象产品
type Computer interface {
	doUse()
}
//具体产品
type NotebookComputer struct {

}
type PersonalComputer struct {

}
func (computer NotebookComputer) doUse()  {

	println("这是笔记本电脑的功能")

}
func (computer PersonalComputer) doUse()  {

	println("这是个人计算机的功能")

}

//抽象产品
type Telephone interface {
	doUse()
}

//具体产品
type Mobile struct {

}
type DesktopPhone struct {

}
func (telephone Mobile) doUse() {

	println("这是手机的功能")

}
func (telephone DesktopPhone) doUse() {

	println("这是座机电话的功能")

}



type CompanyXX struct {

}
func (this *CompanyXX) buildComputer(parameters string) Computer{
	if  parameters == "个人电脑" {

		computer  := PersonalComputer{}

		return  computer

	}  else if parameters == "笔记本电脑" {

		computer  := NotebookComputer{}

		return  computer

	}  else {

		return  nil

	}
}

func (this *CompanyXX) buildTelephone(parameters string) Telephone{
	if  parameters == "座机电话" {

		telephone  := DesktopPhone{}

		return  telephone

	}  else if parameters == "手机" {

		telephone  := Mobile{}

		return  telephone

	}  else {

		return  nil

	}
}
func main(){
	//根据传入的参数得到Computer产品

	company  := CompanyXX{}

	computer1  := company.buildComputer("个人电脑")

	computer1.doUse()



	computer2  := company.buildComputer("笔记本电脑")

	computer2.doUse()



	//根据传入的参数得到Telephone产品

	telephone1  := company.buildTelephone("座机电话")

	telephone1.doUse()



	telephone2  := company.buildTelephone("手机")

	telephone2.doUse()
}