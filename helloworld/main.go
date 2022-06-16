package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parammeter is :", *name)
	fullStr := fmt.Sprintf("Hello %s form GO\n", *name)
	fmt.Println(fullStr)
}
