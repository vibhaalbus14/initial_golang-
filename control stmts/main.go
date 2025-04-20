/*
1.if {
}else => closing } and else must be on same line

2.func func_name(){
}
=> func signature and opening braces on the same line

3.You can pass an integer to os.Exit() to indicate the exit status.
os.Exit(0) means the program was successful, while non-zero values (e.g., os.Exit(1)) indicate an error.
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeToFile(currBalance float32) {
	writeBalance := fmt.Sprint(currBalance)
	err := os.WriteFile("balance.txt", []byte(writeBalance), 0644) //string ->[]byte
	if err != nil {
		fmt.Print("error while writing to file", err)
	} else {
		fmt.Println("write to file successful !")
	}
}

func readFromFile() (float32, error) {
	readBalance, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Print("error while reading from file", err)
		return 1000, errors.New("failed to fetch balance")
		//panic("sorry we cannot continue due to an error") //return the program after printing a message
		// return 0.0
		// return
	} else {
		fmt.Println("read from  file successful !")
		val, _ := strconv.ParseFloat(string(readBalance), 32) //string to float conv
		return float32(val), nil
	}
}

func main() {
	//var balance float32
	//get balance from balance.txt
	balance, errorMessage := readFromFile()
	if errorMessage != nil {
		fmt.Println("error encountered while reading balance ,", errorMessage)
		return
	}

	fmt.Print("balance is", balance)
	var choice int

	for i := 0; i < 100; i++ {
		fmt.Println("Welcome to banking app.What would you like to do ? \n 1.Check Balance \n 2.Withdraw money \n 3.Deposit money \n 4.Exit")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Here's your balance : %0.2f \n", balance)
		case 2:
			var withdrawl float32
			fmt.Printf("Enter the amount to be withdrawn:")
			fmt.Scanln(&withdrawl)
			//check if its possible
			if withdrawl > balance {
				fmt.Println("Withdrawl money greater than balance. Cannot proceed")
			} else {
				balance -= withdrawl
				fmt.Printf("Withdraw %0.2f , remaining balance %0.2f \n", withdrawl, balance)
			}
		case 3:
			var deposit float32
			fmt.Println("Enter money to be deposited")
			fmt.Scanln(&deposit)
			if deposit < 0 {
				fmt.Println("Deposit amount cannot be negative")
			} else {
				balance += deposit
				fmt.Printf("Amount %0.3f deposited.Remaining balance is %0.2f", deposit, balance)

			}
		case 4:
			writeToFile(balance)
			fmt.Println("Program about to exit")
			os.Exit(0)
		}
	}
}
