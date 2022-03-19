package main

import "flag"

func main() {
	var q string
	flag.StringVar(&q, "q", "", "package to search")
	flag.Parse()

	items := search(q)
	for _, item := range items {
		item.Print()
	}
}
