package location

import "lab3/internal/enemy"

// DragonBarrow - логово дракона
type DragonBarrow struct {
	*BaseLocation
}

func NewDragonBarrow() *DragonBarrow {
	return &DragonBarrow{
		BaseLocation: NewBaseLocation(),
	}
}

func (db *DragonBarrow) EnterLocation(playerName string) enemy.Enemy {
	db.logger.Log("Вы спускаетесь в Драконье Логово...")
	enemy := db.spawnEnemy()
	db.logger.Log("Из темноты появляется огромный " + enemy.GetName() + "!")
	return enemy
}

func (db *DragonBarrow) GetLocationName() string {
	return "Драконье Логово"
}

// spawnEnemy - фабричный метод
func (db *DragonBarrow) spawnEnemy() enemy.Enemy {
	return enemy.NewDragon()
}
