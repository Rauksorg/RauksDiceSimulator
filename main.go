package main

import "fmt"

type count struct {
	one, two, three, four, five, six int
}
type results struct {
	count
	total int
}

func main() {
	var myResults results

	fmt.Println(myResults)
}
