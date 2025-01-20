package lasagna

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

func Quantities(layers []string) (int, float64) {
	var noodlesLayers, sauceLayers int
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodlesLayers++
		} else if layers[i] == "sauce" {
			sauceLayers++
		}

	}
	noodles := noodlesLayers * 50
	sauce := float64(sauceLayers) * 0.2

	return noodles, sauce

}

func AddSecretIngredient(friendList, ownList []string) {
	ownList[len(ownList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {

	scaledAmounts := make([]float64, len(amounts))
	for i := 0; i < len(amounts); i++ {
		scaledAmounts[i] = amounts[i] / 2 * float64(portions)
	}
	return scaledAmounts

}

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
