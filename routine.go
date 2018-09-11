package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("cat.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(data)
	splitS := strings.Split(str, "\n")
	height := len(splitS)
	width := len(strings.Split(splitS[0], " "))
	matrix := make([][]string, height-1)
	for i:= 0; i < height-1; i++ {
		strSlice := strings.Split(splitS[i], " ")
		matrix[i] = make([]string, width)
		for j:= 0; j < width; j++ {
			fmt.Println(i, j)
			matrix[i][j] = strSlice[j]
			fmt.Println(matrix)
		}
	}	
	fmt.Println(matrix)
}
