package main

type Strategy interface {
    doWork(project Project)
}

type Project struct {
	projectName string
	strategy Strategy
}

func (project *Project)  setStrategy(strategy Strategy) {

	project.strategy  = strategy

}
func (project *Project) getStrategy()  Strategy {

	return  project.strategy

}
func (project *Project) doStrategyWork()  {

	project.strategy.doWork(*project)

}

type BankStrategy struct {

}
func (bankStrategy *BankStrategy)  doWork(project Project) {

	println(project.projectName  + "采用银行策略的工作。")

}



type GovernmentStrategy struct {

}
func (governmentStrategy  *GovernmentStrategy) doWork(project Project) {

	println(project.projectName  + "采用政府策略的工作")

}

type TelecomStrategy struct {

}
func (telecomStrategy *TelecomStrategy)  doWork(project Project) {

	println(project.projectName  + "采用电信策略的工作。")

}

func main(){
	println("———要求projectA采用银行策略———")

	projectA  := Project{}

	projectA.projectName="projectA"

	strategy1  := &BankStrategy{}

	projectA.setStrategy(strategy1)

	projectA.doStrategyWork()



	println("———要求projectB采用政府策略———")

	projectB  := Project{}

	projectB.projectName="projectB"

	strategy2  := &GovernmentStrategy{}

	projectB.setStrategy(strategy2)

	projectB.doStrategyWork()



	println("———要求projectC采用电信策略———")

	projectC  := Project{}

	projectC.projectName="projectC"

	strategy3  := &TelecomStrategy{}

	projectC.setStrategy(strategy3)

	projectC.doStrategyWork()
}