package equipment

import "lab3/internal/logger"

// Armor - интерфейс брони
type Armor interface {
	GetDefense() float32
	Use()
}

// HeavyArmor - тяжелая броня
type HeavyArmor struct {
	defense float32
	logger  *logger.GameLogger
}

func NewHeavyArmor() *HeavyArmor {
	return &HeavyArmor{
		defense: 0.3,
		logger:  logger.GetInstance(),
	}
}

func (ha *HeavyArmor) GetDefense() float32 {
	return ha.defense
}

func (ha *HeavyArmor) Use() {
	ha.logger.Log("Тяжелая броня блокирует значительную часть урона")
}

// LightArmor - легкая броня
type LightArmor struct {
	defense float32
	logger  *logger.GameLogger
}

func NewLightArmor() *LightArmor {
	return &LightArmor{
		defense: 0.2,
		logger:  logger.GetInstance(),
	}
}

func (la *LightArmor) GetDefense() float32 {
	return la.defense
}

func (la *LightArmor) Use() {
	la.logger.Log("Легкая броня блокирует урон")
}

// Robe - роба
type Robe struct {
	defense float32
	logger  *logger.GameLogger
}

func NewRobe() *Robe {
	return &Robe{
		defense: 0.1,
		logger:  logger.GetInstance(),
	}
}

func (r *Robe) GetDefense() float32 {
	return r.defense
}

func (r *Robe) Use() {
	r.logger.Log("Роба блокирует немного урона")
}
