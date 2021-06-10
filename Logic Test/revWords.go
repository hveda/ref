// go run ./revWords.go "I am A Great human"
package main
  
import (
    "fmt"
    "os"
    "strings"
)

func isUpper(v byte) bool {
    return 'A' <= v && v <= 'Z'
}

func reverse(s string) string {
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
    return string(rns)
}

func main() {
    if len(os.Args) <= 1 {
        fmt.Printf("Not enough argument")
        os.Exit(1)
    }

    str := os.Args[1]
    s := strings.Split(str, " ")

    res := s
    for i, v := range s {
        res[i] = reverse(v)
    }

    fmt.Println(strings.Join(res, " "))
}
