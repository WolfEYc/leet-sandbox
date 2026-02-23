package main

import "log"

func main() {
	n := 6
	var i int
	for i = range n {
		log.Println(i)
	}
	log.Println(i)
}
