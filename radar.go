package radar

import "math"

const earthRadius = 6371210 // радиус земли
var DISTANCE float64 = 20000

// Computer Delta ...
func computeDelta(degrees float64) float64 {
	sinus, cosinus := math.Sincos(deg2rad(degrees))
	print(sinus)
	print(cosinus)

	delta := math.Pi / 180 * earthRadius * cosinus
	return delta
}

func deg2rad(degrees float64) float64 {
	rad := degrees * math.Pi / 180
	return rad
}

func isRadius() (float64, float64) {
	latitude := 55.460531  //Интересующие нас координаты широты
	longitude := 37.210488 //Интересующие нас координаты долготы

	deltaLat := computeDelta(latitude)  //Получаем дельту по широте
	deltaLon := computeDelta(longitude) // Дельту по долготе

	aroundLat := DISTANCE / deltaLat // Вычисляем диапазон координат по широте
	aroundLon := DISTANCE / deltaLon // Вычисляем диапазон координат по долготе

	return aroundLat, aroundLon
}
