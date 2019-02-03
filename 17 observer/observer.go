package main

import (
	"math/rand"
	"fmt"
)

type NumberGenerator struct {
    observer []Observer
}
func (this *NumberGenerator) AddObserver(observer Observer){
    this.observer=append(this.observer,observer)
}
func (self *NumberGenerator) notifyObservers() {

	for _, observer := range self.observer {
		observer.update()
	}

}

type randomNumberGenerator struct {
	*NumberGenerator
}

func (self *randomNumberGenerator) getNumber() int {
	return rand.Intn(10)
}
func (self *randomNumberGenerator) Execute(){
	self.notifyObservers()
}


type Observer interface {
	update()
}
type DigitObserver struct {
	generator randomNumberGenerator
}

func (self *DigitObserver) update() {
	fmt.Println(self.generator.getNumber())
}

type GraphObserver struct {
	generator randomNumberGenerator
}

func (self *GraphObserver) update() {
	i:=self.generator.getNumber()
	for j:=0;j<i;j++{
		fmt.Print("*")
	}
}

func main(){
	random := randomNumberGenerator{&NumberGenerator{}}

	o1 := &DigitObserver{random}
	o2 := &GraphObserver{random}

	random.AddObserver(o1)
	random.AddObserver(o2)

	random.Execute()

}