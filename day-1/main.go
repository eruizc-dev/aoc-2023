package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    document := read("input.txt")
    int_result := calibrate_document(document)
    fmt.Printf("1. Result: %d\n", int_result)
}

func read(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        panic("Could not open file")
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    index := 0
    result := make([]string, 1024)
    for scanner.Scan() {
        result[index] = scanner.Text()
        index++
    }
    return result
}

func calibrate_document(s []string) int {
    sum := 0
    for _, line := range s {
        res, err := calibrate_line(line)
        if err != nil {
            panic("error reading " + line)
        }
        sum += res
    }
    return sum
}

func calibrate_line(s string) (int, error) {
    if len(s) == 0 {
        return 0, nil
    }
    var first byte
    var last byte
    for i := 0; i <= len(s); i++ {
        if first == 0 && s[i] >= '1' && s[i] <= '9' {
            first = s[i]
        }

        if last == 0 && s[len(s) - 1 - i] >= '1' && s[len(s) - 1 - i] <= '9' {
            last = s[len(s) - 1 - i]
        }

        if first != 0 && last != 0 {
            break
        }
    }
    return strconv.Atoi(string(first) + string(last))
}
