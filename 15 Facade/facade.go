package main

import "fmt"

var db=map[string]string{
	"a@a.com":"a",
	"b@b.com":"b",
}

type database struct {

}

func (this *database) getNameByEmail(email string)  string{
	return db[email]
}


type mdWriter struct {

}
func (this *mdWriter) title(title string) string{
	return "# Welcome to " + title + "'s page!"
}

type PageMaker struct {
}

func (self *PageMaker) MakeWelcomePage(mail string) string {
	database := database{}
	writer := mdWriter{}

	name := database.getNameByEmail(mail)
	page := writer.title(name)

	return page
}
func main(){
	maker := PageMaker{}
	result := maker.MakeWelcomePage("a@a.com")
	fmt.Println(result)
}