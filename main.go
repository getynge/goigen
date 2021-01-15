package main

import (
	"fmt"
	"github.com/getynge/goigen/generator"
	"github.com/getynge/goigen/processor"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
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

	fullPath := path.Join(directory, fileTemplate.FileName+".go")
	err = ioutil.WriteFile(fullPath, []byte(text), 0755)

	if err != nil {
		l.Fatal(err)
	}

	cmd := exec.Command("go", "generate", fullPath)

	_, err = cmd.Output()

	if err != nil {
		l.Printf("Failed to generate mocks due to error: %s\nPlease generate your mocks manually", err.Error())
	}
}
