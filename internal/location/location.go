package location

import (
	"lab3/internal/enemy"
	"lab3/internal/logger"
)

// Location - абстрактная локация (фабричный метод)
type Location interface {
	EnterLocation(playerName string) enemy.Enemy
	GetLocationName() string
}

// BaseLocation - базовая реализация локации
type BaseLocation struct {
	logger *logger.GameLogger
}

func NewBaseLocation() *BaseLocation {
	return &BaseLocation{
		logger: logger.GetInstance(),
	}
}
