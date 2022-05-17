package test

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &veggeMania{}

	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	if pizza.getPrice() != 15 {
		t.Errorf("pizza.getPrice() returns %v, not 15", pizza.getPrice())
	}
	if pizzaWithCheese.getPrice() != 25 {
		t.Errorf("pizzaWithCheese.getPrice() returns %v, not 22", pizzaWithCheese.getPrice())
	}
	if pizzaWithCheeseAndTomato.getPrice() != 32 {
		t.Errorf("pizzaWithCheeseAndTomato.getPrice() returns %v, not 32", pizzaWithCheeseAndTomato.getPrice())
	}

}

type pizza interface {
	getPrice() int
}

type veggeMania struct{}

func (p *veggeMania) getPrice() int {
	return 15
}

type tomatoTopping struct {
	pizza pizza
}

type cheeseTopping struct {
	pizza pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
