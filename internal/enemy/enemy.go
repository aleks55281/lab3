package enemy

// Enemy - абстрактный враг
type Enemy interface {
	GetName() string
	GetHealth() int
	TakeDamage(damage int)
	Attack(player interface {
		GetName() string
		TakeDamage(int)
	})
	IsAlive() bool
}
