package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %v!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
  years := map[bool]string {true: "year", false: "years"} [age == 1]

	return fmt.Sprintf("Happy birthday %v! You are now %v %v old!", name, age, years)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
  return fmt.Sprintf("%v\n" +
    "You have been assigned to table %03d" +
    ". Your table is %v, exactly %.1f meters from here.\n" +
    "You will be sitting next to %v.",
    Welcome(name),
    table,
    direction, distance,
    neighbor)
}
