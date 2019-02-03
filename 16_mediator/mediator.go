package mediator

import (
	"fmt"
)

//调停者（Mediator）模式标准定义：用一个中介对象来封装一系列的对象交互。中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互

//调停者
type Mediator interface {
	createColleagues()
	colleagueChanged()
}

//组员
type Colleague interface {
	setMediator(mediator Mediator)
	setColleagueEnabled(enabled bool)
}

type ColleagueButton struct {
	enabled  bool
	mediator Mediator
}

func (this *ColleagueButton) setMediator(mediator Mediator) {
	this.mediator = mediator
}
func (this *ColleagueButton) setColleagueEnabled(enabled bool) {
	this.enabled = enabled
}

type ColleagueRadioButton struct {
	checked  bool
	mediator Mediator
}

func (this *ColleagueRadioButton) setMediator(mediator Mediator) {
	this.mediator = mediator
}
func (this *ColleagueRadioButton) check(checked bool) {
	this.checked = checked
}
func (this *ColleagueRadioButton) statusChange() {
	this.mediator.colleagueChanged()
}

//输入框
type ColleagueTextField struct {
	value    string
	enabled  bool
	mediator Mediator
}

func (this *ColleagueTextField) setMediator(mediator Mediator) {
	this.mediator = mediator
}
func (this *ColleagueTextField) setColleagueEnabled(enabled bool) {
	this.enabled = enabled
}
func (this *ColleagueTextField) input(value string) {
	this.value = value
}
func (this *ColleagueTextField) textChange() {
	this.mediator.colleagueChanged()
}

type loginForm struct {
	GuestRadio ColleagueRadioButton
	LoginRadio ColleagueRadioButton
	UserText   ColleagueTextField
	PassText   ColleagueTextField
	ButtonOk   ColleagueButton
}

func (self *loginForm) createColleagues() {
	self.GuestRadio = ColleagueRadioButton{}
	self.LoginRadio = ColleagueRadioButton{}
	self.UserText = ColleagueTextField{}
	self.PassText = ColleagueTextField{}
	self.ButtonOk = ColleagueButton{}

	self.GuestRadio.setMediator(self)
	self.LoginRadio.setMediator(self)
	self.UserText.setMediator(self)
	self.PassText.setMediator(self)
	self.ButtonOk.setMediator(self)
}

func (self *loginForm) colleagueChanged() {
	if self.GuestRadio.checked {

		self.UserText.setColleagueEnabled(false)
		self.PassText.setColleagueEnabled(false)
		self.ButtonOk.setColleagueEnabled(false)
	} else {

		self.UserText.setColleagueEnabled(true)
		self.userOrPassChanged()
	}

	fmt.Println("Guest Mode", self.GuestRadio.checked)
	fmt.Println("Login Mode", self.LoginRadio.checked)
	fmt.Println("UserText", self.UserText.enabled)
	fmt.Println("PassText", self.PassText.enabled)
	fmt.Println("ButtonOk", self.ButtonOk.enabled)
}
func (self *loginForm) userOrPassChanged() {
	if len(self.UserText.value) > 0 {
		self.PassText.setColleagueEnabled(true)
		if len(self.PassText.value) > 0 {
			self.ButtonOk.setColleagueEnabled(true)
		} else {
			self.ButtonOk.setColleagueEnabled(false)
		}
	} else {

		self.PassText.setColleagueEnabled(false)
		self.ButtonOk.setColleagueEnabled(false)

	}
}

func main() {
	form := loginForm{}
	form.createColleagues()

	form.GuestRadio.check(true)
	form.GuestRadio.statusChange()

	fmt.Println("-----------------------------------")
	form.GuestRadio.check(false)
	form.LoginRadio.check(true)
	form.LoginRadio.statusChange()
	fmt.Println("-----------------------------------")
	form.UserText.input("122")
	form.UserText.textChange()
	fmt.Println("-----------------------------------")
	form.PassText.input("122")
	form.PassText.textChange()
}
