package fuzz

import "strconv"
import "github.com/nbari/violetear/v7"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)

            length := len(content)
            str1 := content[0:length/2]
            str2 := content[length/2:length-1]

            var test violetear.Trie
            test.Get(str1, str2)
            return 0

        case 1:
            content := string(bytes)
            var test violetear.Trie
            test.SplitPath(content)
            return 0

        case 2:
            content := string(bytes)

            length := len(content)
            str1 := content[0:length/2]
            str2 := content[length/2:length-1]

            var test violetear.Params
            test.Add(str1, str2)
            return 0

        default:
            content := string(bytes)
            var test violetear.Trie
            test.Name(content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}