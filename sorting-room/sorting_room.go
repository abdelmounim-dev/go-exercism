package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
  return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}


// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
  // return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
  floatNum := float64(nb.Number())
  return fmt.Sprintf("This is a box containing the number %.1f", floatNum)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
  _, ok := fnb.(FancyNumber)
  if !ok {
    return 0
  }
  val, err := strconv.Atoi(fnb.Value())
  if err != nil {
    return 0
  }
  return val
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
  msg := "This is a fancy box containing the number"
  num :=ExtractFancyNumber(fnb)
  if num == 0 {
    return msg + " 0.0"
}
  return fmt.Sprintf("%v %v.0", msg, num)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
  switch i.(type) {
  case int:
    val := i.(int)
    return DescribeNumber(float64(val))
  case float64:
    val := i.(float64)
    return DescribeNumber(val)
  case NumberBox:
    val := i.(NumberBox)
    return DescribeNumberBox(val)
  case FancyNumberBox:
    val := i.(FancyNumberBox)
    return DescribeFancyNumberBox(val)
  default:
    return "Return to sender"
  }
}
