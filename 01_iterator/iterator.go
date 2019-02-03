package iterator

//抽象迭代子(Iterator)角色：此抽象角色定义出遍历元素所需的接口
type Iterator interface {
	hasNext() bool
	next() *Book
}

//聚集（Aggregate）角色：此抽象角色给出创建迭代子（Iterator）对象的接口
type Aggregate interface {
	iterator() Iterator
}

type Book struct {
	name string
}

func (book *Book) setName(name string) {
	book.name = name
}
func (book *Book) getName() string {
	return book.name
}

//具体聚集（Concrete Aggregate）角色：实现了创建迭代子（Iterator）对象的接口，返回一个合适的具体迭代子实例
type BookCollection struct {
	bookList []Book
	bookMax  int
}

func (bookCollection *BookCollection) addBook(book Book) {
	bookCollection.bookList = append(bookCollection.bookList, book)
	bookCollection.bookMax++
}
func (bookCollection *BookCollection) getBook(index int) Book {
	return bookCollection.bookList[index]
}
func (bookCollection *BookCollection) getLength() int {
	return bookCollection.bookMax
}
func (bookCollection *BookCollection) iterator() Iterator {
	var iterator Iterator
	bookIterator := BookCollectionIterator{}
	bookIterator.init(*bookCollection)
	iterator = &bookIterator
	return iterator
}

//具体迭代子（Concrete Iterator）角色：此角色实现了Iterator接口，并保持迭代过程中的游标位置
type BookCollectionIterator struct {
	bookCollection BookCollection
	index          int
}

func (bookCollectionIterator *BookCollectionIterator) init(bookCollection BookCollection) {
	bookCollectionIterator.bookCollection = bookCollection
	bookCollectionIterator.index = 0
}
func (bookCollectionIterator *BookCollectionIterator) hasNext() bool {
	if bookCollectionIterator.index < bookCollectionIterator.bookCollection.getLength() {
		return true
	}
	return false
}

func (bookCollectionIterator *BookCollectionIterator) next() *Book {
	book := bookCollectionIterator.bookCollection.getBook(bookCollectionIterator.index)
	bookCollectionIterator.index++
	return &book
}

func main() {
	book1 := Book{name: "Mysql从入门到删库"}
	book2 := Book{name: "Go从入门到放弃"}
	bookList := BookCollection{}
	bookList.addBook(book1)
	bookList.addBook(book2)

	iterator := BookCollectionIterator{}
	iterator.init(bookList)

	for iterator.hasNext() {
		book := iterator.next()
		println(book.getName())
	}

	//迭代器最好作为通过聚集（bookList）获取（c# List就是实现了可迭代的接口，获得了迭代器，就可以使用foreach遍历）
	bookIterator := bookList.iterator()
	for bookIterator.hasNext() {
		book := bookIterator.next()
		println(book.getName())
	}
}
