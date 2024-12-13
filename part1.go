package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	//read inputs
	scanner := bufio.NewScanner(file)
	var a [][]byte

	// Read each line
	for scanner.Scan() {
		line := scanner.Bytes()
		copiedLine := append([]byte{}, line...) // Copy the slice
		a = append(a, copiedLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)

	}

	for i:= 0; i<len(a); i++ {
		if(len(a[i])!=len(a[0])){
			fmt.Println("Error bad input", err)
			os.Exit(1)
		}
	}

	total := 0
	for i:= 0; i<len(a); i++ {
		for j := 0; j<len(a[0]);j++{
			down := i<len(a)-3
			right := j<len(a[0])-3

			//diags
			if (right && down){
				//diag1

				if a[i][j] == 'X' &&
				   a[i+1][j+1] == 'M' &&
				   a[i+2][j+2] == 'A' &&
				   a[i+3][j+3] == 'S' {
				   	//fmt.Printf("Found diag XMAS: Start (%d, %d), End (%d, %d)\n", i, j, i+3, j+3)
				    total++
				}
				if a[i][j] == 'S' &&
				   a[i+1][j+1] == 'A' &&
				   a[i+2][j+2] == 'M' &&
				   a[i+3][j+3] == 'X' {
				   	//fmt.Printf("Found diag XMAS: Start (%d, %d), End (%d, %d)\n", i+3, j+3, i, j)
				    total++
				}

				//diag 2
				if a[i][j+3] == 'X' &&
				   a[i+1][j+2] == 'M' &&
				   a[i+2][j+1] == 'A' &&
				   a[i+3][j] == 'S' {
				   	//fmt.Printf("Found diag XMAS: Start (%d, %d), End (%d, %d)\n", i, j+3, i+3, j)
				    total++
				}
				if a[i][j+3] == 'S' &&
				   a[i+1][j+2] == 'A' &&
				   a[i+2][j+1] == 'M' &&
				   a[i+3][j+0] == 'X' {
				   	//fmt.Printf("Found diag XMAS: Start (%d, %d), End (%d, %d)\n", i+3, j, i, j+3)
				    total++
				}
			}
			//line
			if (right){
				if a[i][j] == 'X' &&
				   a[i][j+1] == 'M' &&
				   a[i][j+2] == 'A' &&
				   a[i][j+3] == 'S' {
				   	//fmt.Printf("Found horiz XMAS: Start (%d, %d), End (%d, %d)\n", i, j, i, j+3)
				    total++
				}
				if a[i][j] == 'S' &&
				   a[i][j+1] == 'A' &&
				   a[i][j+2] == 'M' &&
				   a[i][j+3] == 'X' {
				   	//fmt.Printf("Found horiz XMAS: Start (%d, %d), End (%d, %d)\n", i, j+3, i, j)
				    total++
				}
			}
			//colum
			if (down){
				if a[i][j] == 'X' &&
				   a[i+1][j] == 'M' &&
				   a[i+2][j] == 'A' &&
				   a[i+3][j] == 'S' {
				   	//fmt.Printf("Found ver XMAS: Start (%d, %d), End (%d, %d)\n", i, j, i+3, j)
				    total++
				}
				if a[i][j] == 'S' &&
				   a[i+1][j] == 'A' &&
				   a[i+2][j] == 'M' &&
				   a[i+3][j] == 'X' {
				   	//fmt.Printf("Found ver XMAS: Start (%d, %d), End (%d, %d)\n", i+3, j, i, j)
				    total++
				}
			}
			   
		}
	}

	fmt.Println("total", total)
}
