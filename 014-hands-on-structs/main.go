package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "firstname",
		last:     "lsatname",
		icecream: []string{"chocolate", "butterscotch", "black currant"},
	}

	p2 := person{
		first:    "firstname2",
		last:     "lastname2",
		icecream: []string{"vanilla", "strawberry"},
	}

	//printStructPerson(p1)
	//printStructPerson(p2)

	mp := make(map[string]person)
	mp[p1.last] = p1
	mp[p2.last] = p2

	for _, val := range mp {
		fmt.Println(val.first, val.last)
		for _, v := range val.icecream {
			fmt.Println(v)
		}

	}

	//CreateEmbbededStructs()
	createAnonStruct()

}

func printStructPerson(p person) {
	fmt.Println(p.first, p.last)
	for _, val := range p.icecream {
		fmt.Println(val)
	}
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func CreateEmbbededStructs() {
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Honda",
		model: "Civic",
		doors: 5,
		color: "black",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Hyundai",
		model: "Ioniq5",
		doors: 5,
		color: "grey",
	}

	fmt.Println(v1.engine, v1.make, v1.model, v1.doors, v1.color)
	fmt.Println(v1.engine, v2.make, v2.model, v2.doors, v2.color)

}

func createAnonStruct() {
	as := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "firstname",
		friends: map[string]int{
			"friend1": 30,
			"friend2": 40,
		},
		favDrinks: []string{"Coke", "Fanta"},
	}

	fmt.Println(as)
}
