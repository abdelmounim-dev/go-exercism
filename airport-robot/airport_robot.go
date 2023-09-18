package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
  LanguageName() string
  Greet(a string) string
}


func SayHello(name string, g Greeter)  string {
  return "I can speak " + g.LanguageName() + ": " + g.Greet(name)
}
