// go run ./fibo.go 15 1 3
package main
  
import (
    "fmt"
    "math"
    "os"
    "strconv"
)

func fib() func() int {
  a, b := 0, 1
  return func() int {
    a, b = b, a+b
    return a
  }
}
  
func main() {
    data := os.Args[1:]
    var nums = []int{}
    for _, i := range data {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        nums = append(nums, j)
    }
    sum := 0  
    for _, v := range nums {  
        sum += v
    }

    f := fib()
    fres := []int{}
    res := 0
    for res <= sum {
        fres = append(fres, f())
        res = fres[len(fres)-1]
    }


    if int64(math.Abs(float64(sum - fres[len(fres)-2]))) < int64(math.Abs(float64(fres[len(fres)-1] - sum))) {
        fmt.Println(int64(math.Abs(float64(sum - fres[len(fres)-2]))))
    } else {
        fmt.Println(int64(math.Abs(float64(fres[len(fres)-1] - sum))))
    }

}
