package equipment

import (
	"lab3/internal/logger"
	"math/rand"
)

// Weapon - интерфейс оружия
type Weapon interface {
	GetDamage() int
	Use()
}

// Sword - меч
type Sword struct {
	damage int
	logger *logger.GameLogger
}

func NewSword() *Sword {
	return &Sword{
		damage: 20,
		logger: logger.GetInstance(),
	}
}

func (s *Sword) GetDamage() int {
	return s.damage
}

func (s *Sword) Use() {
	s.logger.Log("Удар мечом!")
}

// Bow - лук
type Bow struct {
	damage           int
	criticalChance   float64
	criticalModifier int
	logger           *logger.GameLogger
}

func NewBow() *Bow {
	return &Bow{
		damage:           15,
		criticalChance:   0.3,
		criticalModifier: 2,
		logger:           logger.GetInstance(),
	}
}

func (b *Bow) GetDamage() int {
	if rand.Float64() <= b.criticalChance {
		b.logger.Log("Критический урон!")
		return b.damage * b.criticalModifier
	}
	return b.damage
}

func (b *Bow) Use() {
	b.logger.Log("Выстрел из лука!")
}

// Staff - посох
type Staff struct {
	damage  int
	scatter float64
	logger  *logger.GameLogger
}

func NewStaff() *Staff {
	return &Staff{
		damage:  25,
		scatter: 0.2,
		logger:  logger.GetInstance(),
	}
}

func (s *Staff) GetDamage() int {
	roll := rand.Float64()
	factor := 1 + (roll*2*s.scatter - s.scatter)
	return int(float64(s.damage) * factor)
}

func (s *Staff) Use() {
	s.logger.Log("Воздух накаляется, из посоха вылетает огненный шар!")
}
