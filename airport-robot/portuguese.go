package airportrobot

type Portuguese struct {
}

func (p Portuguese) Greet(name string) string {
  return "Olá " + name + "!"
}

func (p Portuguese) LanguageName() string {
  return "Portuguese"
}
