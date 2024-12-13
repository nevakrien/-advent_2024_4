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
	for i:= 0; i<len(a)-2; i++ {
		for j := 0; j<len(a[0])-2;j++{
			// fmt.Println("yay")
			
			if a[i+1][j+1] != 'A'{
				continue
			}

			// fmt.Println("A found")


			var diag1,diag2 bool
			if(a[i][j] == 'M' && a[i+2][j+2] == 'S'){
				diag1=true
			} else if(a[i][j] == 'S' && a[i+2][j+2] == 'M'){
				diag1=true
			}

			if(a[i+2][j] == 'M' && a[i][j+2] == 'S'){
				diag2=true
			} else if(a[i+2][j] == 'S' && a[i][j+2] == 'M'){
				diag2=true
			}

			if(diag1 && diag2){
				total++
			}

			// fmt.Println("diag1",diag1,"diag2",diag2)

		}
	}
	fmt.Println("total", total)
}
