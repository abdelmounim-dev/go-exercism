package lasagna

// import "strings"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) (totalTime int) {
  totalTime = len(layers) * map[bool]int {true: 2, false: time} [time == 0]
  return
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
  noodles, sauce = 0, 0.0
  for _, v := range layers {
    switch v {
    case "noodles":
      noodles += 50
      break;
    case "sauce":
      sauce += .2
      break;
    }
  }

  return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
  fLen, mLen := len(friendsList), len(myList)
  myList[mLen - 1] = friendsList[fLen - 1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {

  var scaled []float64

  mul := float64(portions )/ 2.0
  for _, v := range quantities {
    scaled = append(scaled, v * mul )
  }
  return scaled
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
