package decorator

import (
	"fmt"
)

type display interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
	show(display display)
}
type defaultDisplay struct {
	display
}

func (this *defaultDisplay) show(display display) {
	for i := 0; i < display.getRows(); i++ {
		fmt.Println(display.getRowText(i))
	}
}

type StrDisplay struct {
	display
	str string
}

func (this *StrDisplay) getColumns() int {
	return len(this.str)
}
func (this *StrDisplay) getRows() int {
	return 1
}
func (this *StrDisplay) getRowText(row int) string {
	if row == 0 {
		return this.str
	}
	return ""
}

type Border struct {
	display
}

type SideBorder struct {
	char string
	display
}

func (this *SideBorder) getColumns() int {
	return 1 + this.display.getColumns() + 1
}
func (this *SideBorder) getRows() int {
	return this.display.getRows()
}
func (this *SideBorder) getRowText(row int) string {
	return this.char + this.display.getRowText(row) + this.char
}

type FullBorder struct {
	display
}

func (this *FullBorder) getColumns() int {
	return 1 + this.display.getColumns() + 1
}
func (this *FullBorder) getRows() int {
	return 1 + this.display.getRows() + 1
}
func (this *FullBorder) getRowText(row int) string {
	if row == 0 {
		return "+" + this.makeLine("-", this.display.getColumns()) + "+"
	} else if row == this.display.getRows()+1 {
		return "+" + this.makeLine("-", this.display.getColumns()) + "+"
	}
	return "|" + this.display.getRowText(row-1) + "|"
}
func (this *FullBorder) makeLine(ch string, count int) string {
	var str string
	for i := 0; i < count; i++ {
		str += ch
	}
	return str
}

func main() {
	b1 := StrDisplay{&defaultDisplay{}, "Hello Word"}
	b1.show(&b1)

	b2 := SideBorder{char: "#", display: &Border{&b1}}
	b2.show(&b2)

	b3 := FullBorder{&b2}
	b3.show(&b3)

	b4 := FullBorder{&b3}
	b4.show(&b4)
}
