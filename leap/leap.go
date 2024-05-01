// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear determines wheather a given year is leap or not
func IsLeapYear(year int) bool {
	return (year%400 == 0) || (year%4 == 0) && !(year%100 == 0)
}
