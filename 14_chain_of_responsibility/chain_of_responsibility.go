package chain_of_responsibility

import (
	"fmt"
	"strconv"
)

type Trouble struct {
	number int
}

func (this *Trouble) getNumber() int {
	return this.number
}

type Support interface {
	resolve(trouble Trouble) bool
	handle(support Support, trouble Trouble) string
}
type defaultSupport struct {
	Support
	name string
	next Support
}

func (this *defaultSupport) setNext(support Support) {
	this.next = support
}
func (self *defaultSupport) done(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " is resolved by " + self.name
}
func (self *defaultSupport) fail(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " cannot be resolved"
}
func (this *defaultSupport) handle(support Support, trouble Trouble) string {
	if support.resolve(trouble) {
		return this.done(trouble)
	} else if this.next != nil {
		return this.next.handle(this.next, trouble)
	} else {
		return this.fail(trouble)
	}
}

type noSupport struct {
	*defaultSupport
}

func (self *noSupport) resolve(trouble Trouble) bool {
	return false
}

type limitSupport struct {
	*defaultSupport
	limit int
}

func (self *limitSupport) resolve(trouble Trouble) bool {
	if trouble.getNumber() < self.limit {
		return true
	} else {
		return false
	}
}

func main() {
	a := noSupport{&defaultSupport{name: "A"}}
	b := limitSupport{&defaultSupport{name: "B"}, 5}
	c := limitSupport{&defaultSupport{name: "C"}, 3}
	a.setNext(&b)
	b.setNext(&c)

	result := a.handle(&a, Trouble{3})

	fmt.Println(result)
}
