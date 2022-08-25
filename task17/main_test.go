package main

import "testing"

func TestBinSearch2(t *testing.T) {
	testTable := []struct {
		slice  []int
		target int
		want   int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			3,
			2,
		},
		{
			[]int{1, 2, 300, 4032},
			300,
			2,
		},
		{
			[]int{1}, //Массив состоит из 1 значения
			1,
			0,
		},
		{
			[]int{}, //Поиск в пустом массиве
			3,
			-1,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8}, //Несущ. элемент
			306,
			-1,
		},
	}

	for _, testCase := range testTable {
		if got := binSearch2(testCase.slice, testCase.target); testCase.want != got {
			t.Errorf("Для slice: %v, target: %v ожидалось %v. Получено %v", testCase.slice, testCase.target, testCase.want, got)
		}
	}
}
