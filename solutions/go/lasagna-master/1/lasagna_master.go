package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, meanTime int) int {
    if meanTime == 0 {
        meanTime = 2
    }

    return len(layers) * meanTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles, sauce := 0, 0.0
    for _, value := range layers {
        if value == "noodles" {
            noodles += 50
        } else if value == "sauce" {
            sauce += 0.2 
        }
    }

    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
    
    myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities  []float64, servingsCount int) []float64 {
    newQuantities := []float64 {} 
    for i := 0; i < len(quantities ); i++ {
        newQuantities = append(newQuantities, quantities[i] * (float64(servingsCount) / 2.0))
    }

    return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.