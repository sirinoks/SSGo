package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-v", "--version":
			fmt.Println("0.0")
		case "-t", "--twitch":
			fmt.Println("dead chat")
		case "-nv", "--nodeVersion":
			dat, err := os.ReadFile("go.mod")
			check(err)
			contents := string(dat[:])
			splits := strings.Split(contents, "\n\n")
			fmt.Println(splits[len(splits)-1])
		default:
			fmt.Println("No arg dum-dum")
		}
	}

	//outputDir := "dist"

	files, err := ioutil.ReadDir("./")

	for _, f := range files {
		if !(f.IsDir() && f.Name() == "dist") {
			fmt.Println("Making new dist dir")

			err = os.Mkdir("dist", 0755)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
	/*
		files, err = ioutil.ReadDir("./input")

		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			fmt.Println(f.Name())
		}
	*/
}
