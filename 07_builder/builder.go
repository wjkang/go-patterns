package builder

import "fmt"

type Builder interface {
	makeTitle(title string)
	makeString(str string)
	makeItems(items []string)
	close()
	getResult() string
}

type Director struct {
	Builder
}

func (this *Director) construct() {
	this.Builder.makeTitle("Greeting")
	this.Builder.makeString("从早上到下午")
	this.Builder.makeItems([]string{"早上好", "下午好"})
	this.Builder.makeString("晚上")
	this.Builder.makeItems([]string{"晚上好", "晚安", "再见"})
	this.Builder.close()
}

type TextBuilder struct {
	result string
}

func (this *TextBuilder) makeTitle(title string) {
	this.result = this.result + "==============================\n"
	this.result = this.result + "[ " + title + " ]\n"
	this.result = this.result + "\n"
}

func (this *TextBuilder) makeString(title string) {
	this.result = this.result + "☆" + title + "\n"
	this.result = this.result + "\n"
}
func (this *TextBuilder) makeItems(items []string) {
	for _, s := range items {
		this.result = this.result + s + "\n"
	}
	this.result = this.result + "\n"
}
func (this *TextBuilder) close() {
	this.result = this.result + "\n"
}

func (this *TextBuilder) getResult() string {
	return this.result
}

type HTMLBuilder struct {
	result string
}

func (this *HTMLBuilder) makeTitle(title string) {
	this.result = this.result + "==============================\n"
	this.result = this.result + "<h1>" + title + "</h1>\n"
	this.result = this.result + "\n"
}

func (this *HTMLBuilder) makeString(title string) {
	this.result = this.result + "<p>" + title + "</p>\n"
	this.result = this.result + "\n"
}
func (this *HTMLBuilder) makeItems(items []string) {
	this.result = this.result + "<ul>\n"
	for _, s := range items {
		this.result = this.result + "<li>" + s + "</li>\n"
	}
	this.result = this.result + "</ul>\n"
}
func (this *HTMLBuilder) close() {
	this.result = this.result + "\n"
}

func (this *HTMLBuilder) getResult() string {
	return this.result
}

func main() {
	var textBuilder Builder
	var htmlBuilder Builder

	textBuilder = &TextBuilder{}

	htmlBuilder = &HTMLBuilder{}

	director1 := Director{Builder(textBuilder)}
	director1.construct()
	fmt.Print(textBuilder.getResult())

	director2 := Director{Builder(htmlBuilder)}
	director2.construct()
	fmt.Print(htmlBuilder.getResult())
}
