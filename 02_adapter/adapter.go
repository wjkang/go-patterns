package adapter

import "fmt"

type IModel interface {
	add() string
	update() string
}

type XOrmModel struct {
	xOperation string
}

type AdapterModel struct {
	XOrmModel
}

func (this *XOrmModel) xAdd() string {
	return this.xOperation + ":insert into t_user(name) values('test')"
}
func (this *XOrmModel) xUpdate() string {
	return this.xOperation + ":update t_user set name='test'"
}

func (this *AdapterModel) add() string {
	return this.xAdd()
}

func (this *AdapterModel) update() string {
	return this.xUpdate()
}

func main() {
	var model IModel

	model = &AdapterModel{XOrmModel{xOperation: "XOrmModel Operation"}}

	fmt.Println(model.add())

	fmt.Println(model.update())
}
