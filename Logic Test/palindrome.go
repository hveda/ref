// go run ./palindrome.go malam
package main
  
import (
    "fmt"
    "os"
)

func reverse(s string) string {
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
    return string(rns)
}
  
func main() {
    if len(os.Args) <= 1 {
        fmt.Printf("No argument given")
        os.Exit(1)
    }
    str := os.Args[1]
  
    // returns the reversed string.
    strRev := reverse(str)

    if str == strRev {
        fmt.Println("Palindrome")
    } else {
        fmt.Println("NOT palindrome")
    }
}
