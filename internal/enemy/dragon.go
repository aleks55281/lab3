package enemy

import (
	"fmt"
	"lab3/internal/logger"
	"math"
)

// Dragon - дракон
type Dragon struct {
	logger     *logger.GameLogger
	name       string
	health     int
	damage     int
	resistance float32
}

func NewDragon() *Dragon {
	return &Dragon{
		logger:     logger.GetInstance(),
		name:       "Дракон",
		health:     100,
		damage:     30,
		resistance: 0.2,
	}
}

func (d *Dragon) GetName() string {
	return d.name
}

func (d *Dragon) GetHealth() int {
	return d.health
}

func (d *Dragon) TakeDamage(damage int) {
	reducedDamage := int(math.Round(float64(damage) * float64(1-d.resistance)))
	d.logger.Log(fmt.Sprintf("%s получает %d урона!", d.name, reducedDamage))
	d.health -= reducedDamage
	if d.health > 0 {
		d.logger.Log(fmt.Sprintf("У %s осталось %d здоровья", d.name, d.health))
	}
}

func (d *Dragon) Attack(player interface {
	GetName() string
	TakeDamage(int)
}) {
	d.logger.Log("Дракон дышит огнем!")
	player.TakeDamage(d.damage)
}

func (d *Dragon) IsAlive() bool {
	return d.health > 0
}
