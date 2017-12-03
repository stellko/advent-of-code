package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

func main() {
    sum := 0

    lines := getInput()

    for _, line := range lines {
        min := math.MaxInt32
        max := 0
        values := strings.Split(line, "\t")

        for _, value := range values {
            intValue, err := strconv.Atoi(string(value))
            if (err == nil) {
                if (intValue < min) {
                    min = intValue
                }
                if (intValue > max) {
                    max = intValue
                }
            }
        }

        diff := max - min
        
        sum += diff
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
