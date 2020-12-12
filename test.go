package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"code.sajari.com/docconv"
)

func main() {
	var bTitle [][2]string
	// var sTitle [][]string
	var body []string

	// targetDir := "."
	// if len(os.Args) == 2 {
	// 	targetDir = os.Args[1]
	// }
	// pattern := targetDir + "/*.docx"
	// files, err := filepath.Glob(pattern)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, file := range files {
	// 	fmt.Println(file)
	// }

	// bTitle = append(bTitle, "test")
	// bTitle = append(bTitle, "test1")
	// fmt.Println(bTitle)

	// sTitle = append(sTitle, bTitle)
	// fmt.Println(sTitle)
	// fmt.Println(sTitle[0][1])
	// fmt.Println("b\nc\r\nd")

	file, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	content, _, err := docconv.ConvertDocx(file)
	content = strings.TrimSpace(content)

	r := regexp.MustCompile("\r\n|\n\r|\n|\r")
	line := r.ReplaceAllString(content, "\n")
	fmt.Println(strings.Split(line, "\n"))

	var b string

	slice := strings.Split(line, "\n")
	for _, s := range slice {
		if strings.HasPrefix(s, "・") == true {
			b = s
		} else if strings.HasPrefix(s, "＊") == true {
			rep := regexp.MustCompile(`-\[.*\]`)
			s = rep.ReplaceAllString(s, "")
			title := [...]string{b, s}
			bTitle = append(bTitle, title)
		} else if s == "" {
			fmt.Println("empty")
		} else {
			body = append(body, s)
		}
	}
	fmt.Println(bTitle)
	fmt.Println(body)
	for _, i := range body {
		fmt.Printf("[%s]", i)
	}
}
