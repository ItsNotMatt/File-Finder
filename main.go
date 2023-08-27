package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var result []string

func IndexDir(path string, search string) {
    dir, err := os.ReadDir(path)
    if err != nil {
        panic("Cant open dir")
    }
    for _, file := range dir {
        if file.IsDir() {
            newPth := path + `\` + file.Name() 
            IndexDir(newPth, search)
        }
        fileName := strings.ToLower(file.Name())
        if strings.Contains(fileName, search) {
            result = append(result, fileName)
        }
    }
}

func main() {
    path, err := os.Getwd()
    if err != nil {
        panic("Cant get working directory")
    }
    path = `D:\Games\Steam`
    if len(os.Args) != 2 {
        panic("Command accepts one argument")
    }
    input := os.Args[1]
    search := strings.ToLower(input)

    start := time.Now()
    IndexDir(path, search)
    end := time.Now()
    elapsed := end.Sub(start)
    //3.1s with this
    //2m with explorer

    if len(result) != 0 {
        fmt.Printf("\n\n----Found Files---\n")
        for _, file := range result {
            fmt.Println("File: ", file)
        }
    } else {
        fmt.Println("No results found")
    }
    fmt.Println("Time:", elapsed)
}
