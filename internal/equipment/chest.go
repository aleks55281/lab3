package equipment

// EquipmentChest - абстрактная фабрика
type EquipmentChest interface {
	GetWeapon() Weapon
	GetArmor() Armor
}

// WarriorEquipmentChest - сундук для воина
type WarriorEquipmentChest struct{}

func NewWarriorEquipmentChest() *WarriorEquipmentChest {
	return &WarriorEquipmentChest{}
}

func (wec *WarriorEquipmentChest) GetWeapon() Weapon {
	return NewSword()
}

func (wec *WarriorEquipmentChest) GetArmor() Armor {
	return NewHeavyArmor()
}

// MagicalEquipmentChest - сундук для мага
type MagicalEquipmentChest struct{}

func NewMagicalEquipmentChest() *MagicalEquipmentChest {
	return &MagicalEquipmentChest{}
}

func (mec *MagicalEquipmentChest) GetWeapon() Weapon {
	return NewStaff()
}

func (mec *MagicalEquipmentChest) GetArmor() Armor {
	return NewRobe()
}

// ThiefEquipmentChest - сундук для вора
type ThiefEquipmentChest struct{}

func NewThiefEquipmentChest() *ThiefEquipmentChest {
	return &ThiefEquipmentChest{}
}

func (tec *ThiefEquipmentChest) GetWeapon() Weapon {
	return NewBow()
}

func (tec *ThiefEquipmentChest) GetArmor() Armor {
	return NewLightArmor()
}
