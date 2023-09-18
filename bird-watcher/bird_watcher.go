package birdwatcher

// import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
  sum := 0

  for _, v := range (birdsPerDay) {
    sum += v
  }
  return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

  sum := 0
  startOfWeek := (week -1) * 7
  endOfWeek := (week -1) * 7 + 7

  for i := startOfWeek; i >= startOfWeek && i < endOfWeek ; i++{
    // fmt.Println(i, startOfWeek, endOfWeek, birdsPerDay[i])
   sum += birdsPerDay[i]
}

  return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
  for i := range(birdsPerDay) {
    if (i % 2) == 0 {birdsPerDay[i]++}
  }

  return birdsPerDay
}
