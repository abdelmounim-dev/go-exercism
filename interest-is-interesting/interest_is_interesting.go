package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {

  switch  {
  case balance < 0:
    return 3.213
  case balance < 1000:
    return  .5
  case balance < 5000:
    return 1.621
  default:
    return 2.475
  }

}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
  return balance + Interest(balance)
}

// iterationo is  used in the recursive function  below.
var iteration = 0
// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
  // this function is  done recursively, might be more efficient if done in the conventinal way
  if balance * targetBalance < 0 {
    return -1
  }

  if balance >= targetBalance {
    value := iteration
    iteration = 0
    return value
  }
  iteration++
  return YearsBeforeDesiredBalance(AnnualBalanceUpdate(balance), targetBalance)
}
