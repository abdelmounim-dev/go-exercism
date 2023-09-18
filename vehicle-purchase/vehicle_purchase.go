package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
// TODO maybe not the most readable code ever? IDK
  return (map[bool]string{true: option1, false:option2} [option1 < option2]) +
  " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
  if age < float64(3) {return originalPrice * 0.8}
  if age >= float64(10) {return originalPrice * 0.5}
	return originalPrice * .7

}
