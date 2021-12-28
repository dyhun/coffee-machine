package main

import (
	"fmt"
	"math"
)

func main() {
	water := 400
	milk := 540
	beans := 120
	cups := 9
	money := 550

	var action string
	for action != "exit" {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		switch action {
		case "remaining":
			PrintState(water, milk, beans, cups, money)
		case "buy":
			BuyCoffee(&water, &milk, &beans, &cups, &money)
		case "fill":
			FillMachine(&water, &milk, &beans, &cups)
		case "take":
			TakeMoney(&money)
		}
	}
}

func BuyCoffee(water, milk, beans, cups, money *int) {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	coffee := 0
	for coffee < 1 || coffee > 3 {
		fmt.Scan(&coffee)
	}
	MakeCoffee(&coffee, water, milk, beans, cups, money)
}

/*

	One espresso requires 250 ml of water and 16 g of coffee beans. It costs $4;
	One latte requires 350 ml of water, 75 ml of milk, and 20 g of coffee beans. It costs $7;
	One cappuccino requires 200 ml of water, 100 ml of milk, and 12 g of coffee beans. It costs $6.

*/
func MakeCoffee(coffee, water, milk, beans, cups, money *int) {
	switch *coffee {
	case 1:
		{
			MakeEspresso(water, beans, cups, money)
		}
	case 2:
		{
			MakeLatte(water, milk, beans, cups, money)
		}
	case 3:
		{
			MakeCappuccino(water, milk, beans, cups, money)
		}
	}
}

func MakeCappuccino(water *int, milk *int, beans *int, cups *int, money *int) {
	if *water >= 200 {
		if *milk >= 100 {
			if *beans >= 12 {
				if *cups >= 1 {
					*water -= 200
					*milk -= 100
					*beans -= 12
					*money += 6
					fmt.Println("I have enough resources, making you a coffee!")
				} else {
					fmt.Println("Sorry, not enough disposable cups!")
				}
			} else {
				fmt.Println("Sorry, not enough coffee beans!")
			}
		} else {
			fmt.Println("Sorry, not enough milk!")
		}
	} else {
		fmt.Println("Sorry, not enough water!")
	}
}

func MakeLatte(water *int, milk *int, beans *int, cups *int, money *int) {
	if *water >= 350 {
		if *milk >= 75 {
			if *beans >= 20 {
				if *cups >= 1 {
					*water -= 350
					*milk -= 75
					*beans -= 20
					*money += 7
					fmt.Println("I have enough resources, making you a coffee!")
				} else {
					fmt.Println("Sorry, not enough disposable cups!")
				}
			} else {
				fmt.Println("Sorry, not enough coffee beans!")
			}
		} else {
			fmt.Println("Sorry, not enough milk!")
		}
	} else {
		fmt.Println("Sorry, not enough water!")
	}
}

func MakeEspresso(water *int, beans *int, cups *int, money *int) {
	if *water >= 250 {
		if *beans >= 16 {
			if *cups >= 1 {
				*water -= 250
				*beans -= 16
				*cups--
				*money += 4
				fmt.Println("I have enough resources, making you a coffee!")
			} else {
				fmt.Println("Sorry, not enough disposable cups!")
			}
		} else {
			fmt.Println("Sorry, not enough coffee beans!")
		}
	} else {
		fmt.Println("Sorry, not enough water!")
	}
}

func TakeMoney(money *int) {
	fmt.Printf("I gave you $%d\n\n", *money)
	*money = 0
}

func FillMachine(water, milk, beans, cups *int) {
	var add int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&add)
	*water += add
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&add)
	*milk += add
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&add)
	*beans += add
	fmt.Println("Write how many disposable coffee cups you want to add:")
	fmt.Scan(&add)
	*cups += add
}

func PrintState(water, milk, beans, cups, money int) {
	fmt.Printf("\nThe coffee machine has:\n%d of water\n%d of milk\n%d of coffee beans\n%d of disposable cups\n%d of money\n",
		water, milk, beans, cups, money)
	fmt.Println()
}

func NumberOfCups(water, milk, beans int) int {
	waterLimit := water / 200
	milkLimit := milk / 50
	beansLimit := beans / 15
	return int(math.Min(float64(waterLimit), math.Min(float64(milkLimit), float64(beansLimit))))
}
