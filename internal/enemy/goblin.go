package enemy

import (
	"fmt"
	"lab3/internal/logger"
)

// Goblin - гоблин
type Goblin struct {
	logger *logger.GameLogger
	name   string
	health int
	damage int
}

func NewGoblin() *Goblin {
	return &Goblin{
		logger: logger.GetInstance(),
		name:   "Гоблин",
		health: 50,
		damage: 10,
	}
}

func (g *Goblin) GetName() string {
	return g.name
}

func (g *Goblin) GetHealth() int {
	return g.health
}

func (g *Goblin) TakeDamage(damage int) {
	g.logger.Log(fmt.Sprintf("%s получает %d урона!", g.name, damage))
	g.health -= damage
	if g.health > 0 {
		g.logger.Log(fmt.Sprintf("У %s осталось %d здоровья", g.name, g.health))
	}
}

func (g *Goblin) Attack(player interface {
	GetName() string
	TakeDamage(int)
}) {
	g.logger.Log(fmt.Sprintf("%s атакует %s!", g.name, player.GetName()))
	player.TakeDamage(g.damage)
}

func (g *Goblin) IsAlive() bool {
	return g.health > 0
}
