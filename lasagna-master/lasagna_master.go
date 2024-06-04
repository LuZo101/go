package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgLayerTime int) int {
	if avgLayerTime == 0 {
		prepT := len(layers) * 2
		return prepT
	} else {

		return len(layers) * avgLayerTime
	}
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {

	cSauce := 0.0
	cNoodles := 0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			cSauce += 0.2
		}
		if layers[i] == "noodles" {
			cNoodles += 50
		}
	}
	return cNoodles, cSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friends []string, myList []string) []string {

	indexSecret := len(friends) - 1
	secret := friends[indexSecret]
	myList[len(myList)-1] = secret
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantites []float64, portions int) []float64 {

	// slice_name := make([]type, length, capacity)
	newList := make([]float64, len(quantites))

	for i := 0; i < len(newList); i++ {
		newQuant := (quantites[i] / 2.0) * float64(portions)
		newList[i] = newQuant
	}
	return newList
}
