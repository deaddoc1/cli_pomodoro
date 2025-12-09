package main

import "fmt"

func main() {

	var pomodoro, rest, answ int

	fmt.Println("Welcome to CLI_Pomodoro!")
	fmt.Print("Pomodoro time (in min) (25 or 50): ")
	fmt.Scan(&pomodoro)

	fmt.Print("Rest time (in min) (5 or 10): ")
	fmt.Scan(&rest)

	for {
		clearConsole()

		fmt.Println("Good luck (>ω<)!")

		countdown(pomodoro, "Pomodoro")
		fmt.Println("Pomodoro is over! Time to rest.")
		beep()
		beep()
		beep()
		beep()
		beep()
		fmt.Println("Are you ready for rest?: ")
		fmt.Println("1 - yes")
		fmt.Println("2 - no")
		fmt.Scan(&answ)

		if answ != 1 {
			fmt.Println("Good luck, bye!")
			break
		}

		clearConsole()

		fmt.Println("Chill (─‿‿─)!")

		countdown(rest, "Rest")
		fmt.Println("Rest is over! Time for Pomodoro!")
		beep()
		beep()
		beep()
		beep()
		beep()
		fmt.Println("Are you ready for pomodoro?: ")
		fmt.Println("1 - yes")
		fmt.Println("2 - no")
		fmt.Scan(&answ)

		if answ != 1 {
			fmt.Println("Good luck, bye!")
			break
		}

	}

}
