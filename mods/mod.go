package mods

import "fmt"

type mod struct {
	name string
	id   int
}

func (m mod) GetName() string {
	return m.name
}

func (m mod) GetId() int {
	return m.id
}

// String returns a string with the name and id of the mod.
func (m mod) String() string {
	return fmt.Sprintf("{name: %s, id: %d}", m.name, m.id)
}
