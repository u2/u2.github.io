package main

import (
    "fmt"
    "bytes"
    "strings"
)

func main() {
    result := comma("12345677.33333333")
    fmt.Println(result)

    result = comma("12345677.3334")
    fmt.Println(result)

    result = comma("12345677.333")
    fmt.Println(result)

    result = comma("12345677.33")
    fmt.Println(result)

    result = comma("123")
    fmt.Println(result)

    result = comma("123456")
    fmt.Println(result)

    result = comma("1239945677")
    fmt.Println(result)
}
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
    var sls = strings.Split(s, ".")
    var buf bytes.Buffer
    var integralPart = sls[0]
    var integralPartBytes []byte = []byte(integralPart)
    length := len(integralPartBytes)
    for i, v := range integralPartBytes {
        buf.WriteString(string(v))
        if i != (length - 1) && ((length - i - 1)) % 3 == 0 {
            buf.WriteString(", ")
        }
    }

    if len(sls) == 2 {
        buf.WriteString(".")

        var fractionalPart = sls[1]
        var fractionalPartBytes []byte = []byte(fractionalPart)
        length = len(fractionalPartBytes)
        for i, v := range fractionalPartBytes {
            buf.WriteString(string(v))
            if i > 0 && i != (length - 1) && (i + 1) % 3 == 0 {
                buf.WriteString(", ")
            }
        }
    }

    return buf.String()
}
