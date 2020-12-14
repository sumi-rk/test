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
	var body [][][2]string

	fmt.Println(os.Args[1])

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	content, _, err := docconv.ConvertDocx(file)
	index := strings.Index(content, "日報") + 6
	fmt.Println(index)
	content = content[index:]
	content = strings.TrimSpace(content)

	r := regexp.MustCompile("\r\n|\n\r|\n|\r")
	line := r.ReplaceAllString(content, "\n")
	fmt.Println(strings.Split(line, "\n"))

	var b string
	var date string
	var title_flag bool
	var title_index int
	var body_part [][2]string
	var body_part2 [2]string

	title_flag = false
	title_index = 0

	slice := strings.Split(line, "\n")
	// for _, i := range slice {
	// 	fmt.Printf("[%s]", i)
	// }
	for _, s := range slice {
		if strings.HasPrefix(s, "作成日") {
			date = s[12:19]
			date = strings.Replace(date, "月", "/", 1)
		} else if strings.HasPrefix(s, "作成者") {
		} else if strings.HasPrefix(s, "・") {
			b = s
		} else if strings.HasPrefix(s, "＊") {
			title_flag = false
			title_index = 0
			rep := regexp.MustCompile(`-\[.*\]`)
			s = rep.ReplaceAllString(s, "")
			for _, t := range bTitle {
				// fmt.Println(t[0])
				if b == t[0] && s == t[1] {
					title_flag = true
					// fmt.Println(title_flag)
					break
				}
				title_index += 1
			}
			if !title_flag {
				// fmt.Println("!!!!!!!!!!!!!!!")
				title := [...]string{b, s}
				bTitle = append(bTitle, title)
			}
		} else if s == "" {
			fmt.Println("empty")
		} else {
			// fmt.Println(title_flag)
			body_part2[0] = date
			body_part2[1] = s
			if title_flag {
				// fmt.Println(body[title_index])
				body_part = append(body[title_index], body_part2)
				body[title_index] = body_part
			} else {
				body_part = nil
				body_part = append(body_part, body_part2)
				body = append(body, body_part)
			}
		}
		// fmt.Println(date)
	}
	fmt.Println(bTitle)
	fmt.Println(body)
	// for _, i := range body {
	// 	fmt.Printf("[%s]", i)
	// }
}
