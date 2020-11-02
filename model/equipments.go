package model

type equipment struct {
	name string
	attack int
	def int
}

func newEquipment(name string, attack int, def int) *equipment{
	return &equipment{
		name: name,
		attack: attack,
		def: def,
		}
}

