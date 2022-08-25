package main
//Адаптер предусматривает создание класса оболочки с требуемым интерфейсом

//Адаптер реализует интерфейс недоступного для изменения модуля fSensor, являясь оберткой
type adapterCtoF struct{
	sensor *cSensor
}

func (a adapterCtoF) getTemperatureF() fahrenheit{
	return fahrenheit(a.sensor.getTemperatureC() * 1.8 + 32)
}

func (a adapterCtoF) setTemperatureF(t fahrenheit){
	a.sensor.temperature = celsius((t - 32) / 1.8)
}
