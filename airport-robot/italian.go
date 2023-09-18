package airportrobot

type Italian struct {
}

func (i Italian) Greet(name string) string {
  return "Ciao " + name + "!"
}

func (i Italian) LanguageName() string {
  return "Italian"
}
