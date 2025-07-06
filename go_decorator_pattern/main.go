package main

import "fmt"

type Pizza interface {
	getPrice() int
}

// Pizzas
type SimplePizza struct {
}

func (s *SimplePizza) getPrice() int {
	return 200
}

type GarlicPizza struct {
}

func (g *GarlicPizza) getPrice() int {
	return 250
}

// Toppings
type VegiTopping struct {
	pizza Pizza
}

func (v *VegiTopping) getPrice() int {
	pizzaPrice := v.pizza.getPrice()
	return pizzaPrice + 100
}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 150
}

type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 50
}

func main() {
	simplePizza := &SimplePizza{}

	simplePizzaWithVegi := &VegiTopping{
		pizza: simplePizza,
	}
	simplePizzaWithVegiAndCheese := &CheeseTopping{
		pizza: simplePizzaWithVegi,
	}
	simplePizzaWithAllToppings := &TomatoTopping{
		pizza: simplePizzaWithVegiAndCheese,
	}

	fmt.Println("Price of Simple Pizza with Vegi, Cheese and Tomato Topping: ", simplePizzaWithAllToppings.getPrice())

	garlicPizza := &GarlicPizza{}

	garlicPizzaWithVegi := &VegiTopping{
		pizza: garlicPizza,
	}
	garlicPizzaWithCheese := &CheeseTopping{
		pizza: garlicPizzaWithVegi,
	}
	garlicPizzaWithAllToppings := &TomatoTopping{
		pizza: garlicPizzaWithCheese,
	}

	fmt.Println("Price of Garlic Pizza with Vegi, Cheese and Tomato Topping: ", garlicPizzaWithAllToppings.getPrice())
}
