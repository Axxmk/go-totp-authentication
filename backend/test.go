package main

import "fmt"

func main() {
	x, _ := evenOnly(10)
	fmt.Println(x)
	
}

func evenOnly(n int) (int, error) {
	if n%2 == 0 {
		return n / 2, nil
	}
	return 0, fmt.Errorf("not even")
}
