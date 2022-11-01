package radar

import "math"

var EarthRadius float64 = 6371210 // радиус земли
var DISTANCE float64 = 20000      // константа для перых версий тестов (будем принимать в будущем)

// Computer Delta ...
func ComputeDelta(degrees float64) float64 {
	sinus, cosinus := math.Sincos(Deg2rad(degrees))
	print(sinus)
	print(cosinus)

	delta := math.Pi / 180 * EarthRadius * cosinus
	return delta
}

// Deg2Rad ...
func Deg2rad(degrees float64) float64 {
	rad := degrees * math.Pi / 180 // определяем радиус
	return rad                     // возвращаем радиус
}

// Is Radius ...
func IsRadius() (float64, float64) {
	latitude := 55.460531  // интересующие нас координаты широты (будем принимать в будущем)
	longitude := 37.210488 // интересующие нас координаты долготы (будем принимать в будущем)

	deltaLat := ComputeDelta(latitude)  // получаем дельту по широте
	deltaLon := ComputeDelta(longitude) // получаем дельту по долготе

	aroundLat := DISTANCE / deltaLat // вычисляем диапазон координат по широте
	aroundLon := DISTANCE / deltaLon // вычисляем диапазон координат по долготе

	return aroundLat, aroundLon // возвращаем диапазон координат по ш. и д.
}
