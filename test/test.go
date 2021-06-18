package main

import "fmt"

func main() {
	maap := make(map[string][][]string)
	maap["make"] = [][]string{{"1", "2"}}
	maap["make"] = append(maap["make"], []string{"1", "2"})
	vare, ok := maap["make"]
	vare = append(vare, []string{"3", "4"})
	fmt.Println(vare)
	fmt.Println(ok)
	fmt.Println(maap)
}
