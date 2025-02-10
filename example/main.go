package main

import "log"

func main() {
	err := ErrorUnauthorized()
	if !IsUnauthorized(err) {
		log.Fatalln("error")
	}

	log.Println("ok")
}
