package model

type hero struct {
	name string
	equipments []equipment
}

func (h *hero) getEquip() []equipment {
	return h.equipments
}

type action interface {
	attack()
	def()
	changeEquip() *equipment
}

func (h *hero) attack ()  {
	j :=0
	for range h.equipments{
		h.equipments[j].attack += h.equipments[j].attack
	}
}
