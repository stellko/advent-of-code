package main

import "fmt"
import "strings"
import "os"
import "bufio"
import "sort"

func main() {   

    sum := 0
    lines := getInput()
    linesWithoutDuplicates := []string{}

    for _, line := range lines {
        values := strings.Split(line, "\t")

        for _, value := range values {
            
            if !hasDuplicates(strings.Fields(value)) {
                linesWithoutDuplicates = append(linesWithoutDuplicates, value)
            }
        }
    }

    for _, value := range linesWithoutDuplicates {
        if !hasAnagram(strings.Fields(value)) {
            sum++
        }
    }

	fmt.Println(sum)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func hasAnagram(elements []string) bool {
    
    for i := 0; i < len(elements); i++ {
        for v := 0; v < i; v++ {
            sortedWeerd := SortString(elements[v])
            sortedOtherWord := SortString(elements[i])
            if sortedWeerd == sortedOtherWord {
                return true
            }
        }
    }
    return false
}

func hasDuplicates(elements []string) bool {
    
    for i := 0; i < len(elements); i++ {
        for v := 0; v < i; v++ {
            if elements[v] == elements[i] {
                return true
            }
        }
    }
    return false
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