package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"

    "github.com/stretchr/stew/slice"
)

func main() {
    lst := []string{"sh", "gl", "ch", "ph", "tr", "br", "fr", "bl", "gr", "st", "sl", "cl", "pl", "fl", "th"}
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Type what you would like translated into pig-latin and press ENTER: ")
    sentenceraw, _ := reader.ReadString('\n')
    sentence := strings.Split(strings.TrimSpace(sentenceraw), " ")
    isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
    newsentence := make([]string, 0)
    for _, i := range sentence {
        if slice.Contains([]byte{'a', 'e', 'i', 'o', 'u'}, i[0]) {
            newsentence = append(newsentence, strings.Join([]string{string(i), "ay"}, ""))
        } else if slice.Contains(lst, string(i[0])+string(i[1])) {
            newsentence = append(newsentence, strings.Join([]string{string(i[2:]), string(i[:2]), "ay"}, ""))
        } else if !isAlpha(string(i)) {
            newsentence = append(newsentence, strings.Join([]string{string(i)}, ""))
        } else {
            newsentence = append(newsentence, strings.Join([]string{string(i[1:]), string(i[0]), "ay"}, ""))
        }
    }
    fmt.Println(strings.Join(newsentence, " "))
}
