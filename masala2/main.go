package main

import (
	"log"
	"os"
)

func fileContent() {
	file, err := os.ReadFile("f1.txt")
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.ReadFile("f2.txt")
	if err != nil {
		log.Fatal(err)
	}

	f3, err := os.OpenFile("f3.txt", os.O_RDWR|os.O_CREATE| os.O_TRUNC| os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	f3.Write(file)
	f3.Write(f2)

	var i string
	i = string(file)+" "+string(f2)
	log.Println(i)
	f3.Close()
	


}

func main() {
	fileContent()
}
