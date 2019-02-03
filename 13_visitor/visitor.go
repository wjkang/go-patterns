package visitor

import (
	"fmt"
	"strconv"
)

type Visitor interface {
	visitFile(file *File)
	visitDir(directory *Directory)
}

type Element interface {
	accept(visitor Visitor)
}

type Entry interface {
	Element
	getName() string
	getSize() int
	toString() string
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
func (this *File) accept(visitor Visitor) {
	visitor.visitFile(this)
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
func (this *Directory) accept(visitor Visitor) {
	visitor.visitDir(this)
}

type ListVisitor struct {
	currentDir string
}

func (this *ListVisitor) visitFile(file *File) {
	fmt.Println(this.currentDir + "/" + file.toString())
}
func (this *ListVisitor) visitDir(dir *Directory) {
	saveDir := this.currentDir
	fmt.Println(this.currentDir + "/" + dir.toString())
	this.currentDir += "/" + dir.getName()
	if dir.directory != nil {
		for _, entry := range dir.directory {

			entry.accept(this)
		}
	}
	this.currentDir = saveDir
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

	visitor := &ListVisitor{}

	rootDir.accept(visitor)

}
