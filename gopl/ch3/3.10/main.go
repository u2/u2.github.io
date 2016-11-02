package main

import (
    "fmt"
    "bytes"
)

func main() {
    result := comma("12345677")
    fmt.Println(result)
}
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
    var ss []byte = []byte(s)
    var buf bytes.Buffer
    length := len(s)
    for i, v := range ss {
        buf.WriteString(string(v))
        if i > 0 && i != (length - 1) && ((length - i - 1)) % 3 == 0 {
            buf.WriteString(", ")
        }
    }
    return buf.String()
}
