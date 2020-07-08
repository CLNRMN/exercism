package space

type Planet string

const EarthSeconds = 31557600.0

// Create a map with the key type Planet and the value the OrbitalPeriod
var PlanetOrbitalPeriod = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Calculate Seconds from Planet X to Earth Years
func Age(seconds float64, pl Planet) float64 {
	// seconds / 31557600 * orbital period = Age on Earth
	return (seconds / EarthSeconds) / PlanetOrbitalPeriod[pl]
}
