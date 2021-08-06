package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var Version = "development"

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("This program only works on linux!!")
		return
	}
	vFlag := flag.Bool("v", false, "print version")
	cFlag := flag.Bool("c", false, "only assemble")
	oFlag := flag.String("o", "a.out", "output path")
	iFlag := flag.String("i", "", "input path")
	flag.Parse()
	if *vFlag {
		fmt.Println("Version: ", Version)
		return
	}
	if *iFlag == "" {
		fmt.Printf("Usage: %s [-i inpath] {-o outpath}", os.Args[0])
		return
	}
	file := *iFlag
	preProcess(file)
	cstring := ""
	cstring += "#include<stdio.h>\n"
	cstring += "int main() {char array[65535];char*ptr=array;"

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := string(bytes)
	cstring += getC(s)
	cstring += "}"
	brainPath := "/tmp/brainfuck.c"
	err = ioutil.WriteFile(brainPath, []byte(cstring), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *cFlag {
		if *oFlag == "a.out" {
			tmp := strings.TrimSuffix(*iFlag, ".b")
			*oFlag = tmp + ".o"
		}
		_, err = exec.Command("gcc", "-c", "-o", *oFlag, "-O2", brainPath).Output()
	} else {
		_, err = exec.Command("gcc", "-o", *oFlag, "-O2", brainPath).Output()
	}
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	exec.Command("strip", *oFlag)
	err = os.Remove(brainPath)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}

func getC(s string) string {
	cstring := ""
	for _, char := range s {
		switch char {
		case '>':
			cstring += "++ptr;"
		case '<':
			cstring += "--ptr;"
		case '+':
			cstring += "++*ptr;"

		case '-':
			cstring += "--*ptr;"

		case '.':
			cstring += "putchar(*ptr);"

		case ',':
			cstring += "*ptr = getchar();"

		case '[':
			cstring += "while (*ptr) {"

		case ']':
			cstring += "}"

		}
	}
	return cstring
}

func preProcess(filePath string) *string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !strings.HasPrefix("@", scanner.Text()) {
			continue
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	return nil
}
