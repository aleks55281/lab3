package main

import (
	"bufio"
	"fmt"
	"lab3/internal/character"
	"lab3/internal/equipment"
	"lab3/internal/location"
	"lab3/internal/logger"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	gameLogger := logger.GetInstance()

	fmt.Println("=== ТРОПА СУДЬБЫ ===")
	fmt.Println("Создайте своего персонажа:")

	// Ввод имени
	fmt.Print("Введите имя персонажа: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Выбор класса
	fmt.Println("\nВыберите класс персонажа:")
	fmt.Println("1. Воин (WARRIOR) - 100 HP")
	fmt.Println("2. Вор (THIEF) - 90 HP")
	fmt.Println("3. Маг (MAGE) - 80 HP")
	fmt.Print("Ваш выбор (1-3): ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var charClass character.CharacterClass
	switch choice {
	case "1":
		charClass = character.Warrior
	case "2":
		charClass = character.Thief
	case "3":
		charClass = character.Mage
	default:
		fmt.Println("Неверный выбор. По умолчанию выбран Воин.")
		charClass = character.Warrior
	}

	// Получение стартового снаряжения (Абстрактная фабрика)
	var chest equipment.EquipmentChest
	switch charClass {
	case character.Warrior:
		chest = equipment.NewWarriorEquipmentChest()
	case character.Thief:
		chest = equipment.NewThiefEquipmentChest()
	case character.Mage:
		chest = equipment.NewMagicalEquipmentChest()
	}

	startingWeapon := chest.GetWeapon()
	startingArmor := chest.GetArmor()

	// Создание персонажа (Builder)
	player := character.NewBuilder().
		SetName(name).
		SetClass(charClass).
		SetWeapon(startingWeapon).
		SetArmor(startingArmor).
		Build()

	gameLogger.Log(fmt.Sprintf("%s очнулся на распутье!", player.Name))

	// Выбор локации
	fmt.Println("\nКуда вы двинетесь?")
	fmt.Println("1. Мистический Лес (против Гоблина)")
	fmt.Println("2. Драконье Логово (против Дракона)")
	fmt.Print("Ваш выбор (1-2): ")

	locationChoice, _ := reader.ReadString('\n')
	locationChoice = strings.TrimSpace(locationChoice)

	var currentLocation location.Location
	switch locationChoice {
	case "1":
		currentLocation = location.NewForest()
	case "2":
		currentLocation = location.NewDragonBarrow()
	default:
		fmt.Println("Неверный выбор. По умолчанию выбран Мистический Лес.")
		currentLocation = location.NewForest()
	}

	// Вход в локацию (Фабричный метод создает врага)
	enemy := currentLocation.EnterLocation(player.Name)

	fmt.Println("\n=== БОЙ НАЧИНАЕТСЯ! ===")
	fmt.Println("Нажмите Enter для атаки...")

	// Игровой цикл
	for player.IsAlive() && enemy.IsAlive() {
		reader.ReadString('\n')

		// Атака игрока
		player.Attack(enemy)

		if !enemy.IsAlive() {
			break
		}

		// Есть шанс, что враг будет оглушен
		if rand.Intn(2) == 0 {
			gameLogger.Log(fmt.Sprintf("%s был оглушен атакой %s!", enemy.GetName(), player.Name))
			continue
		}

		// Атака врага
		enemy.Attack(player)
	}

	fmt.Println("\n=== БОЙ ЗАКОНЧЕН! ===")

	if !player.IsAlive() {
		gameLogger.Log(fmt.Sprintf("%s был убит...", player.Name))
	} else {
		gameLogger.Log(fmt.Sprintf("Злой %s был побежден! %s отправился дальше по тропе судьбы...",
			enemy.GetName(), player.Name))
	}

	// Вывод всех логов
	fmt.Println("\n=== ИГРОВЫЕ ЛОГИ ===")
	for _, log := range gameLogger.GetLogs() {
		fmt.Println(log)
	}
}
