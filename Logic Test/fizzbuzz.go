// go run ./fizzbuzz.go 15
package main
  
import (
    "fmt"
    "os"
    "strconv"
)
  
func main() {
    if len(os.Args) <= 1 {
        fmt.Printf("Not enough argument")
        os.Exit(1)
    }

    num,_ := strconv.Atoi(os.Args[1])
    data := make([]string, num)

    for i := 1; i <= num; i++ {
        if i%3 == 0 && i%5 == 0 {
            data[i-1] = "FizzBuzz"
        } else if i%3 == 0 {
            data[i-1] = "Fizz"
        } else if i%5 == 0 {
            data[i-1] = "Buzz"
        } else {
            data[i-1] = strconv.Itoa(i)
        }
    }
    fmt.Printf("%v",data) 
}
