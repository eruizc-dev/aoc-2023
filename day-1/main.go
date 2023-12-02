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
    //fmt.Printf("%s: %s, %s\n", s, string(first), string(last))
    return strconv.Atoi(string(first) + string(last))
}

func find_first(s string) byte {
    for i := 0; i < len(s); i++ {
        if s[i] >= '1' && s[i] <= '9' {
            return s[i]
        }

        if i < len(s) - 5 {
            if s[i] == 't' && s[i+1] == 'h' && s[i+2] == 'r'  && s[i+3] == 'e' && s[i+4] == 'e' {
                return '3'
            } else if s[i] == 'e' && s[i+1] == 'e' && s[i+2] == 'r'  && s[i+3] == 'h' && s[i+4] == 't' {
                return '3'
            } else if s[i] == 's' && s[i+1] == 'e' && s[i+2] == 'v'  && s[i+3] == 'e' && s[i+4] == 'n' {
                return '7'
            } else if s[i] == 'n' && s[i+1] == 'e' && s[i+2] == 'v'  && s[i+3] == 'e' && s[i+4] == 's' {
                return '7'
            } else if s[i] == 'e' && s[i+1] == 'i' && s[i+2] == 'g'  && s[i+3] == 'h' && s[i+4] == 't' {
                return '8'
            } else if s[i] == 't' && s[i+1] == 'h' && s[i+2] == 'g'  && s[i+3] == 'i' && s[i+4] == 'e' {
                return '8'
            }
        }

        if i < len(s) - 4 {
            if s[i] == 'f' && s[i+1] == 'o' && s[i+2] == 'u' && s[i+3] == 'r' {
                return '4'
            } else if s[i] == 'r' && s[i+1] == 'u' && s[i+2] == 'o' && s[i+3] == 'f' {
                return '4'
            } else if s[i] == 'f' && s[i+1] == 'i' && s[i+2] == 'v' && s[i+3] == 'e' {
                return '5'
            } else if s[i] == 'e' && s[i+1] == 'v' && s[i+2] == 'i' && s[i+3] == 'f' {
                return '5'
            } else if s[i] == 'n' && s[i+1] == 'i' && s[i+2] == 'n' && s[i+3] == 'e' {
                return '9'
            } else if s[i] == 'e' && s[i+1] == 'n' && s[i+2] == 'i' && s[i+3] == 'n' {
                return '9'
            }
        }

        if i < len(s) - 3 {
            if s[i] == 'o' && s[i+1] == 'n' && s[i+2] == 'e' {
                return '1'
            } else if s[i] == 'e' && s[i+1] == 'n' && s[i+2] == 'o' {
                return '1'
            } else if s[i] == 't' && s[i+1] == 'w' && s[i+2] == 'o' {
                return '2'
            } else if s[i] == 'o' && s[i+1] == 'w' && s[i+2] == 't' {
                return '2'
            } else if s[i] == 's' && s[i+1] == 'i' && s[i+2] == 'x' {
                return '6'
            } else if s[i] == 'x' && s[i+1] == 'i' && s[i+2] == 's' {
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
