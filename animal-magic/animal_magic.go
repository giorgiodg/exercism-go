package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	random := rand.Intn(20)
	if random < 1 {
		return 1
	} else {
		return random
	}
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	min := 0.0
	max := 12.0
	return (min + rand.Float64()*(max-min))
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"cat", "dog", "elephant", "giraffe", "fox", "ant", "beaver", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) { animals[i], animals[j] = animals[j], animals[i] })
	return animals
}
