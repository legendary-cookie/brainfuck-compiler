package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("This program only works on linux!!")
		return
	}
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s FILE OUT\n",os.Args[0])
		return
	}

	file := os.Args[1] 
	cstring := ""
	cstring += "#include<stdio.h>\n"
	cstring += "int main() {char array[65535];char*ptr=array;"
	
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := string(bytes)
	for _, char := range s {
		switch(char) {
		case '>':
			cstring += "++ptr;"
			break
		case '<':
			cstring += "--ptr;"
			break
		case '+':
			cstring += "++*ptr;"
			break
		case '-':
			cstring += "--*ptr;"
			break
		case '.':
			cstring += "putchar(*ptr);"
			break
		case ',':
			cstring += "*ptr = getchar();"
			break
		case '[':
			cstring += "while (*ptr) {"
			break
		case ']':
			cstring += "}"
			break
		}
	}
	cstring += "}"
	err = ioutil.WriteFile("/tmp/brainfuck.c",[]byte(cstring),0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	_,err = exec.Command("gcc", "-O2","-o", os.Args[2], "/tmp/brainfuck.c").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	exec.Command("strip", os.Args[2])
	err = os.Remove("/tmp/brainfuck.c")
	if err != nil {
		fmt.Printf("%s",err)
	}
}
