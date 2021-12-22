package lasagna

func PreparationTime(layers []string, timePerlayer int) int {
	minutesToAdd := timePerlayer
	if timePerlayer == 0 {
		minutesToAdd = 2
	}

	return (len(layers) * minutesToAdd)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles, sauce := 0, 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}

		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(list1 []string, list2 []string) []string {
	result := list2
	result = append(result, list1[len(list1)-1])
	return result
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsForPortions []float64, portions int) []float64 {
	result := make([]float64, len(amountsForPortions))
	for idx, val := range amountsForPortions {
		result[idx] = (val / 2) * float64(portions)
	}

	return result
}
