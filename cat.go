package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func cat(s string) {
	file, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("error: Not Such file")
	}
	fmt.Println(string(file))
}

func main() {
	var op_n, op_b bool
	flag.BoolVar(&op_n, "n", false, "Number the output lines, starting at 1.")
	flag.BoolVar(&op_b, "b", false, "Number the non-blank output lines, starting at 1.")
	flag.Parse()
	command_args := flag.Args()
	if op_b {
		for _, v := range command_args {
			file, err := ioutil.ReadFile(v)
			if err != nil {
				fmt.Println("error: Not Such file")
			}
			string_file := string(file)
			string_file = strings.TrimRight(string_file, "\n")
			sp_string := strings.Split(string_file, "\n")
			sp_length := len(sp_string)
			sp_length_string := len(fmt.Sprint(sp_length))
			number := 1
			for _, v := range sp_string {
				if g, _ := regexp.MatchString(`\A(\b+|\s+|)\z`, v); g || v == "" {
					fmt.Print("     " + strings.Repeat(" ", sp_length_string) + "  " + string(v) + "\n")
				} else {
					fmt.Print("     " + fmt.Sprint(number) + strings.Repeat(" ", sp_length_string-len(fmt.Sprint(number))) + "  " + string(v) + "\n")
					number += 1
				}
			}
		}
	} else if op_n {
		for _, v := range command_args {
			file, err := ioutil.ReadFile(v)
			if err != nil {
				fmt.Println("error: Not Such file")
			}
			string_file := string(file)
			string_file = strings.TrimRight(string_file, "\n")
			sp_string := strings.Split(string_file, "\n")
			for n, v := range sp_string {
				number := n + 1
				fmt.Print("     " + fmt.Sprint(number) + "  " + string(v) + "\n")
			}
		}
	} else {
		for _, v := range command_args {
			cat(v)
		}
	}
	// How to read file.
	// file, _ := ioutil.ReadFile("/home/readme.txt")
	// fmt.Println(string(file))
	//	u := string(file)
	//	u += "#created by go\n"
	//	m := []byte(u)
	//	ioutil.WriteFile("/home/readme.txt", m, 0755)
}
