package main

type celsius float32

type SensorC interface {
	getTemperatureС() fahrenheit
	setTemperatureС(t fahrenheit)
}

//cSensor Датчик который работает с цельсия и имеет другой интерфейс несовместимый с SensorF
type cSensor struct {
	temperature celsius
}

func (ts cSensor) getTemperatureC() celsius{
	return ts.temperature
}

func (ts cSensor) setTemperatureC(t celsius) {
	ts.temperature = t
}

