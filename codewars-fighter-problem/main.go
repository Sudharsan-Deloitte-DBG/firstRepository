package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func Attack(f1, f2 *Fighter) {
	f2.Health -= f1.DamagePerAttack
}

func DeclareWinner(f1, f2 Fighter, firstAttacker string) string {
	if f1.Name == firstAttacker {
		for {
			if f1.Health <= 0 || f2.Health <= 0 {
				break
			}
			Attack(&f1, &f2)
			fmt.Println(f1, f2)
			Attack(&f2, &f1)
			fmt.Println(f1, f2)
		}
	} else {
		for {
			if f1.Health <= 0 || f2.Health <= 0 {
				break
			}
			Attack(&f2, &f1)
			fmt.Println(f1, f2)
			Attack(&f1, &f2)
			fmt.Println(f1, f2)
		}
	}
	if f1.Health > 0 {
		return f1.Name
	} else {
		return f2.Name
	}
}

func main() {
	f1 := Fighter{
		Name:            "James",
		Health:          10,
		DamagePerAttack: 2,
	}
	f2 := Fighter{
		Name:            "Bond",
		Health:          10,
		DamagePerAttack: 1,
	}

	winner := DeclareWinner(f1, f2, f1.Name)
	fmt.Println(winner)

}
