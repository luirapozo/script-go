package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	//create a file that tracks the times the .exe has bee run
	script.Exec("bash -c 'touch number'").Stdout()
	script.Exec("bash -c 'echo 1 >> number'").Stdout()

	data, err := script.File("number").String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data)
	data = strings.Split(data, "\n")[0]

	num, err := strconv.Atoi(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(num)
	//all the Hello World!'s that have been written are recorded
	for i := 0; i < num; i++ {
		script.Exec("bash -c 'echo Hello World! >> hola'").Stdout()
	}
	num++
}
