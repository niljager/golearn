package main

import (
	"fmt"
)

func main() {
    romans := map[string]int {
        "I": 1, //II - 2 ...
        "IV": 4,
        "V": 5, //VI - 6 ...
        "IX": 9,
        "X": 10,
        "XL": 40,
        "L": 50, //lx - 60 ...
        "XC": 90,
        "C": 100,
        "CD": 400,
        "D": 500, //dcc - 600 ...
        "CM": 900,
        "M": 1000,
    }

    var (
        roman_num string
        arabic_num int
    )

    roman_num = "MCXI"

    for i :=0; i <= len(roman_num) - 1; i++ {
        if  i < len(roman_num) - 1 && 
                ((string(roman_num[i]) == "I" && string(roman_num[i + 1]) == "V") || 
                (string(roman_num[i]) == "I" && string(roman_num[i + 1]) == "X") ||
                (string(roman_num[i]) == "X" && string(roman_num[i + 1]) == "L") ||
                (string(roman_num[i]) == "X" && string(roman_num[i + 1]) == "C") ||
                (string(roman_num[i]) == "C" && string(roman_num[i + 1]) == "D") ||
                (string(roman_num[i]) == "C" && string(roman_num[i + 1]) == "M")) {
            val, ok := romans[string(roman_num[i]) + string(roman_num[i+1])];
            if ok {
                arabic_num += val
            }
            i++
        } else {
            val, ok := romans[string(roman_num[i])];
            if ok {
                arabic_num += val
            }
        }
    }

    fmt.Println(arabic_num)
}