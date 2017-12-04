package main

import "fmt";

// 0 -> right
// 1 -> up
// 2 -> left
// 3 -> right

func main() {
    input := 83
    steps := 0

    number := 1
    loopNumber := 0

        for dir := 0; dir <= 3; dir++ {
            if dir == 2 || dir == 0 {
                steps++
            }
            for step := 1; step <= steps; step++ {
                if number >= input { break }
                number++
                fmt.Printf("dir: %d, step: %d of %d steps, number: %d \n", dir, step, steps, number)
            }
            if dir == 3 {
                dir = 0
                loopNumber++
            }
            if number >= input { break }
        }
    
}
