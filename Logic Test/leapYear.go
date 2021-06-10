// go run ./leapYear.go 1900 2020
package main
  
import (
    "fmt"
    "os"
    "strconv"
)
  
func main() {
    if len(os.Args) <= 2 {
        fmt.Printf("Not enough argument")
        os.Exit(1)
    }

    start,_ := strconv.Atoi(os.Args[1])
    end,_ := strconv.Atoi(os.Args[2])

    year := start

    for year < end {
        if year % 4 == 0  {
            if year % 100 == 0 {
                if year % 400 == 0 {
                    fmt.Printf("%d, ", year)
                    year += 4
                } else {
                    year++
                }
            } else {
                fmt.Printf("%d, ", year)
                year += 4
            }
        } else {
            year++
        }
    }
}
