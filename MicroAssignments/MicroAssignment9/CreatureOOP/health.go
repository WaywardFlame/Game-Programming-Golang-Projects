package main

type Health struct {
	MaxHP     float32
	CurrentHP float32
}

func NewHealth(max float32, current float32) Health {
	if current > max {
		current = max
	}
	return Health{MaxHP: max, CurrentHP: current}
}

func (h *Health) Damage(dmg float32) {
	h.CurrentHP -= dmg
	if h.CurrentHP < 0 {
		h.CurrentHP = 0
	}
}

func (h *Health) Heal(restore float32) {
	h.CurrentHP += restore
	if h.CurrentHP > h.MaxHP {
		h.CurrentHP = h.MaxHP
	}
}
