package character

// CharacterClass - классы персонажей
type CharacterClass string

const (
	Warrior CharacterClass = "WARRIOR"
	Thief   CharacterClass = "THIEF"
	Mage    CharacterClass = "MAGE"
)

// StartingHealth - начальное здоровье для каждого класса
func (cc CharacterClass) StartingHealth() int {
	switch cc {
	case Warrior:
		return 100
	case Thief:
		return 90
	case Mage:
		return 80
	default:
		return 0
	}
}

// String - строковое представление
func (cc CharacterClass) String() string {
	return string(cc)
}
