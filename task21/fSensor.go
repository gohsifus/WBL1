package main

type fahrenheit float32

type SensorF interface {
	getTemperatureF() fahrenheit
	setTemperatureF(t fahrenheit)
}

//fSensor датчик температуры, работающий с температурой в фаренгейтах
//по какой-то причине мы не можем переписать/дописать эту сущность, например это сторонний пакет,
//который используется в программе где-то еще
type fSensor struct {
	temperature fahrenheit
}

func (ts fSensor) getTemperatureF() fahrenheit{
	return ts.temperature
}

func (ts fSensor) setTemperatureF(t fahrenheit) {
	ts.temperature = t
}

