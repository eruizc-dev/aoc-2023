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
    last := find_last(s)
    return strconv.Atoi(string(first) + string(last))
}

func find_first(s string) byte {
    for i := 0; i < len(s); i++ {
        num := find_num(s[i:])
        if num != '0' {
            return num
        }
    }
    return 0
}

func find_last(s string) byte {
    for i := len(s) - 1; i >= 0; i-- {
        num := find_num(s[i:])
        if num != '0' {
            return num
        }
    }
    return 0
}

func find_num(num string) byte {
    if num[0] >= '1' && num[0] <= '9' {
        return num[0]
    }

    if len(num) >= 5 {
        slice := num[0:5]
        if slice == "three" {
            return '3'
        } else if slice == "seven" {
            return '7'
        } else if slice == "eight" {
            return '8'
        }
    }

    if len(num) >= 4 {
        slice := num[0:4]
        if slice == "four" {
            return '4'
        } else if slice == "five" {
            return '5'
        } else if slice == "nine" {
            return '9'
        }
    }

    if len(num) >= 3 {
        slice := num[0:3]
        if slice == "one" {
            return '1'
        } else if slice == "two" {
            return '2'
        } else if slice == "six" {
            return '6'
        }
    }
    return '0'
}
