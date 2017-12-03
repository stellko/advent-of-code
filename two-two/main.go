package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
    sum := 0

    lines := getInput()

    for _, line := range lines {
        values := strings.Split(line, "\t")

        for _, value := range values {
            intValue, err := strconv.Atoi(string(value))
            if (err == nil) {
                for _, otherValue := range values {
                    intOtherValue, err := strconv.Atoi(string(otherValue))
                    if (err == nil) {
                        if intValue != intOtherValue {
                            if (intValue % intOtherValue == 0) {
                                sum += intValue / intOtherValue
                            }

                        }
                    }
                }
            }
        } 
	}
    
	fmt.Println(sum)
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
