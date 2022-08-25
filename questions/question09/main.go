package main

import "fmt"

func main(){
	s1 := make([]int, 10)
	s2 := new([10]int)[:10]
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(s1, s2, s3)
}
