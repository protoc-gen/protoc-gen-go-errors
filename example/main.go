package main

import "log"

func main() {
	err := ErrorUnauthorized()
	if !IsUnauthorized(err) {
		log.Fatalln("error")
	}

	ExampleI18nErrors()

	log.Println("ok")
}
