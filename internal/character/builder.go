package character

import (
	"lab3/internal/equipment"
)

// CharacterBuilder - строитель для персонажа
type CharacterBuilder struct {
	name           string
	characterClass CharacterClass
	weapon         equipment.Weapon
	armor          equipment.Armor
}

// NewBuilder - создание нового строителя
func NewBuilder() *CharacterBuilder {
	return &CharacterBuilder{}
}

// SetName - установка имени
func (cb *CharacterBuilder) SetName(name string) *CharacterBuilder {
	cb.name = name
	return cb
}

// SetClass - установка класса
func (cb *CharacterBuilder) SetClass(class CharacterClass) *CharacterBuilder {
	cb.characterClass = class
	return cb
}

// SetWeapon - установка оружия
func (cb *CharacterBuilder) SetWeapon(weapon equipment.Weapon) *CharacterBuilder {
	cb.weapon = weapon
	return cb
}

// SetArmor - установка брони
func (cb *CharacterBuilder) SetArmor(armor equipment.Armor) *CharacterBuilder {
	cb.armor = armor
	return cb
}

// Build - создание персонажа
func (cb *CharacterBuilder) Build() *PlayableCharacter {
	return NewCharacter(cb.name, cb.characterClass, cb.weapon, cb.armor)
}
