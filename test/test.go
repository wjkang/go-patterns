package main

import "coding.net/gopattern/singleton"

func main(){
	singleton:=singleton.GetInstance();
	singleton.Display();
}