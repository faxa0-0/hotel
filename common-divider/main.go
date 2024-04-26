package main

import "fmt"

func main() {
    fmt.Println(findCommonDividersArray([]int{42, 12, 18}))
}

func findCommonDividersArray(arr []int) (resultArray []int) {
    if len(arr) == 0 {
        return resultArray
    }

    commonDividersMap := make(map[int]int)
    for _, v := range arr {
        for cm := range findCommonDividers(v) {
            if val, ok := commonDividersMap[cm]; ok {
                commonDividersMap[cm] = val + 1
            } else {
                commonDividersMap[cm] = 1
            }
        }
    }

    for k, v := range commonDividersMap {
        if v == len(arr) {
            resultArray = append(resultArray, k)
        }
    }

    return resultArray
}

func findCommonDividers(number int) (resultMap map[int]bool) {
    resultMap = make(map[int]bool)
    if number <= 1 {
        return resultMap
    }
    for i := 2; i <= number; i++ {
        if number%i == 0 {
            resultMap[i] = true
        }
    }
    return resultMap
}