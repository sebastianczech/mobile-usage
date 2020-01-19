package main

import "fmt"

func checkNetgear(url string) string {
	return "netgear: " + url
}

func main() {
	fmt.Println(checkNetgear("http://10.212.194.1/index.html"))
}
