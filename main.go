package main

import (
	"github.com/name1e5s/brainfuck/Brainfucker"
	"flag"
	"log"
	"io/ioutil"
	"os"
)

var (
	filename	string
)

func init(){
	flag.StringVar(&filename,"File name","","Brainfuck source.")
}


func main() {
	flag.Parse()
	var input string
	var byto []byte
	var err		error
	if filename == "" {
		log.Println("Now read source file from stdin.")
		byto,err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Panic("Read stdin error.")
		}
	} else {
		byto,err = ioutil.ReadFile(filename)
		if err != nil {
			log.Panic("Read file error.")
		}
	}
	input = string(byto)
	fucker := Brainfucker.NewFucker(input)
	fucker.Check()
	fucker.Run()
}
