package main

import "fmt"

func main() {

	Run(&Chinese{})

	Run(&English{})

}

func Run(service Greet) {
	fmt.Println(service.Hello())
	fmt.Println(service.HowAreYou())
}
