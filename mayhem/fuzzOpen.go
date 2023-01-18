package fuzzOpen

import "github.com/apache/calcite-avatica-go/v5"


func mayhemit(bytes []byte) int {

    content := string(bytes)
    var test avatica.Driver
    test.Open(content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}