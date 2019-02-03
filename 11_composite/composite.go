package composite

import (
	"fmt"
	"strconv"
)

type Entry interface {
	getName() string
	getSize() int
	toString() string
	printList(prefix string)
	add(entry Entry) Entry
}

type File struct {
	name string
	size int
}

func (this *File) getName() string {
	return this.name
}
func (this *File) getSize() int {
	return this.size
}
func (this *File) add(entry Entry) Entry {
	return nil
}
func (this *File) toString() string {
	return this.getName() + "(" + strconv.Itoa(this.getSize()) + ")"
}
func (this *File) printList(prefix string) {
	fmt.Println(prefix + "/" + this.toString())
}

type Directory struct {
	name      string
	directory []Entry
}

func (this *Directory) getName() string {
	return this.name
}

func (this *Directory) getSize() int {
	size := 0
	if this.directory != nil {
		for _, entry := range this.directory {
			size += entry.getSize()
		}
	}
	return size
}
func (this *Directory) add(entry Entry) Entry {
	this.directory = append(this.directory, entry)
	return this
}
func (this *Directory) toString() string {
	return this.getName() + "(" + strconv.Itoa(this.getSize()) + ")"
}
func (this *Directory) printList(prefix string) {
	fmt.Println(prefix + "/" + this.toString())
	if this.directory != nil {
		for _, entry := range this.directory {
			entry.printList(prefix + "/" + this.getName())
		}
	}
}

func main() {
	rootDir := Directory{name: "root"}
	binDir := Directory{name: "bin"}
	tmpDir := Directory{name: "tmp"}
	usrDir := Directory{name: "usr"}
	rootDir.add(&binDir)
	rootDir.add(&tmpDir)
	rootDir.add(&usrDir)

	binDir.add(&File{name: "vi", size: 10000})
	binDir.add(&File{name: "latex", size: 20000})
	rootDir.printList("")
}
