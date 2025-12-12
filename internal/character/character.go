package character

import (
	"fmt"
	"lab3/internal/equipment"
	"lab3/internal/logger"
)

// PlayableCharacter - игровой персонаж
type PlayableCharacter struct {
	logger *logger.GameLogger
	Name   string
	Class  CharacterClass
	Weapon equipment.Weapon
	Armor  equipment.Armor
	Health int
}

// NewCharacter - конструктор
func NewCharacter(name string, class CharacterClass, weapon equipment.Weapon, armor equipment.Armor) *PlayableCharacter {
	return &PlayableCharacter{
		logger: logger.GetInstance(),
		Name:   name,
		Class:  class,
		Weapon: weapon,
		Armor:  armor,
		Health: class.StartingHealth(),
	}
}

// GetName - получаем имя персонажа
func (pc *PlayableCharacter) GetName() string {
	return pc.Name
}

// TakeDamage - получение урона
func (pc *PlayableCharacter) TakeDamage(damage int) {
	reducedDamage := int(float32(damage) * (1 - pc.Armor.GetDefense()))
	if reducedDamage < 0 {
		reducedDamage = 0
	}

	pc.Health -= reducedDamage
	pc.Armor.Use()
	pc.logger.Log(fmt.Sprintf("%s получил урон: %d", pc.Name, reducedDamage))

	if pc.Health > 0 {
		pc.logger.Log(fmt.Sprintf("У %s осталось %d здоровья", pc.Name, pc.Health))
	}
}

// Attack - атака врага
func (pc *PlayableCharacter) Attack(enemy interface {
	GetName() string
	TakeDamage(int)
}) {
	pc.logger.Log(fmt.Sprintf("%s атакует врага %s", pc.Name, enemy.GetName()))
	pc.Weapon.Use()
	enemy.TakeDamage(pc.Weapon.GetDamage())
}

// IsAlive - проверка жив ли персонаж
func (pc *PlayableCharacter) IsAlive() bool {
	return pc.Health > 0
}
