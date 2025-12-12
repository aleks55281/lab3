package location

import "lab3/internal/enemy"

// Forest - мистический лес
type Forest struct {
	*BaseLocation
}

func NewForest() *Forest {
	return &Forest{
		BaseLocation: NewBaseLocation(),
	}
}

func (f *Forest) EnterLocation(playerName string) enemy.Enemy {
	f.logger.Log("Вы вошли в Мистический Лес...")
	enemy := f.spawnEnemy()
	f.logger.Log("Из-за деревьев появляется " + enemy.GetName() + "!")
	return enemy
}

func (f *Forest) GetLocationName() string {
	return "Мистический Лес"
}

// spawnEnemy - фабричный метод
func (f *Forest) spawnEnemy() enemy.Enemy {
	return enemy.NewGoblin()
}
