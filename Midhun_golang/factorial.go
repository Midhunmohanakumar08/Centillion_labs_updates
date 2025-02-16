package main
import "fmt"
func factorialIterative(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}
func factorialRecursive(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorialRecursive(n-1)
}

func main() {
    var number int
    fmt.Print("Enter a number to calculate its factorial: ")
    fmt.Scanln(&number)
    if number < 0 {
        fmt.Println("Factorial is not defined for negative numbers.")
        return
    }
    iterativeResult := factorialIterative(number)
    fmt.Printf("Factorial (Iterative) of %d is: %d\n", number, iterativeResult)
    recursiveResult := factorialRecursive(number)
    fmt.Printf("Factorial (Recursive) of %d is: %d\n", number, recursiveResult)
}
