package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func pastMonthCommits() (int, error) {
	wc := exec.Command("wc", "-l")
	wcPipe, err := wc.StdinPipe()
	if err != nil {
		fmt.Println("error creating pipe")
		return 0, err
	}
	commits, err := exec.Command("git", "log", "--since='last month'", "--pretty=oneline").Output()
	if err != nil {
		fmt.Println("error executing git log")
		return 0, err
	}
	go func() {
		defer wcPipe.Close()
		wcPipe.Write(commits)
	}()
	amount, err := wc.Output()
	if err != nil {
		fmt.Println("error getting output")
		return 0, err
	}
	if len(amount) < 1 {
		return 0, fmt.Errorf("git log is broken")
	}
	amount = amount[:len(amount)-1]

	// SHAME SHAME SHAME SHAME
	return strconv.Atoi(fmt.Sprintf("%s", amount))
}

func littlePiggyPiggy(howMuch int) string {
	return fmt.Sprintf(""+
		"|\\_,,____ \n"+
		"( o__o \\/            PIGGY IS PROUD OF YOU\n"+
		"/(..)  \\                 YOU'VE EARNED\n"+
		"(_ )--( _)                  \033[1;92m  $%d \033[0m \n"+
		"/ \"\"--\"\" \\ \n"+
		"", howMuch)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Chanchite needs your salary as a parameter!")
		os.Exit(1)
	}
	salary := os.Args[1]
	intSalary, err := strconv.Atoi(salary)
	if err != nil {
		fmt.Printf("Check your env variable is a valid int: %s", salary)
		os.Exit(1)
	}
	commits, err := pastMonthCommits()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		os.Exit(1)
	}
	fmt.Print(littlePiggyPiggy(intSalary / commits))
}
