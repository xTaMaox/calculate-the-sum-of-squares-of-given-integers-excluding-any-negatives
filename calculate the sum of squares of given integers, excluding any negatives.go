// Calculate the sum of squares of given integers, excluding any negatives In Go lang
// Code written by Muhamad Adi Prasetyo

// Enter Total number of test cases := N
// N Times :
// Enter Total number of values.    := num
// Enter [num] times int values
// sample input:
// 2
// 4
// 3 -1 1 14
// 5
// 9 6 -53 32 16

// sample output :
// 206
// 1397


// sum of square, only positive numbers without using any loop
// Thats why i use recurssion in Go

// my Code in Go lang

package main

import "fmt"

// Taking Number of Test Cases and values

func test_cases(n int) {
    if n <= 0 {
        return
    }
    var num int
    // total number of values
    fmt.Scanf("%d", &num)
    fmt.Println(sum_of_square(num))
    test_cases(n-1)
}

// Calculating sum of square for each test case

func sumOfSquare(X int) int {
    if X == 0 {
        return 0
    }
    var Y int
    fmt.Scanf("%d", &Y)
    if Y > 0 {
        return Y*Y + sumOfSquare(X-1)
    }
    return sumOfSquare(X-1)
}

func sum_of_square(value_count int) int {
    if value_count == 0 {
        return 0
    }
    var value int
    // take input value for generating sum of square
    fmt.Scanf("%d", &value)
    // if only value is positive
    if value > 0 {
        return value*value + sum_of_square(value_count - 1)
    }
    return sum_of_square(value_count - 1)
}

// Main function

func main() {
    var N int
    // number of total Test Cases
    fmt.Scanf("%d", &N)
    // Take input for each Test Case
    test_cases(N)
} 