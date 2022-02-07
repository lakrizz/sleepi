package main

import (
	"io/ioutil"

	"github.com/k0kubun/pp"
	"krizz.org/sleepi/pkg/library"
)

func main() {
	lib, err := library.GetLibrary()
	if err != nil {
		panic(err)
	}
	pp.Println(lib.Files)

	test_file := "/Users/kristofkipp/Documents/bbank.mp3"
	dat, err := ioutil.ReadFile(test_file)
	if err != nil {
		panic(err)
	}

	err = lib.AddFile(dat, "blood_bank.mp3")
	if err != nil {
		panic(err)
	}

	pp.Println(lib.Files)
}
