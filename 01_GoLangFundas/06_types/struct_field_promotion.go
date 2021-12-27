package main

import "fmt"

type cost struct {
	sellingPrice, costPrice int
}

type item struct {
	itemName string
	price    cost
}

type goood struct {
	itemName string
	cost
}

func main() {
	item1 := item{
		itemName: "Rice Bag",
		price:    cost{450, 400},
	}
	goood1 := goood{
		itemName:     "Channa Dal Bag",
		cost:cost{
			sellingPrice: 679,
			costPrice:    625,
		}
	}

	fmt.Println("Margin: ",item1.price.sellingPrice-item1.price.costPrice)
	fmt.Println("Margin: ",goood1.sellingPrice-goood1.costPrice)
}
