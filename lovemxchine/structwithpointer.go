package lovemxchine

import "fmt"

type Weapon struct {
	Type       string
	Rarity     string
	Sharpness  int
	Durability int
}

func changeStats(weapon *Weapon, sharp int) {
	weapon.Sharpness += sharp
}

func SetWeapon() {
	katana := Weapon{
		Type:       "sword",
		Rarity:     "rare",
		Sharpness:  20,
		Durability: 100,
	}
	changeStats(&katana, 5)
	fmt.Println(katana)
}