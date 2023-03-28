package main

import (
	"fmt"
	"math/rand"
)

func main() {
    var (
        rand_nums [3][5] int
        max_num int
    )

    for i := 0; i < cap(rand_nums); i++ {
        for j := 0; j < cap(rand_nums[i]); j++ {
            rand_nums[i][j] = rand.Intn(100)
        }
    }

    for i := 0; i < cap(rand_nums); i++ {
        for j := 0; j < cap(rand_nums[i]); j++ {
            if rand_nums[i][j] > max_num {
                max_num = rand_nums[i][j]
            }
        }
    }

    fmt.Println(rand_nums)
    fmt.Println(max_num)
}