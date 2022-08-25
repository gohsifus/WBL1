package main

import "fmt"

func main(){
	//var i *[]int
	i := new([]int)
	fmt.Println(i, len(*i), cap(*i))                   // &[] 0 0

	//var istring *[]string
	istring := new([]string)
	fmt.Println(istring, len(*istring), cap(*istring)) // &[] 0 0

	im := make([]int, 10)
	fmt.Println(im, len(im), cap(im))

	ims := make([]string, 10)
	fmt.Println(ims, len(ims), cap(ims))
}
