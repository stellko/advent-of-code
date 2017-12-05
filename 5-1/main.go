package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {   

    steps := 0
    elements := getInput()
    i := 0;
    jumps := 0
    intList := []int{}

    for _, value := range elements {
        number, err := strconv.Atoi(value)
        if (err == nil) {
            intList = append(intList, number)
        }
    }

    for i < len(intList) {
        i = i + jumps
        if i >= len(intList) { break }
        jumps = intList[i]
        intList[i] += 1
        steps++
    }
  
	fmt.Println(steps)
}

func getInput() []string {
    file, err := os.Open("input.txt")
    
    if err != nil {
    fmt.Println("Failed to open file")
    }
    defer file.Close()
    
    lines := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    lines = append(lines, scanner.Text())
    }
    
    return lines
}