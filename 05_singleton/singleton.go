package singleton

import "fmt"

type singleton struct {

}
func (this *singleton) Display(){
	fmt.Println("单例模式");
}

var instance *singleton;

func GetInstance() *singleton{
	if instance==nil{
		instance=&singleton{};
	}
	return instance;
}
