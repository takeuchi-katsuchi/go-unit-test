package car

import "fmt"

type Car struct {
	name  string
	color string
	price int
	rank  string
}

func InitCar() {
	Honda := Car{"honda", "青", 500, ""}
	Honda.rank = Rank(Honda)
	fmt.Printf("name:%v\n", Honda.name)
	fmt.Printf("color:%v\n", Honda.color)
	fmt.Printf("price:%v\n", Honda.price)
	fmt.Printf("rank:%v\n", Honda.rank)
}

func Rank(p Car) string {
	if p.price > 500 {
		return "high"
	} else if p.price > 300 {
		return "middle"
	} else {
		return "low"
	}
}
