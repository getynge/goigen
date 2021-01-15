package main

import (
	"fmt"
	"goigen/generator"
	"goigen/processor"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	if len(os.Args) != 4 {
		fmt.Println("Usage: goigen DIRECTORY STRUCT INTERFACE\nexample: goigen . Example IExample")
		os.Exit(0)
	}

	directory := os.Args[1]
	targetStruct := os.Args[2]
	targetInterface := os.Args[3]

	methods, pkg, err := processor.ProcessDirectory(directory, targetStruct)

	if err != nil {
		l.Fatal(err)
	}

	fileTemplate := generator.NewFileTemplate(methods, pkg, targetInterface)
	text, err := fileTemplate.Generate(directory)

	if err != nil {
		l.Fatal(err)
	}

	fmt.Println(text)
}
