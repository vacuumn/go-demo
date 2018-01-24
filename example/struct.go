package example

type Rabbit struct {
	Name  string
	Color int
}

func (r Rabbit) Speak() string {
	return "How long is forever?"
}
