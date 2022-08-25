package main

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Set map[int]struct{}

//newRandSet вернет множество с случайными элементами
func newRandSet(count int) Set{
	retSet := make(map[int]struct{}, count)

	for i := 0; i < count; i++ {
		retSet[rand.Intn(20)] = struct{}{}
	}

	return retSet
}

func (s Set) String() string{
	builder := strings.Builder{}
	builder.WriteRune('{')
	for k, _ := range s{
		builder.WriteString(fmt.Sprintf(" %v ", k))
	}
	builder.WriteRune('}')

	return builder.String()
}

func intersect(s1, s2 Set) Set{
	ret := Set{}

	for k, _ := range s2{
		if _, ok := s1[k]; ok {
			ret[k] = struct{}{}
		}
	}

	return ret
}

func main(){
	rand.Seed(time.Now().UnixNano())
	set1 := newRandSet(12)
	set2 := newRandSet(6)

	fmt.Println(fmt.Sprintf("Set1: %v", set1))
	fmt.Println(fmt.Sprintf("Set2: %v\n", set2))

	fmt.Println("intersect: ", intersect(set1, set2))
}
