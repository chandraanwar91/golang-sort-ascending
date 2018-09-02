package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var fileName string
	var extention string
	fmt.Println("Enter filename : ")
	fmt.Scan(&fileName)
	fmt.Println("Enter Extention : ")
	fmt.Scan(&extention)
	lines, err := readLines(fileName+"."+extention)

	if err != nil {
		fmt.Println("Error Reading File "+ err.Error())
	}
	//set variable for check
	countLine := len(lines)
	swapped := true
	for swapped {
        swapped = false
        for i := 1; i < countLine; i++ {
			var temp int
            if lines[i-1] > lines[i] {
				temp = lines[i-1]
                lines[i-1], lines[i] = lines[i], temp
                swapped = true
            }
		}
	}
	
	writeLines(lines,fileName+"_sort."+extention)
	
}

func readLines(path string) ([]int, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []int
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
	  textNumber,_ := strconv.Atoi(scanner.Text())
    lines = append(lines,textNumber )
  }
  return lines, scanner.Err()
}

func writeLines(lines []int, path string) error {
	file, err := os.Create(path)
	if err != nil {
	  return err
	}
	defer file.Close()
  
	w := bufio.NewWriter(file)
	for _, line := range lines {
	  fmt.Fprintln(w, line)
	}
	return w.Flush()
  }

