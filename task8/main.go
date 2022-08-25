package main

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

import (
	"fmt"
)

//Вернет значение i-го бита
func getBitValue(target int64, i int) bool {
	//Сдвигаем 1 на i - 1 чтобы получить маску вида ..000100..
	//где i бит установлен в 1 а остальные 0
	var mask int64 = 1 << (i - 1)
	// после применения маски либо все биты target устанавливаются в 0 либо все кроме бита i посокльку он 1
	if target&mask == mask {
		return true
	}
	return false
}

func setBitValue(target int64, i int, value bool) int64 {
	//1 бит зарезервирован для знака - чтобы его инвертировать надо умножить число на -1
	if i == 64 {
		if value == (target < 0){
			return target
		}
		return target * -1
	}

	bitValue := getBitValue(target, i)
	//Если бит уже установлен в нужное состояние
	if bitValue == value {
		return target
	}

	var mask int64 = 1 << (i - 1)
	if bitValue { //было 1 надо 0
		// %^ операция И Не - сбрасывает биты target если биты mask 1
		// 0110 &^ 0010 = 0100
		target = target &^ mask
	} else { //было 0 надо 1
		// 0100 | 0010 = 0110
		target = target|mask
	}

	return target
}

func main() {
	var number int64 = 123
	fmt.Println(fmt.Sprintf("%64b %v", number, number))
	fmt.Println("setBitValue")
	fmt.Println(fmt.Sprintf("%64b %v", setBitValue(number, 3, true), setBitValue(number, 3, true)))
}
