package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) (out []Record) {
	for _, r := range in {
	  if predicate(r) {
      out = append(out, r)
    }
	}
  return
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
  return func (r Record) bool {
    if (p.From <= r.Day && p.To >= r.Day) {
      return true
    }
    return false
  }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
  return func (r Record) bool {
    if (r.Category == c) {
      return true
    }
    return false
  }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) (total float64) {
  for _, r := range Filter(in, ByDaysPeriod(p)) {
    total += r.Amount;
  }

  return
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (total float64, err error) {
  record1 :=Filter(in, ByCategory(c))
  if len(record1) == 0 {
    return 0, errors.New("unknown category " + c)
  }
  record2 :=Filter(record1, ByDaysPeriod(p))
  // if len(record2) == 0 {
  //   return 0, errors.New("no reccords of category " + c + " in the given period")
  // }
  for _, r := range record2  {
    total += r.Amount;
  }
  return
}
