package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type InputParser struct {

}

func (input *InputParser) ParsePoint() (x int,y int) {
	var i int
	var j int
	fmt.Print("Enter X: ")
	fmt.Scanf("%d", &i)
	fmt.Println("You entered : "+strconv.Itoa(i))
	fmt.Print("Enter Y: ")
	fmt.Scanf("%d", &j)
	fmt.Println("You entered : "+strconv.Itoa(j))
	return i,j
}

func (input *InputParser) ParseName() string{
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	return name

}