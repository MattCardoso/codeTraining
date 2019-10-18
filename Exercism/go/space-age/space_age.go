package space


type Planet string

func getOrbitalPeriod(p Planet) float64 {
	switch p {
		case "Earth":
			return 365.25
		case "Mercury":  
			return 0.2408467 * 365.25
		case "Venus":  
			return 0.61519726 * 365.25
		case "Mars":  
			return 1.8808158 * 365.25
		case "Jupiter":  
			return 11.862615 * 365.25
		case "Saturn":  
			return 29.447498 * 365.25
		case "Uranus":  
			return 84.016846 * 365.25
		case "Neptune":  
			return 164.79132 * 365.25
	}
	return 1.00
}
//Age(tc.seconds, tc.planet)
func Age(seconds float64, planet Planet) float64{
	orbitalPeriod := getOrbitalPeriod(planet);
	return seconds / (60 * 60 * 24 * orbitalPeriod)
}