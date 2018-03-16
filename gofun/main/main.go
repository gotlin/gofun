package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
)

func main() {

	io.Copy(os.Stdout,os.Stdin)
}

func ReadInput(){

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Println(input)

}
