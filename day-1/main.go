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
    first := find_first(s)
    last := find_first(flip(s))
    return strconv.Atoi(string(first) + string(last))
}

func find_first(s string) byte {
    for i := 0; i < len(s); i++ {
        if s[i] >= '1' && s[i] <= '9' {
            return s[i]
        }

        if i < len(s) - 5 {
            num := s[i:i+5]
            if num == "three" || num == "eerht" {
                return '3'
            } else if num == "seven" || num == "neves" {
                return '7'
            } else if num == "eight" || num == "thgie" {
                return '8'
            }
        }

        if i < len(s) - 4 {
            num := s[i:i+4]
            if num == "four" || num == "ruof" {
                return '4'
            } else if num == "five" || num == "evif" {
                return '5'
            } else if num == "nine" || num == "enin" {
                return '9'
            }
        }

        if i < len(s) - 3 {
            num := s[i:i+3]
            if num == "one" || num == "eno" {
                return '1'
            } else if num == "two" || num == "owt" {
                return '2'
            } else if num == "six" || num == "xis" {
                return '6'
            }
        }
    }
    return 0
}

func flip(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
