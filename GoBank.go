package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//start screen what do you want to do?
//try go prompt

type Client struct {
	name      string
	balance   int
	statement []int
}

func (c *Client) setBalance() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())
	c.balance = input
}

func (c *Client) deposit() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())
	c.balance += input

}

func (c Client) getBalance() {
	fmt.Println(c.balance)
}

func main() {

	c1 := Client{"abdi", 40, []int{30, 40, 50}}

	//input wants to be used all the time even when I dont call deposit

	// c1.deposit(input)
	// fmt.Println(c1.balance)
	c1.setBalance()
	fmt.Println(c1.balance)

	fmt.Println("now add to your balance")
	c1.deposit()
	fmt.Println(c1.balance)
	//c1.getBalance()
}
