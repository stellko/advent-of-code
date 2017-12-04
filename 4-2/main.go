package main

import "fmt"
import "strings"
import "os"
import "bufio"


func main() {   

    sum := 0
    lines := getInput()

    for _, line := range lines {
        values := strings.Split(line, "\t")

        for _, value := range values {
            if !hasDuplicates(strings.Fields(value)) {
                sum++
            }
        } 
	}

	fmt.Println(sum)
}

func hasDuplicates(elements []string) bool {
    result := []string{}

    for i := 0; i < len(elements); i++ {
        exists := false
        for v := 0; v < i; v++ {
            if elements[v] == elements[i] {
                exists = true
                return true
            }
        }

        if !exists {
            result = append(result, elements[i])
        } else {
            return true
        }
    }
    return false
}

func hasAnagram(elements []string) bool {

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