package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("./data.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // fmt.Println("Successfully Opened user_access.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // Convert json file to Map
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    fmt.Printf("%d", result)

}
