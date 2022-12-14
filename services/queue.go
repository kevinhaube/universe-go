package services

import (
	"math/rand"
	"time"
	ingredients "universe-go/models"
	order "universe-go/models"
	ordering "universe-go/services/handlers"
)

// Randomly select an ingredient from an array of choices
// mimicking the behavior of a customer choosing between white or brown rice
// for example.
func chooseIngredient(choices []ingredients.MealIngredient) ingredients.MealIngredient {
	index := rand.Intn(len(choices) - 1)
	return choices[index]
}

// Randomly decide (with bias to false) whether item is on the side or not
// mimicking a baseless assumption that there's a 20% chance of of somebody
// asking for an ingredient on the side.
func chooseOnTheSide() bool {
	return rand.Intn(10) < 2
}

// Compiles the Option structure, like a customer choosing ingredients for
// their meal.
func MakeOption(choices []ingredients.MealIngredient) order.CustomerChoice {
	return order.CustomerChoice{
		Ingredient: chooseIngredient(choices),
		Multiplier: rand.Intn(2) + 1,
		Side:       chooseOnTheSide(),
	}
}

// Creates a channel and starts a Goroutine that adds a new customer to the
// queue every random-n seconds. The queue uses a shared channel to coordinate
// with the OrderTaker.
func StartCustomerQueue() {
	customerQueue := make(chan int)
	go ordering.HandleCustomerQueue(customerQueue)
	nextCustomerInQueue := 0
	for {
		nextCustomerInQueue += 1
		customerQueue <- nextCustomerInQueue
		nextCustomerTimeout := rand.Intn(10) // Simulate randomly timed queue growth
		time.Sleep(time.Second * time.Duration(nextCustomerTimeout))
	}
	// TODO: How to keep queue running while the restaurant moves through the line
}
