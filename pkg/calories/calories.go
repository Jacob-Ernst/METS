package calories

// For calculating Calories per minute.
const calorieConst float64 = 3.5

// Calculate Calories/min burnt.
func BurnRate(MET, kg float64) float64 {
	return (MET * calorieConst * kg) / 200
}

func TotalBurn(kg, MET, time float64) float64 {
	calPerMinute := BurnRate(MET, kg)

	return calPerMinute * time
}
