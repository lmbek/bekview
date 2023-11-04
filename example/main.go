package main

import "bekview"

func main() {
	err := bekview.Start()
	if err != nil {
		panic(err)
	}
}
